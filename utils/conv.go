package utils

import (
	"reflect"
	"strconv"
)

//interface TO MapInterface
func Inter2Map(inParam interface{}) (map[string]interface{}, bool) {
	var reData = make(map[string]interface{})
	outItem := reflect.ValueOf(inParam)
	if outItem.Kind() == reflect.Map {
		for _, k := range outItem.MapKeys() {
			v := outItem.MapIndex(k)
			kStr, _ := Inter2Str(k.Interface())
			reData[kStr] = v.Interface()
		}
	} else {
		return nil, false
	}
	return reData, true
} //end func InterfaceToMapInterface

//interface 类型转string
func Inter2Str(v interface{}) (string, bool) {
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
