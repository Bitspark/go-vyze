package system

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bitspark/go-vyze/core"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Consts

type Code int

const (
	// CodeUndefined should not be used and will most likely trigger an error
	CodeUndefined Code = iota

	// CodeSuccess200 represents a 200 (OK) status
	CodeSuccess200 Code = iota

	// CodePartial206 represents a 206 (partial content) status
	CodePartial206 Code = iota

	// CodeError403 represents a 403 (authentication required) error
	CodeError403 Code = iota

	// CodeError400BadRequest represents a 400 (bad request) error
	CodeError400BadRequest Code = iota

	// CodeError404NotFound represents a 404 (not found) error
	CodeError404NotFound Code = iota

	// CodeError416InvalidRange represents a 416 (range not satisfiable) error
	CodeError416InvalidRange Code = iota

	// CodeError500OrUnknown represents a server-side error or an unknown error
	CodeError500OrUnknown Code = iota
)

func IsError(code Code) bool {
	return code != CodeSuccess200 && code != CodePartial206
}

// Structs

type LayerInfo struct {
	// LayerID represents the layer of this info
	LayerID core.ID `json:"layerId"`

	// UserID represents the user for whom this token was issues
	UserID core.ID `json:"userId"`

	// Permissions contains granted, mandatory and exclusive permissions for this layer
	Permissions string `json:"permissions"`
}

type Info struct {
	// APIVersion represents the API version
	APIVersion string `json:"apiVersion"`

	// ServerVersion represents the version of the server
	ServerVersion string `json:"serverVersion"`

	// UnixTime contains the server time
	UnixTime int64 `json:"unixTime"`

	// Layers contains layer info for the connecting client
	Layers []LayerInfo `json:"layers"`
}

type LayerTokenInfo struct {
	// LayerID represents the layer of this token
	LayerID core.ID `json:"layer"`

	// UserID represents the user for whom this token was issued
	UserID core.ID `json:"user"`

	// Granted contains the granted permissions
	Granted uint32 `json:"granted"`

	// Mandatory contains the mandatory permissions
	Mandatory uint32 `json:"mandatory"`

	// Exclusive contains the exclusive permissions
	Exclusive uint32 `json:"exclusive"`

	// Created contains the creation date
	Created int64 `json:"created"`

	// Expiry contains the duration of validity
	Expiry int64 `json:"expiry"`

	// Admin indicates whether this user is an admin
	Admin bool `json:"admin"`
}

type Object struct {
	ID      core.ID `json:"id"`
	Name    string  `json:"name"`
	Created int64   `json:"created"`
	User    core.ID `json:"user"`
}

// Requests

type CreateObjectRequest struct {
	Abstracts core.IDSet `json:"abstracts"`
	Name      string     `json:"name"`
	Dependent bool       `json:"dependent"`
	Layer     core.ID    `json:"layer"`
}

type CreateRelationRequest struct {
	Name      string     `json:"name,omitempty"`
	Origin    core.ID    `json:"origin"`
	Target    core.ID    `json:"target"`
	Abstracts core.IDSet `json:"abstracts"`
	Layer     core.ID    `json:"layer"`
	Key       any        `json:"key,omitempty"`
	KeyFormat FormatType `json:"keyFormat,omitempty"`
}

type CreateLayerRequest struct {
	Abstracts core.IDSet `json:"abstracts"`
	Name      string     `json:"name"`
}

type PutResourceRequest struct {
	ObjectID core.ID        `json:"objectId"`
	Object   ResourceType   `json:"object"`
	Schema   ResourceSchema `json:"schema"`
	Value    any            `json:"value"`
	Layer    core.ID        `json:"layer"`
}

type GetResourceRequest struct {
	ObjectID core.ID         `json:"objectId"`
	Object   ResourceType    `json:"object"`
	Schema   ResourceSchema  `json:"schema"`
	Filter   ResourceFilters `json:"filter,omitempty"`
	Order    ResourceOrders  `json:"order,omitempty"`
	Offset   *int            `json:"offset,omitempty"`
	Limit    *int            `json:"limit,omitempty"`
}

// Client

type AccessOptions struct {
	Access      string
	AccessNames []string `json:"accessNames"`
}

type SystemClient struct {
	endpoint       string
	layerProfile   *LayerProfile
	defaultOptions *AccessOptions
}

