package service

import "github.com/Bitspark/go-vyze/core"

type ComponentType string

const (
	ComponentTypeView    ComponentType = "view"
	ComponentTypeService ComponentType = "service"
	ComponentTypeAction  ComponentType = "action"
)

const (
	ViewComponentTypeWeb  = "web"
	ViewComponentTypeFlex = "flex"
)

type Component struct {
	ID          core.ID       `json:"id"`
	Type        ComponentType `json:"type"`
	SubType     string        `json:"subType"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Author      core.ID       `json:"author"`
	Universe    core.ID       `json:"universe"`
	Public      bool          `json:"public"`
	Listed      bool          `json:"listed"`
	Profile     core.ID       `json:"profile"`
}

type ViewComponent struct {
	Component
	Model string `json:"model"`
	List  bool   `json:"list"`
}

type ServiceComponent struct {
	Component
}

type ActionComponent struct {
	Component
}

type WebViewComponent struct {
	ViewComponent
	Tag string `json:"tag"`
}

type FlexViewComponent struct {
	ViewComponent
	Definition string `json:"definition"` // TODO: Replace with actual form definition (TBD)
}
