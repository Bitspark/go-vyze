package state

import (
	"fmt"
	"log"
)

type Expression struct {
	Name string `json:"name" yaml:"name"`

	Value     any           `json:"value" yaml:"value"`
	Variable  string        `json:"variable" yaml:"variable"`
	Reference string        `json:"reference" yaml:"reference"`
	Operation string        `json:"operation" yaml:"operation"`
	Children  []*Expression `json:"children" yaml:"children"`

	Native ExpressionNative `json:"-" yaml:"-"`

	valType  *Type
	compiled bool
}

func (c *Expression) Compile(lib *Library) error {
	if c == nil {
		return nil
	}
	if c.Reference != "" {
		if lib == nil {
			return fmt.Errorf("library not set")
		}
		rc, err := lib.GetExpression(c.Reference)
		if err != nil {
			log.Printf("condition %s: %v", c.Reference, err)
		}
		if !rc.compiled {
			if err := rc.Compile(lib); err != nil {
				return err
			}
		}
		*c = *rc
	}
	for _, cc := range c.Children {
		if err := cc.Compile(lib); err != nil {
			return err
		}
	}
	c.compiled = true
	return nil
}

func (c Expression) MustEval(stt State) any {
	ok, err := c.Eval(stt)
	if err != nil {
		panic(err)
	}
	return ok
}

func (c *Expression) Eval(stt State) (any, error) {
	if c == nil {
		return true, nil
	}

	if !c.compiled && c.Reference != "" {
		return false, fmt.Errorf("condition not compiled: %s", c.Name)
	}

	if c.Value != nil {
		return c.Value, nil
	}

	if c.Variable != "" {
		val, err := stt.GetValue(c.Variable)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	if c.Operation != "" {
		switch c.valType.Leaf {
		case LeafString:
			val := ""
			for _, ch := range c.Children {
				valEl, err := ch.Eval(stt)
				if err != nil {
					return nil, err
				}
				valStr, ok := valEl.(string)
				if !ok {
					return nil, fmt.Errorf("expected string, got %v", valEl)
				}
				switch c.Operation {
				case "+":
					val += valStr
				}
			}
			return val, nil

		case LeafInteger:
			val := 0
			for _, ch := range c.Children {
				valEl, err := ch.Eval(stt)
				if err != nil {
					return nil, err
				}
				valInt, ok := valEl.(int)
				if !ok {
					return nil, fmt.Errorf("expected integer, got %v", valEl)
				}
				switch c.Operation {
				case "+":
					val += valInt
				case "-":
					val -= valInt
				case "*":
					val *= valInt
				case "/":
					val /= valInt
				}
			}
			return val, nil

		case LeafFloat:
			val := 0.0
			for _, ch := range c.Children {
				valEl, err := ch.Eval(stt)
				if err != nil {
					return nil, err
				}
				valFloat, ok := valEl.(float64)
				if !ok {
					return nil, fmt.Errorf("expected float, got %v", valEl)
				}
				switch c.Operation {
				case "+":
					val += valFloat
				case "-":
					val -= valFloat
				case "*":
					val *= valFloat
				case "/":
					val /= valFloat
				}
			}
			return val, nil

		case LeafBoolean:
			val := true
			for _, ch := range c.Children {
				valEl, err := ch.Eval(stt)
				if err != nil {
					return nil, err
				}
				valBool, ok := valEl.(bool)
				if !ok {
					return nil, fmt.Errorf("expected boolean, got %v", valEl)
				}
				switch c.Operation {
				case "&&":
					val = val && valBool
				case "||":
					val = val || valBool
				}
			}
			return val, nil
		}
	}

	return nil, nil
}
