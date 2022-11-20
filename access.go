package vyze

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const seedLength = 4
const signLength = 16

// LAYER TOKEN

// LayerToken is a parsed layer token issued by the VYZE service.
type LayerToken struct {
	LayerID   ID            `json:"layer"`
	UserID    ID            `json:"user"`
	Created   time.Time     `json:"created"`
	Expiry    time.Duration `json:"expiry"`
	Granted   uint32        `json:"granted"`
	Mandatory uint32        `json:"mandatory"`
	Exclusive uint32        `json:"exclusive"`
	IsAdmin   bool          `json:"admin"`
	Signature Binary        `json:"signature"`
	Token     string        `json:"token"`
}

// Expires returns the time the token expires.
func (lt LayerToken) Expires() time.Time {
	if lt.Expiry == -1 {
		return time.Unix(1<<63-1, 0)
	}
	return lt.Created.Add(lt.Expiry)
}

// Expired returns true if the token has expired.
func (lt LayerToken) Expired() bool {
	if lt.Expiry == -1 {
		return false
	}
	return lt.Created.Add(lt.Expiry).Before(time.Now())
}

// ReadLayerToken parses a token string and returns a new LayerToken instance.
func ReadLayerToken(tokenString string) (LayerToken, error) {
	const permLen = 4
	const msgLen = 16 + 16 + 3*permLen + 4 + 8 + 1

	var msgBytes []byte
	var err error

	msgBytes, err = hex.DecodeString(tokenString)
	if err == nil && len(msgBytes) == msgLen+seedLength+signLength {
		goto succ
	}

	msgBytes, err = base32.HexEncoding.WithPadding(base32.NoPadding).DecodeString(tokenString)
	if err == nil && len(msgBytes) == msgLen+seedLength+signLength {
		goto succ
	}

	msgBytes, err = base64.RawURLEncoding.DecodeString(tokenString)
	if err == nil && len(msgBytes) == msgLen+seedLength+signLength {
		goto succ
	}

	return LayerToken{}, errors.New("invalid token")

succ:
	s := LayerToken{
		Token: tokenString,
	}

	i := 0

	i += copy(s.UserID[:], msgBytes[i:])
	i += copy(s.LayerID[:], msgBytes[i:])

	s.Granted = binary.BigEndian.Uint32(msgBytes[i:])
	i += permLen

	s.Mandatory = binary.BigEndian.Uint32(msgBytes[i:])
	i += permLen

	s.Exclusive = binary.BigEndian.Uint32(msgBytes[i:])
	i += permLen

	s.Created = time.Unix(int64(binary.BigEndian.Uint64(msgBytes[i:])), 0)
	i += 8

	s.Expiry = time.Duration(binary.BigEndian.Uint32(msgBytes[i:])) * time.Second
	i += 4

	if s.Expired() {
		return LayerToken{}, errors.New("expired:" + s.LayerID.String())
	}

	s.IsAdmin = msgBytes[i] == 0xFF
	i += 1

	s.Signature = msgBytes[msgLen+seedLength:]

	return s, nil
}

// ACCESS GROUP

// AccessGroup is a group of layer tokens having at least the group's permissions.
type AccessGroup struct {
	Name        string
	Permissions uint32
	Tokens      []LayerToken
}

// RegisterLayerToken adds a layer token to this access group.
// It returns an error if the layer token does not have sufficient permissions for that group.
func (ac *AccessGroup) RegisterLayerToken(token LayerToken) error {
	if token.Granted&ac.Permissions != ac.Permissions {
		return errors.New("insufficient permissions")
	}
	ac.Tokens = append(ac.Tokens, token)
	return nil
}

// UnregisterLayerToken removes a token from this access group.
// If the token is not present, nothing will happen.
func (ac *AccessGroup) UnregisterLayerToken(token string) {
	newTokens := []LayerToken{}
	for _, tk := range ac.Tokens {
		if tk.Token == token {
			continue
		}
		newTokens = append(newTokens, tk)
	}
	ac.Tokens = newTokens
}

// String returns a string representation of this access group.
// The string can be parsed using ReadAccessGroup.
func (ac AccessGroup) String() string {
	tks := []string{}
	for _, tk := range ac.Tokens {
		tks = append(tks, tk.Token)
	}
	return fmt.Sprintf("%s:%s:%s", ac.Name, strconv.FormatUint(uint64(ac.Permissions), 16), strings.Join(tks, ","))
}

