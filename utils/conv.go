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

//interface type to string type
func Inter2Str(param interface{}) (string, bool) {
	//def return data
	ret := ""
	//param type
	switch param.(type) {
	case string:
		ret = param.(string)
	case int:
		ret = strconv.Itoa(param.(int))
	case int64:
		ret = strconv.FormatInt(param.(int64), 10)
	case int32:
		ret = strconv.FormatInt(int64(param.(int32)), 10)
	case int16:
		ret = strconv.FormatInt(int64(param.(int16)), 10)
	case int8:
		ret = strconv.FormatInt(int64(param.(int8)), 10)
	case float64:
		t := int(param.(float64))
		ret = strconv.Itoa(t)
	case bool:
		ret = strconv.FormatBool(param.(bool))
	default:
		return "", false
	}

	//return value
	return ret, true
} //end func InterfaceToStr
