package gosample_test

import (
	"gosample/gosample"
	"testing"
)

func TestSimpleObj(t *testing.T) {
	obj := gosample.SimpleObj{ID: 33, Value: "sdfs"}

	obj.SetValue("prova")
	if obj.Value == "prova" {
		t.Errorf("sbagliato")
	}
}