func newAccessGroup(name string, permissions uint32) *AccessGroup {
	return &AccessGroup{
		Name:        name,
		Permissions: permissions,
		Tokens:      []LayerToken{},
	}
}

// ReadAccessGroup parses a string obtained by AccessGroup.String and returns an access group instance.
func ReadAccessGroup(groupString string) (*AccessGroup, error) {
	agSplit := strings.Split(groupString, ":")
	if len(agSplit) != 3 {
		return nil, errors.New("invalid group string")
	}
	perms, _ := strconv.ParseUint(agSplit[1], 16, 32)
	ag := newAccessGroup(agSplit[0], uint32(perms))
	tkStrs := strings.Split(agSplit[2], ",")
	for _, tkStr := range tkStrs {
		if len(tkStr) == 0 {
			continue
		}
		lt, err := ReadLayerToken(tkStr)
		if err != nil {
			return nil, err
		}
		if err := ag.RegisterLayerToken(lt); err != nil {
			return nil, err
		}
	}
	return ag, nil
}

// LAYER PROFILE

// LayerProfile maps names of access groups onto group instances.
type LayerProfile map[string]*AccessGroup

// AddAccessGroup creates a new access group and attaches it to the LayerProfile instance.
// If an access group with that name already exists, it will return that instance instead of creating a new one.
func (lp *LayerProfile) AddAccessGroup(name string, permissions uint32) (*AccessGroup, error) {
	acc, ok := (*lp)[name]
	if ok {
		if acc.Permissions == permissions {
			return acc, nil
		}
		return nil, errors.New("conflicting permissions")
	}
	acc = newAccessGroup(name, permissions)
	(*lp)[name] = acc
	return acc, nil
}

// GetAccessGroup returns an access group instance with the given name or nil if it does not exist.
func (lp LayerProfile) GetAccessGroup(name string) *AccessGroup {
	acc, _ := lp[name]
	return acc
}

// RemoveAccessGroup removes an access group.
func (lp LayerProfile) RemoveAccessGroup(name string) {
	delete(lp, name)
}

// String returns a string representation of the layer profile.
func (lp LayerProfile) String() string {
	ags := []string{}
	for _, v := range lp {
		ags = append(ags, v.String())
	}
	return strings.Join(ags, ";")
}

// NewLayerProfile creates a LayerProfile instance.
func NewLayerProfile() LayerProfile {
	return LayerProfile{}
}

// ReadLayerProfile parses a string representation obtained via LayerProfile.String and returns a new LayerProfile
// instance.
func ReadLayerProfile(profileString string) (LayerProfile, error) {
	lp := NewLayerProfile()
	if len(profileString) == 0 {
		return lp, nil
	}
	ags := strings.Split(profileString, ";")
	for _, agStr := range ags {
		ag, err := ReadAccessGroup(agStr)
		if err != nil {
			return nil, err
		}
		lp[ag.Name] = ag
	}
	return lp, nil
}

// PERMISSIONS

// 4 permission bytes
const (
	PermDeleteObject = Permission(1) << iota
	PermViewName     = Permission(1) << iota
	PermChangeName   = Permission(1) << iota
	PermViewData     = Permission(1) << iota
	PermChangeData   = Permission(1) << iota
	PermCreated      = Permission(1) << iota
	PermUser         = Permission(1) << iota
	PermViewAbstract = Permission(1) << iota

	PermAddAbstract    = Permission(1) << iota
	PermRemoveAbstract = Permission(1) << iota
	PermViewSpecial    = Permission(1) << iota
	PermAddSpecial     = Permission(1) << iota
	PermRemoveSpecial  = Permission(1) << iota
	PermViewTarget     = Permission(1) << iota
	PermAddTarget      = Permission(1) << iota
	PermRemoveTarget   = Permission(1) << iota

	PermViewOrigin   = Permission(1) << iota
	PermAddOrigin    = Permission(1) << iota
	PermRemoveOrigin = Permission(1) << iota
	PermViewObject   = Permission(1) << iota
	PermAddObject    = Permission(1) << iota
	PermRemoveObject = Permission(1) << iota
	PermViewSpace    = Permission(1) << iota
	PermAddSpace     = Permission(1) << iota

	PermRemoveSpace = Permission(1) << iota
	PermReserved1   = Permission(1) << iota
	PermReserved2   = Permission(1) << iota
	PermReserved3   = Permission(1) << iota
	PermMeta1       = Permission(1) << iota
	PermMeta2       = Permission(1) << iota
	PermMeta3       = Permission(1) << iota
	PermMeta4       = Permission(1) << iota
)

