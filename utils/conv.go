package utils

import (
	"errors"
	"reflect"
	"strconv"
)

//interface TO MapInterface
func InterfaceToMapInterface(in interface{}) (map[string]interface{}, error) {
	var R_data = make(map[string]interface{})
	item := reflect.ValueOf(in)
	if item.Kind() == reflect.Map {
		for _, k := range item.MapKeys() {
			v := item.MapIndex(k)
			kStr, _ := InterfaceToStr(k.Interface())
			R_data[kStr] = v.Interface()
		}
	} else {
		return nil, errors.New("Interface To MapInterface is wrong!")
	}
	return R_data, nil
} //end func InterfaceToMapInterface

//interface 类型转string
func InterfaceToStr(v interface{}) (string, bool) {
	var R_str string
	switch v.(type) {
	case string:
		R_str = v.(string)
	case int:
		R_str = strconv.Itoa(v.(int))
	case int64:
		R_str = strconv.FormatInt(v.(int64), 10)
	case int32:
		R_str = strconv.FormatInt(int64(v.(int32)), 10)
	case int16:
		R_str = strconv.FormatInt(int64(v.(int16)), 10)
	case int8:
		R_str = strconv.FormatInt(int64(v.(int8)), 10)
	case float64:
		t := int(v.(float64))
		R_str = strconv.Itoa(t)
	case bool:
		R_str = strconv.FormatBool(v.(bool))
	default:
		return "", false
	}
	return R_str, true
} //end func InterfaceToStr
