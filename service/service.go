package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Bitspark/go-vyze/core"
	"github.com/Bitspark/go-vyze/system"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ServiceClient struct {
	endpoint string
	token    string
}

func NewServiceClient(endpoint string) *ServiceClient {
	c := &ServiceClient{
		endpoint: endpoint,
	}
	return c
}

func (sc *ServiceClient) SetToken(token string) {
	sc.token = token
}

type ServiceInfo struct {
	// APIVersion represents the API version
	APIVersion string `json:"apiVersion"`

	// ServerVersion represents the version of the server
	ServerVersion string `json:"serverVersion"`

	// UnixTime contains the server time
	UnixTime int64 `json:"unixTime"`

	// UserID contains the user core.ID of the provided token
	UserID core.ID `json:"user"`

	// Permissions contains the permissions of the provided token
	Permissions []string `json:"permissions"`
}

func (sc *ServiceClient) Info() (ServiceInfo, error) {
	u := sc.buildURL(fmt.Sprintf("info"), nil)
	infoBytes, err := sc.request(u, "GET", nil)
	if err != nil {
		return ServiceInfo{}, err
	}
	var info ServiceInfo
	if err := json.Unmarshal(infoBytes, &info); err != nil {
		return ServiceInfo{}, err
	}
	return info, nil
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	ID   core.ID `json:"id"`
	Name string  `json:"name"`
}

type LoginResponse struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}

func (sc *ServiceClient) ResolveUniverse(name string) (core.ID, error) {
	u := sc.buildURL(fmt.Sprintf("resolve/universe"), map[string][]string{"i": {name}})
	universeIDBytes, err := sc.request(u, "GET", nil)
	if err != nil {
		return core.ID{}, err
	}
	var universeID core.ID
	if err := json.Unmarshal(universeIDBytes, &universeID); err != nil {
		return core.ID{}, err
	}
	if universeID.IsNull() {
		return core.ID{}, errors.New("universe not found")
	}
	return universeID, nil
}

func (sc *ServiceClient) LoadUniverse(id core.ID) (system.Universe, error) {
	u := sc.buildURL(fmt.Sprintf("universe/%s/export", id.Hex()), map[string][]string{"o": {"1"}})
	universeBytes, err := sc.request(u, "GET", nil)
	if err != nil {
		return system.Universe{}, err
	}
	univ := system.Universe{}
	universeReader := bytes.NewReader(universeBytes)
	if err := univ.Load(universeReader, false); err != nil {
		return system.Universe{}, err
	}
	return univ, nil
}

type UpdateUniverseServiceRequest struct {
	Description *string `json:"description,omitempty"`
	Public      *bool   `json:"public,omitempty"`
	Extensible  *bool   `json:"extensible,omitempty"`
}