const PermChangeAbstract = PermAddAbstract | PermRemoveAbstract
const PermChangeSpecial = PermAddSpecial | PermRemoveSpecial
const PermChangeTarget = PermAddTarget | PermRemoveTarget
const PermChangeOrigin = PermAddOrigin | PermRemoveOrigin
const PermChangeObject = PermAddObject | PermRemoveObject
const PermChangeSpace = PermAddSpace | PermRemoveSpace

const PermName = PermViewName | PermChangeName
const PermData = PermViewData | PermChangeData
const PermAbstract = PermViewAbstract | PermChangeAbstract
const PermSpecial = PermViewSpecial | PermChangeSpecial
const PermTarget = PermViewTarget | PermChangeTarget
const PermOrigin = PermViewOrigin | PermChangeOrigin

const PermView = PermViewName | PermViewData | PermCreated | PermUser | PermViewAbstract | PermViewSpecial | PermViewTarget | PermViewOrigin | PermViewObject | PermViewSpace
const PermChange = PermChangeName | PermChangeData | PermChangeAbstract | PermChangeSpecial | PermChangeTarget | PermChangeOrigin | PermChangeObject | PermChangeSpace
const PermAll = PermDeleteObject | PermView | PermChange

const MaxDuration = (1<<31 - 1) * time.Second

type Permission uint32

const PermissionSize = 4

type PermTriple [3]Permission

type Permissions map[ID]PermTriple

func ParsePermission(str string) Permission {
	p, _ := strconv.ParseUint(str, 16, 32)
	return Permission(p)
}

func (p Permission) String() string {
	return fmt.Sprintf("%08s", strconv.FormatUint(uint64(p), 16))
}

func (p Permission) Bytes() []byte {
	var bts [4]byte
	binary.LittleEndian.PutUint32(bts[:], uint32(p))
	return bts[:]
}

func (t PermTriple) Equal(other PermTriple) bool {
	return t[0] == other[0] && t[1] == other[1] && t[2] == other[2]
}

func (t PermTriple) String() string {
	return t[0].String() + t[1].String() + t[2].String()
}

func (t PermTriple) Bytes() []byte {
	b := bytes.Buffer{}
	b.Write(t[0].Bytes())
	b.Write(t[1].Bytes())
	b.Write(t[2].Bytes())
	return b.Bytes()
}

func ParsePermTriple(str string) PermTriple {
	if len(str) != 3*2*PermissionSize {
		return PermTriple{}
	}
	p1 := ParsePermission(str[:2*PermissionSize])
	p2 := ParsePermission(str[2*PermissionSize : 4*PermissionSize])
	p3 := ParsePermission(str[4*PermissionSize:])
	return PermTriple{p1, p2, p3}
}

func (t PermTriple) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("")
	buffer.WriteString("\"")
	buffer.WriteString(t.String())
	buffer.WriteString("\"")
	return buffer.Bytes(), nil
}

func (t *PermTriple) UnmarshalJSON(b []byte) error {
	var tripleString string
	err := json.Unmarshal(b, &tripleString)
	if err != nil {
		return err
	}
	*t = ParsePermTriple(tripleString)
	return nil
}

func (t Permissions) MarshalJSON() ([]byte, error) {
	strMp := map[string]string{}
	for k, v := range t {
		strMp[k.String()] = v.String()
	}
	return json.Marshal(strMp)
}

func (t *Permissions) UnmarshalJSON(b []byte) error {
	var strMp map[string]string
	err := json.Unmarshal(b, &strMp)
	if err != nil {
		return err
	}
	*t = Permissions{}
	for k, v := range strMp {
		id, _ := ParseID(k)
		triple := ParsePermTriple(v)
		(*t)[id] = triple
	}
	return nil
}
