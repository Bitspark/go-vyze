package core

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"hash/maphash"
	"math"
	"time"
)

type Binary []byte

type Key [32]byte

func encodeLength(l int64) []byte {
	b := []byte{}
	for {
		nb := byte(l % 128)
		l /= 128
		if l == 0 {
			b = append(b, nb)
			break
		} else {
			nb |= 0b10000000
			b = append(b, nb)
		}
	}
	return b
}

func decodeLength(b []byte) (int64, int) {
	l := int64(0)
	m := int64(1)
	for i := 0; i < len(b); i++ {
		l += m * int64(0b01111111&b[i])
		m *= 128
		if 0b10000000&b[i] != 0b10000000 {
			return l, i + 1
		}
	}
	return 0, 0
}

func Encrypt(d []byte, key Key) []byte {
	dl := len(d)
	b := encodeLength(int64(dl))
	el := len(b) + dl
	tl := ((el-1)/aes.BlockSize + 1) * aes.BlockSize
	e := make([]byte, tl)
	copy(e, b)
	copy(e[len(b):], d)
	c, _ := aes.NewCipher(key[:])
	for i := 0; i < tl/aes.BlockSize; i++ {
		c.Encrypt(e[i*aes.BlockSize:], e[i*aes.BlockSize:])
	}
	return e
}

func (d *Binary) Encrypt(key Key) {
	*d = Encrypt(*d, key)
}

func Decrypt(d []byte, key Key) ([]byte, error) {
	c, _ := aes.NewCipher(key[:])
	tl := len(d)
	e := make([]byte, tl)
	for i := 0; i < tl/aes.BlockSize; i++ {
		c.Decrypt(e[i*aes.BlockSize:], d[i*aes.BlockSize:])
	}
	dl, n := decodeLength(e)
	if n+int(dl) > len(e) {
		return nil, errors.New("error decoding")
	}
	return e[n : n+int(dl)], nil
}

func (d *Binary) Decrypt(key Key) error {
	var err error
	*d, err = Decrypt(*d, key)
	return err
}

func (d Binary) MarshalJSON() ([]byte, error) {
	if d == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(d.String())
}

func (d *Binary) UnmarshalJSON(b []byte) error {
	var str *string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	if str == nil {
		*d = nil
		return nil
	}
	*d = ParseBinary(*str)
	return nil
}

func (d Binary) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

func (d *Binary) UnmarshalYAML(value *yaml.Node) error {
	var str string
	err := value.Decode(&str)
	if err != nil {
		return err
	}
	*d = ParseBinary(str)
	return nil
}

func (d Binary) Equals(data Binary) bool {
	if (d == nil) != (data == nil) {
		return false
	}
	return bytes.Compare(d[:], data[:]) == 0
}

func (d Binary) String() string {
	if d == nil {
		return "nil"
	}
	return d.Base64()
}

func (d Binary) Hex() string {
	return hex.EncodeToString(d[:])
}

func (d *Binary) Base64() string {
	if d == nil {
		return ""
	}
	return base64.RawStdEncoding.EncodeToString(*d)
}

func (d Binary) Hash64() uint64 {
	var h maphash.Hash
	_, _ = h.Write(d)
	return h.Sum64()
}

func (d Binary) ToID() ID {
	id := ID{}
	copy(id[:], d)
	return id
}

func (d Binary) ToRID() RichID {
	rid := RichID{}
	copy(rid[:], d)
	return rid
}

func ParseBinary(h string) Binary {
	return ParseBinaryBase64(h)
}

func ParseBinaryHex(h string) Binary {
	b, _ := hex.DecodeString(h)
	return b
}

func ParseBinaryBase64(s string) Binary {
	d, _ := base64.RawStdEncoding.DecodeString(s)
	return d
}

// Converters

func (d Binary) ToString() (string, error) {
	return string(d), nil
}

func (d *Binary) FromString(s string) error {
	*d = Binary(s)
	return nil
}

func BinaryFromString(s string) Binary {
	d := Binary{}
	if err := d.FromString(s); err != nil {
		return nil
	}
	return d
}

func (d Binary) ToInt64() (int64, error) {
	if len(d) > 9 {
		return 0, errors.New("too many bytes")
	}
	if len(d) == 0 {
		return 0, nil
	}
	if len(d) == 1 {
		return 0, errors.New("need at least 2 bytes")
	}
	i := int64(0)
	for k := 1; k < len(d); k++ {
		i += int64(d[k]) << ((k - 1) * 8)
	}
	if d[0] == 0 {
		return i, nil
	} else if d[0] == 0xFF {
		return -i, nil
	} else {
		return 0, errors.New("wrong sign byte")
	}
}

func (d *Binary) FromInt64(i int64) error {
	*d = (*d)[:0]

	if i == 0 {
		return nil
	}

	remainder := i
	if remainder > 0 {
		*d = append(*d, 0)
	} else {
		*d = append(*d, 255)
		remainder = -remainder
	}
	for remainder > 0 {
		*d = append(*d, byte(remainder%256))
		remainder /= 256
	}

	return nil
}

func BinaryFromInt(i int64) Binary {
	d := Binary{}
	if err := d.FromInt64(i); err != nil {
		return nil
	}
	return d
}

func (d Binary) ToUint() (uint64, error) {
	if len(d) > 8 {
		return 0, errors.New("too many bytes")
	}
	if len(d) == 0 {
		return 0, nil
	}
	i := uint64(0)
	for k := 0; k < len(d); k++ {
		i += uint64(d[k]) << (k * 8)
	}
	return i, nil
}

func (d *Binary) FromUint(i uint64) error {
	*d = (*d)[:0]

	if i == 0 {
		return nil
	}

	remainder := i
	for remainder > 0 {
		*d = append(*d, byte(remainder%256))
		remainder /= 256
	}

	return nil
}

func BinaryFromUint(i uint64) Binary {
	d := Binary{}
	if err := d.FromUint(i); err != nil {
		return nil
	}
	return d
}

func (d Binary) ToFloat64() (float64, error) {
	if len(d) == 0 {
		return 0, nil
	}
	bits, err := d.ToUint()
	if err != nil {
		return 0, err
	}
	f := math.Float64frombits(bits)
	return f, nil
}

func (d *Binary) FromFloat64(f float64) error {
	if f == 0 {
		*d = (*d)[:0]
		return nil
	}
	return d.FromUint(math.Float64bits(f))
}

func BinaryFromFloat64(f float64) Binary {
	d := Binary{}
	if err := d.FromFloat64(f); err != nil {
		return nil
	}
	return d
}

func (d Binary) ToBoolean() (bool, error) {
	return len(d) > 0, nil
}

func (d *Binary) FromBoolean(b bool) error {
	if b {
		*d = []byte{255}
	} else {
		*d = []byte{}
	}
	return nil
}

func BinaryFromBoolean(b bool) Binary {
	d := Binary{}
	if err := d.FromBoolean(b); err != nil {
		return nil
	}
	return d
}

func (d Binary) ToTime() (time.Time, error) {
	t := time.Time{}
	err := t.UnmarshalBinary(d)
	return t, err
}

func (d *Binary) FromTime(date time.Time) error {
	var err error
	*d, err = date.MarshalBinary()
	return err
}

func BinaryFromTime(date time.Time) Binary {
	d := Binary{}
	if err := d.FromTime(date); err != nil {
		return nil
	}
	return d
}
