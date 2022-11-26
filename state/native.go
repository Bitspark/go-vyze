package state

type NativeInterface interface {
	Set(name string, value any) error
	Get(name string) (any, error)
}

type nativeInterface struct {
	state State
}

func (ni *nativeInterface) Set(name string, value any) error {
	err := ni.state.SetValue(name, value)
	return err
}

func (ni *nativeInterface) Get(name string) (any, error) {
	val, err := ni.state.GetValue(name)
	if err != nil {
		return nil, err
	}
	return val.Value, nil
}

type ActionNative func(stateInterface NativeInterface) error

type ConditionNative func(stateInterface NativeInterface) (bool, error)

type ExpressionNative func(stateInterface NativeInterface) (any, error)
