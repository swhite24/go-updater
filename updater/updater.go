// Package updater provides a mechanism overwrite struct data with a map.
package updater

import (
	"reflect"
	"strings"
)

// Struct takes a struct model and an update map and updates the values
// on the struct marked with the `update` struct tag.
func Struct(model interface{}, update map[string]interface{}) {
	// Get type / value
	t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem()

	// Ensure struct
	if v.Kind() != reflect.Struct {
		return
	}

	// Iterate over struct fields
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)

		if f.Tag.Get("update") != "" {
			key := getKeyName(f)
			switch f.Type.Kind() {
			case reflect.String:
				if u, ok := update[key].(string); ok {
					fv.SetString(u)
				}
			case reflect.Bool:
				if u, ok := update[key].(bool); ok {
					fv.SetBool(u)
				}
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
				if u, ok := update[key].(int); ok {
					fv.SetInt(int64(u))
				} else if u, ok := update[key].(int64); ok {
					fv.SetInt(u)
				} else if u, ok := update[key].(float64); ok {
					fv.SetInt(int64(u))
				}
			case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
				if u, ok := update[key].(uint64); ok {
					fv.SetUint(u)
				} else if u, ok := update[key].(float64); ok {
					fv.SetUint(uint64(u))
				}
			case reflect.Float32, reflect.Float64:
				if u, ok := update[key].(float64); ok {
					fv.SetFloat(u)
				}
			case reflect.Map:
				if m, ok := update[key].(map[string]interface{}); ok {
					fv.Set(reflect.ValueOf(m))
				}
			case reflect.Slice:
				if u, ok := update[key].([]string); ok {
					fv.Set(reflect.ValueOf(u))
				} else if u, ok := update[key].([]int); ok {
					fv.Set(reflect.ValueOf(u))
				} else if u, ok := update[key].([]interface{}); ok {
					handleSliceInterface(fv, u)
				}
			}
		}
	}
}

func handleSliceInterface(fv reflect.Value, s []interface{}) {
	if len(s) == 0 {
		return
	}

	first := s[0]
	switch first.(type) {
	case string:
		update := []string{}
		for _, v := range s {
			if u, ok := v.(string); ok {
				update = append(update, u)
			}
		}
		fv.Set(reflect.ValueOf(update))
	case int:
		update := []int{}
		for _, v := range s {
			if u, ok := v.(int); ok {
				update = append(update, u)
			}
		}
		fv.Set(reflect.ValueOf(update))
	}
}

func getKeyName(f reflect.StructField) string {
	// Default to bson, otherwise json
	key := f.Tag.Get("json")
	if key == "" {
		key = f.Name
	}
	return strings.Split(key, ",")[0]
}