func NewSystemClient(endpoint string) *SystemClient {
	c := &SystemClient{
		endpoint:     endpoint,
		layerProfile: &LayerProfile{},
		defaultOptions: &AccessOptions{
			Access:      "main_full",
			AccessNames: []string{"main_read", "main_full"},
		},
	}
	return c
}

func (c *SystemClient) SetDefaultOptions(options *AccessOptions) {
	c.defaultOptions = options
}

func (c *SystemClient) LayerProfile() *LayerProfile {
	return c.layerProfile
}

func (c *SystemClient) SetLayerProfile(layerProfile LayerProfile) {
	c.layerProfile = &layerProfile
}

func (c *SystemClient) GetInfo(options *AccessOptions) (Info, error) {
	u := c.buildURL("info", nil)
	info := Info{}
	if err := c.getJSON(u, &info, options); err != nil {
		return Info{}, err
	}
	return info, nil
}

func (c *SystemClient) VerifyToken(token string) (LayerTokenInfo, error) {
	u := c.buildURL("token", map[string][]string{"token": {token}})
	layerToken := LayerTokenInfo{}
	if err := c.getJSON(u, &layerToken, nil); err != nil {
		return LayerTokenInfo{}, err
	}
	return layerToken, nil
}

func (c *SystemClient) GetObject(id core.ID, options *AccessOptions) (Object, error) {
	u := c.buildURL(fmt.Sprintf("object/%s", id.Hex()), nil)
	obj := Object{}
	if err := c.getJSON(u, &obj, options); err != nil {
		return Object{}, err
	}
	return obj, nil
}

func (c *SystemClient) CreateObject(abstractIDs core.IDSet, name string, accessName string, options *AccessOptions) (Object, error) {
	u := c.buildURL(fmt.Sprintf("object"), nil)
	req := CreateObjectRequest{
		Abstracts: abstractIDs,
		Name:      name,
		Layer:     c.getAccessLayerID(accessName),
	}
	obj := Object{}
	err := c.postJSON(u, req, &obj, options)
	if err != nil {
		return Object{}, err
	}
	return obj, nil
}

func (c *SystemClient) DeleteObject(id core.ID, options *AccessOptions) error {
	u := c.buildURL(fmt.Sprintf("object/%s", id.Hex()), nil)
	if _, err := c.deleteRaw(u, options); err != nil {
		return err
	}
	return nil
}

