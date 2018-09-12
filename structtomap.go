package kazstructs

import (
	"fmt"
	"reflect"
	"strings"
)

//StructToMapLower turns all field names no lower case
func StructToMapLower(in interface{}, omitEmpty bool) (map[string]interface{}, error) {
	return StructToMap(in, omitEmpty, true)
}

//StructToMap transforms an input structure into a map
func StructToMap(in interface{}, omitEmpty bool, fieldsToLower bool) (map[string]interface{}, error) {
	v := reflect.ValueOf(in)

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Input %T type is not a struct", v)
	}

	m := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		var key string
		nametag := v.Type().Field(i).Tag.Get("name")
		if nametag != "" {
			key = nametag
		} else {
			key = v.Type().Field(i).Name
		}

		if fieldsToLower {
			key = strings.ToLower(key)
		}
		value := v.Field(i)

		if omitEmpty && value.String() == "" {
			continue
		}

		switch value.Interface().(type) {
		case int, int8, int16, int32, int64:
			m[key] = value.Int()
		case string:
			m[key] = value.String()
		case bool:
			m[key] = value.Bool()
		default:
			m[key] = value
		}
	}

	return m, nil
}
