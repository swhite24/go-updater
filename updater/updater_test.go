package updater

import "testing"

type foo struct {
	String      string   `json:"string" update:"true"`
	Int         int64    `json:"int" update:"true"`
	Float       float64  `json:"float" update:"true"`
	Bool        bool     `json:"bool" update:"true"`
	StringSlice []string `json:"string_slice" update:"true"`
	IntSlice    []int    `json:"int_slice" update:"true"`
	IPWhitelist []string `json:"ip_whitelist" update:"true"`
}

func TestUpdateSlice(t *testing.T) {
	f := &foo{StringSlice: []string{"foo"}}
	u := map[string]interface{}{
		"string_slice": []string{"foo", "bar"},
		"int_slice":    []int{1, 2, 3},
		"ip_whitelist": []interface{}{"0.0.0.0", "10.10.0.0"},
	}

	Struct(f, u)

	if len(f.StringSlice) != 2 {
		t.Errorf("failed to update string slice: want %d, got %d", 2, len(f.StringSlice))
	}
	if len(f.IntSlice) != 3 {
		t.Errorf("failed to update int slice: want %d, got %d", 3, len(f.IntSlice))
	}
	if len(f.IPWhitelist) != 2 {
		t.Errorf("failed to update slice interface: want %d, got %d", 2, len(f.IPWhitelist))
	}
}

func TestUpdateString(t *testing.T) {
	f := &foo{String: "foo"}
	u := map[string]interface{}{
		"string": "foo1",
	}

	Struct(f, u)

	if f.String != "foo1" {
		t.Errorf("failed to update string: want %s, got %s", "foo1", f.String)
	}
}

func TestUpdateInt(t *testing.T) {
	f := &foo{Int: 1}
	u := map[string]interface{}{"int": 10}
	Struct(f, u)
	if f.Int != 10 {
		t.Errorf("failed to update int: want %d, got %d", 10, f.Int)
	}
}

func TestUpdateFloat(t *testing.T) {
	f := &foo{Float: 0.5}
	u := map[string]interface{}{"float": 10.5}
	Struct(f, u)
	if f.Float != 10.5 {
		t.Errorf("failed to update float: want %f, got %f", 10.5, f.Float)
	}
}

func TestUpdateBool(t *testing.T) {
	f := &foo{Bool: true}
	u := map[string]interface{}{"bool": false}
	Struct(f, u)

	if f.Bool {
		t.Error("failed to update bool")
	}
}