func (c *SystemClient) GetData(id core.ID, formatType FormatType, options *AccessOptions) (any, error) {
	u := c.buildURL(fmt.Sprintf("object/%s/data", id.Hex()), map[string][]string{"format": {string(formatType)}})
	var err error
	var obj any
	switch formatType {
	case FormatTypeString, FormatTypeHex, FormatTypeBase64:
		objStr := string("")
		err = c.getJSON(u, &objStr, options)
		obj = objStr
	case FormatTypeInteger:
		objInt := int64(0)
		err = c.getJSON(u, &objInt, options)
		obj = objInt
	case FormatTypeFloat:
		objFloat := float64(0)
		err = c.getJSON(u, &objFloat, options)
		obj = objFloat
	case FormatTypeBoolean:
		objBool := bool(false)
		err = c.getJSON(u, &objBool, options)
		obj = objBool
	case FormatTypeRaw:
		obj, err = c.getRaw(u, options)
	}
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *SystemClient) SetData(id core.ID, value any, formatType FormatType, options *AccessOptions) error {
	u := c.buildURL(fmt.Sprintf("object/%s/data", id.Hex()), map[string][]string{"format": {string(formatType)}})
	var err error
	var resp interface{}
	switch formatType {
	case FormatTypeString, FormatTypeHex, FormatTypeBase64:
		err = c.postJSON(u, value.(string), &resp, options)
	case FormatTypeInteger:
		err = c.postJSON(u, value.(int64), &resp, options)
	case FormatTypeFloat:
		err = c.postJSON(u, value.(float64), &resp, options)
	case FormatTypeBoolean:
		err = c.postJSON(u, value.(bool), &resp, options)
	case FormatTypeRaw:
		resp, err = c.postRaw(u, value.(core.Binary), options)
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *SystemClient) DeleteData(id core.ID, options *AccessOptions) error {
	u := c.buildURL(fmt.Sprintf("object/%s/data", id.Hex()), nil)
	var resp any
	if err := c.deleteJSON(u, &resp, options); err != nil {
		return err
	}
	return nil
}

func (c *SystemClient) GetName(id core.ID, options *AccessOptions) (string, error) {
	u := c.buildURL(fmt.Sprintf("object/%s/name", id.Hex()), nil)
	var name string
	if err := c.getJSON(u, &name, options); err != nil {
		return "", err
	}
	return name, nil
}

func (c *SystemClient) SetName(id core.ID, name string, options *AccessOptions) error {
	u := c.buildURL(fmt.Sprintf("object/%s/name", id.Hex()), nil)
	err := c.postJSON(u, name, nil, options)
	return err
}

func (c *SystemClient) GetAbstracts(id core.ID, self bool, direct bool, indirect bool, options *AccessOptions) (core.IDSet, error) {
	var selfStr, directStr, indirectStr string
	if self {
		selfStr = "1"
	}
	if direct {
		directStr = "1"
	}
	if indirect {
		indirectStr = "1"
	}
	u := c.buildURL(fmt.Sprintf("object/%s/abstracts", id.Hex()), map[string][]string{"self": {selfStr}, "direct": {directStr}, "transitive": {indirectStr}})
	var ids core.IDSet
	if err := c.getJSON(u, &ids, options); err != nil {
		return nil, err
	}
	return ids, nil
}

func (c *SystemClient) GetSpecials(id core.ID, self bool, direct bool, indirect bool, options *AccessOptions) (core.IDSet, error) {
	var selfStr, directStr, indirectStr string
	if self {
		selfStr = "1"
	}
	if direct {
		directStr = "1"
	}
	if indirect {
		indirectStr = "1"
	}
	u := c.buildURL(fmt.Sprintf("object/%s/specials", id.Hex()), map[string][]string{"self": {selfStr}, "direct": {directStr}, "transitive": {indirectStr}})
	var ids core.IDSet
	if err := c.getJSON(u, &ids, options); err != nil {
		return nil, err
	}
	return ids, nil
}

func (c *SystemClient) CreateLayer(abstractIDs core.IDSet, name string, options *AccessOptions) (Object, error) {
	u := c.buildURL(fmt.Sprintf("layer"), nil)
	req := CreateLayerRequest{
		Abstracts: abstractIDs,
		Name:      name,
	}
	obj := Object{}
	err := c.postJSON(u, req, &obj, options)
	if err != nil {
		return Object{}, err
	}
	return obj, nil
}

func (c *SystemClient) CreateRelation(originID core.ID, targetID core.ID, abstractIDs core.IDSet, name string, accessName string, options *AccessOptions) (Object, error) {
	u := c.buildURL(fmt.Sprintf("relation"), nil)
	req := CreateRelationRequest{
		Name:      name,
		Origin:    originID,
		Target:    targetID,
		Abstracts: abstractIDs,
		Layer:     c.getAccessLayerID(accessName),
	}
	obj := Object{}
	err := c.postJSON(u, req, &obj, options)
	if err != nil {
		return Object{}, err
	}
	return obj, nil
}

func (c *SystemClient) CreateKeyedRelation(originID core.ID, targetID core.ID, abstractIDs core.IDSet, name string, key any, keyFormat FormatType, accessName string, options *AccessOptions) (Object, error) {
	u := c.buildURL(fmt.Sprintf("relation"), nil)
	req := CreateRelationRequest{
		Name:      name,
		Origin:    originID,
		Target:    targetID,
		Abstracts: abstractIDs,
		Layer:     c.getAccessLayerID(accessName),
		Key:       key,
		KeyFormat: keyFormat,
	}
	obj := Object{}
	err := c.postJSON(u, req, &obj, options)
	if err != nil {
		return Object{}, err
	}
	return obj, nil
}

// GetResourceInstance returns a value specified by the resource parameter.
func (c *SystemClient) GetResourceInstance(resource *ResourceInstance, options *AccessOptions) (any, error) {
	u := c.buildURL("resource/get", nil)
	obj := map[string]any{}
	err := c.postJSON(u, resource.toGetRequest(), &obj, options)
	return obj, err
}

// GetResourceList returns a slice of values specified by the resource parameter.
func (c *SystemClient) GetResourceList(resource *ResourceSpecials, options *AccessOptions) ([]any, error) {
	u := c.buildURL("resource/get", nil)
	obj := []any{}
	err := c.postJSON(u, resource.toGetRequest(), &obj, options)
	return obj, err
}

func (c *SystemClient) PutResourceInstance(resource *ResourceInstance, value any, accessName string, options *AccessOptions) (core.ID, error) {
	u := c.buildURL(fmt.Sprintf("resource/put"), nil)
	var resp core.ID
	err := c.postJSON(u, resource.toPutRequest(value, c.getAccessLayerID(accessName)), &resp, options)
	if err != nil {
		return core.ID{}, err
	}
	return resp, nil
}

func (c *SystemClient) PutResourceSpecials(resource *ResourceSpecials, values any, accessName string, options *AccessOptions) (core.IDList, error) {
	u := c.buildURL(fmt.Sprintf("resource/put"), nil)
	var resp core.IDList
	err := c.postJSON(u, resource.toPutRequest(values, c.getAccessLayerID(accessName)), &resp, options)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *SystemClient) GetNode(node Node, options *AccessOptions) (any, error) {
	u := c.buildURL("node/get", nil)
	var resp any
	err := c.postJSON(u, GetNodeRequest{
		Node: node,
	}, &resp, options)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *SystemClient) PutNode(node Node, value any, accessName string, options *AccessOptions) (any, error) {
	u := c.buildURL("node/put", nil)
	var resp any
	err := c.postJSON(u, PutNodeRequest{
		Node:  node,
		Value: value,
		Layer: c.getAccessLayerID(accessName),
	}, &resp, options)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *SystemClient) getAccessLayerID(accessName string) core.ID {
	if len(accessName) == 0 && c.defaultOptions != nil {
		accessName = c.defaultOptions.Access
	}
	acc := c.layerProfile.GetAccessGroup(accessName)
	if acc == nil {
		return core.ID{}
	}
	if len(acc.Tokens) == 0 {
		return core.ID{}
	}
	return acc.Tokens[0].LayerID
}

func (c *SystemClient) buildURL(endpoint string, params url.Values) string {
	var paramStr string
	if params != nil {
		paramStr = params.Encode()
	}
	return fmt.Sprintf("%s/v1/%s?%s", c.endpoint, endpoint, paramStr)
}

func (c *SystemClient) getJSON(u string, respObj any, options *AccessOptions) error {
	respBytes, err := c.request(u, "GET", nil, options)
	if err != nil {
		return err
	}
	if respObj != nil {
		if err := json.Unmarshal(respBytes, respObj); err != nil {
			return err
		}
	}
	return nil
}

func (c *SystemClient) getRaw(u string, options *AccessOptions) (core.Binary, error) {
	respBytes, err := c.request(u, "GET", nil, options)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func (c *SystemClient) deleteJSON(u string, respObj any, options *AccessOptions) error {
	respBytes, err := c.request(u, "DELETE", nil, options)
	if err != nil {
		return err
	}
	if respObj != nil {
		if err := json.Unmarshal(respBytes, respObj); err != nil {
			return err
		}
	}
	return nil
}

func (c *SystemClient) deleteRaw(u string, options *AccessOptions) (core.Binary, error) {
	respBytes, err := c.request(u, "DELETE", nil, options)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func (c *SystemClient) postJSON(u string, reqObj any, respObj any, options *AccessOptions) error {
	reqBytes, err := json.Marshal(reqObj)
	if err != nil {
		return err
	}
	respBytes, err := c.request(u, "POST", reqBytes, options)
	if err != nil {
		return err
	}
	if respObj != nil {
		if err := json.Unmarshal(respBytes, respObj); err != nil {
			return err
		}
	}
	return nil
}

func (c *SystemClient) postRaw(u string, reqBytes core.Binary, options *AccessOptions) (core.Binary, error) {
	respBytes, err := c.request(u, "POST", reqBytes, options)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func (c *SystemClient) request(u string, method string, body []byte, options *AccessOptions) ([]byte, error) {
	if options == nil {
		options = c.defaultOptions
	}
	accessTokens := []string{}
	if options.AccessNames != nil {
		for _, an := range options.AccessNames {
			ag := c.layerProfile.GetAccessGroup(an)
			if ag == nil {
				continue
			}
			for _, tk := range ag.Tokens {
				accessTokens = append(accessTokens, tk.Token)
			}
		}
	} else {
		for _, ag := range *c.layerProfile {
			for _, tk := range ag.Tokens {
				accessTokens = append(accessTokens, tk.Token)
			}
		}
	}
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, u, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-vy-layers", strings.Join(accessTokens, ","))
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	codeStr := resp.Header.Get("x-vy-code")
	code, _ := strconv.Atoi(codeStr)
	if IsError(Code(code)) {
		errMessage := resp.Header.Get("x-vy-message")
		return nil, errors.New(errMessage)
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("response: %d", resp.StatusCode)
	}
	return respBytes, nil
}