func (sc *ServiceClient) UpdateUniverse(universeID core.ID, description *string, public *bool, extensible *bool) error {
	req := UpdateUniverseServiceRequest{}
	req.Description = description
	req.Public = public
	req.Extensible = extensible
	u := sc.buildURL(fmt.Sprintf("universe/%s", universeID.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	_, err := sc.request(u, "PUT", reqBytes)
	return err
}

type CreateModelServiceRequest struct {
	Name        string  `json:"name"`
	Object      core.ID `json:"object"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
}

func (sc *ServiceClient) CreateModel(universeID core.ID, name string, objectID core.ID, modelType string, description string) error {
	req := CreateModelServiceRequest{
		Name:        name,
		Object:      objectID,
		Type:        modelType,
		Description: description,
	}
	u := sc.buildURL(fmt.Sprintf("universe/%s/model", universeID.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	_, err := sc.request(u, "POST", reqBytes)
	return err
}

func (sc *ServiceClient) DeleteModel(universeID core.ID, name string) error {
	u := sc.buildURL(fmt.Sprintf("universe/%s/model", universeID.Hex()), map[string][]string{"i": {name}})
	_, err := sc.request(u, "DELETE", nil)
	return err
}

type UpdateModelServiceRequest struct {
	Description *string `json:"description"`
}

func (sc *ServiceClient) UpdateModel(universeID core.ID, name string, description *string) error {
	req := UpdateModelServiceRequest{
		Description: description,
	}
	u := sc.buildURL(fmt.Sprintf("universe/%s/model", universeID.Hex()), map[string][]string{"i": {name}})
	reqBytes, _ := json.Marshal(req)
	_, err := sc.request(u, "PUT", reqBytes)
	return err
}

type CreateEndpointServiceRequest struct {
	Name       string              `json:"name"`
	Definition system.EndpointNode `json:"definition"`
}

func (sc *ServiceClient) CreateEndpoint(universeID core.ID, definition system.EndpointNode) (system.EndpointNode, error) {
	req := CreateEndpointServiceRequest{
		Name:       definition.Name,
		Definition: definition,
	}
	u := sc.buildURL(fmt.Sprintf("universe/%s/endpoint", universeID.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	respBytes, _ := sc.request(u, "POST", reqBytes)
	var resp system.EndpointNode
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return system.EndpointNode{}, err
	}
	return resp, nil
}

type UpdateEndpointServiceRequest struct {
	Name       string              `json:"name"`
	Definition system.EndpointNode `json:"definition"`
}

func (sc *ServiceClient) UpdateEndpoint(universeID core.ID, definition system.EndpointNode) (system.EndpointNode, error) {
	req := UpdateEndpointServiceRequest{
		Name:       definition.Name,
		Definition: definition,
	}
	u := sc.buildURL(fmt.Sprintf("universe/%s/endpoint", universeID.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	respBytes, _ := sc.request(u, "PUT", reqBytes)
	var resp system.EndpointNode
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return system.EndpointNode{}, err
	}
	return resp, nil
}

func (sc *ServiceClient) DeleteEndpoint(universeID core.ID, name string) error {
	u := sc.buildURL(fmt.Sprintf("universe/%s/endpoint", universeID.Hex()), map[string][]string{"i": {name}})
	_, err := sc.request(u, "DELETE", nil)
	return err
}

func (sc *ServiceClient) GetEndpointInterfaces(universeID core.ID, endpointName string) ([]system.Interface, error) {
	u := sc.buildURL(fmt.Sprintf("universe/%s/endpoint/%s/interfaces", universeID.Hex(), endpointName), nil)
	respBytes, err := sc.request(u, "GET", nil)
	if err != nil {
		return nil, err
	}
	var resp []system.Interface
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type CreateInterfaceServiceRequest struct {
	Name       string                `json:"name"`
	Definition system.NamedInterface `json:"definition"`
}

func (sc *ServiceClient) CreateInterface(universeID core.ID, definition system.NamedInterface) (system.NamedInterface, error) {
	req := CreateInterfaceServiceRequest{
		Name:       definition.Name,
		Definition: definition,
	}
	u := sc.buildURL(fmt.Sprintf("universe/%s/interface", universeID.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	respBytes, _ := sc.request(u, "POST", reqBytes)
	var resp system.NamedInterface
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return system.NamedInterface{}, err
	}
	return resp, nil
}

func (sc *ServiceClient) DeleteInterface(universeID core.ID, name string) error {
	u := sc.buildURL(fmt.Sprintf("universe/%s/interface", universeID.Hex()), map[string][]string{"i": {name}})
	_, err := sc.request(u, "DELETE", nil)
	return err
}

func (sc *ServiceClient) GetUniverseJSONSchema(universeName string) (string, error) {
	u := sc.buildURL(fmt.Sprintf("json/%s", universeName), nil)
	schemaBytes, err := sc.request(u, "GET", nil)
	if err != nil {
		return "", err
	}
	return string(schemaBytes), nil
}

func (sc *ServiceClient) GetUniverseCode(universeName string, language string) (string, error) {
	u := sc.buildURL(fmt.Sprintf("code/%s", universeName), map[string][]string{"l": {language}})
	schemaBytes, err := sc.request(u, "GET", nil)
	if err != nil {
		return "", err
	}
	return string(schemaBytes), nil
}

func (sc *ServiceClient) GetLayerProfile(profileID core.ID) (system.LayerProfile, error) {
	u := sc.buildURL(fmt.Sprintf("profile/%s/tokens", profileID.Hex()), nil)
	respBytes, err := sc.request(u, "GET", nil)
	if err != nil {
		return nil, err
	}
	return system.ReadLayerProfile(string(respBytes))
}

func (sc *ServiceClient) GetComponents(universe core.ID, componentType ComponentType, subType string) ([]Component, error) {
	params := map[string][]string{
		"type":    {string(componentType)},
		"subType": {subType},
	}
	if !universe.IsNull() {
		params["universe"] = []string{universe.Hex()}
	}
	u := sc.buildURL("app/component", params)
	respBytes, _ := sc.request(u, "GET", nil)
	var resp []Component
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (sc *ServiceClient) GetComponent(componentID core.ID) (Component, error) {
	u := sc.buildURL(fmt.Sprintf("app/component/%s", componentID.Hex()), nil)
	respBytes, _ := sc.request(u, "GET", nil)
	var resp Component
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return Component{}, err
	}
	return resp, nil
}

type CreateComponentRequest struct {
	Name     string        `json:"name"`
	Type     ComponentType `json:"type"`
	SubType  string        `json:"subType"`
	Universe core.ID       `json:"universe"`
}

func (sc *ServiceClient) CreateComponent(name string, componentType ComponentType, subType string, universe core.ID) (Component, error) {
	req := CreateComponentRequest{
		Name:     name,
		Type:     componentType,
		SubType:  subType,
		Universe: universe,
	}
	u := sc.buildURL("app/component", nil)
	reqBytes, _ := json.Marshal(req)
	respBytes, _ := sc.request(u, "POST", reqBytes)
	var resp Component
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return Component{}, err
	}
	return resp, nil
}

type UpdateComponentRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Public      *bool    `json:"public,omitempty"`
	Listed      *bool    `json:"listed,omitempty"`
	Profile     *core.ID `json:"profile,omitempty"`
	Title       *string  `json:"title,omitempty"`
}

func (sc *ServiceClient) UpdateComponent(id core.ID, name *string, description *string, public *bool, listed *bool) (Component, error) {
	req := UpdateComponentRequest{
		Name:        name,
		Description: description,
		Public:      public,
		Listed:      listed,
	}
	u := sc.buildURL(fmt.Sprintf("app/component/%s", id.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	respBytes, _ := sc.request(u, "PUT", reqBytes)
	var resp Component
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return Component{}, err
	}
	return resp, nil
}

func (sc *ServiceClient) DeleteComponent(id core.ID) error {
	u := sc.buildURL(fmt.Sprintf("app/component/%s", id.Hex()), nil)
	_, err := sc.request(u, "DELETE", nil)
	return err
}

func (sc *ServiceClient) GetWebView(componentID core.ID) (WebViewComponent, error) {
	u := sc.buildURL(fmt.Sprintf("app/component/view/web/%s", componentID.Hex()), nil)
	respBytes, _ := sc.request(u, "GET", nil)
	var resp WebViewComponent
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return WebViewComponent{}, err
	}
	return resp, nil
}

type UpdateViewRequest struct {
	Model *string `json:"model,omitempty"`
	List  *bool   `json:"list,omitempty"`
}

type UpdateWebViewRequest struct {
	Tag *string `json:"tag,omitempty"`
}

func (sc *ServiceClient) UpdateWebView(componentID core.ID, tag *string) (WebViewComponent, error) {
	req := UpdateWebViewRequest{
		Tag: tag,
	}
	u := sc.buildURL(fmt.Sprintf("app/component/view/web/%s", componentID.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	respBytes, _ := sc.request(u, "PUT", reqBytes)
	var resp WebViewComponent
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return WebViewComponent{}, err
	}
	return resp, nil
}

func (sc *ServiceClient) GetFlexView(componentID core.ID) (FlexViewComponent, error) {
	u := sc.buildURL(fmt.Sprintf("app/component/view/flex/%s", componentID.Hex()), nil)
	respBytes, _ := sc.request(u, "GET", nil)
	var resp FlexViewComponent
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return FlexViewComponent{}, err
	}
	return resp, nil
}

type UpdateFlexViewRequest struct {
	Definition *string `json:"definition"`
}

func (sc *ServiceClient) UpdateFlexView(componentID core.ID, definition *string) (FlexViewComponent, error) {
	req := UpdateFlexViewRequest{
		Definition: definition,
	}
	u := sc.buildURL(fmt.Sprintf("app/component/view/flex/%s", componentID.Hex()), nil)
	reqBytes, _ := json.Marshal(req)
	respBytes, _ := sc.request(u, "PUT", reqBytes)
	var resp FlexViewComponent
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return FlexViewComponent{}, err
	}
	return resp, nil
}

type CreateAppRequest struct {
	Name     string  `json:"name"`
	Universe core.ID `json:"universe"`
	Public   bool    `json:"public"`
}

type UpdateAppRequest struct {
	Name        *string `json:"name,omitempty"`
	Public      *bool   `json:"public,omitempty"`
	Description *string `json:"description,omitempty"`
	Path        *string `json:"path,omitempty"`
	Title       *string `json:"title,omitempty"`
}

type CreateRouteRequest struct {
	Route string  `json:"route"`
	View  core.ID `json:"view"`
}

func (sc *ServiceClient) buildURL(endpoint string, params url.Values) string {
	var paramStr string
	if params != nil {
		paramStr = params.Encode()
	}
	return fmt.Sprintf("%s/v1/%s?%s", sc.endpoint, endpoint, paramStr)
}

func (sc *ServiceClient) request(u string, method string, body []byte) ([]byte, error) {
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, u, bodyReader)
	if err != nil {
		return nil, err
	}
	if len(sc.token) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sc.token))
	}
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
	if system.IsError(system.Code(code)) {
		errMessage := resp.Header.Get("x-vy-message")
		return nil, errors.New(errMessage)
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("response: %d", resp.StatusCode)
	}
	return respBytes, nil
}
