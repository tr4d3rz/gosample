package gosample

import "fmt"

//SimpleObj shows nothing
type SimpleObj struct {
	ID    int
	Value string
}

//SetValue sets a value
func (obj *SimpleObj) SetValue(_val string) {
	obj.Value = _val

	fmt.Printf("Value is %s", _val)
}
