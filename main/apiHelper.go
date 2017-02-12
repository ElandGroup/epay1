package main

import (
	"encoding/json"
	"reflect"

	"errors"
)

func BadRequestMessage(reqParam interface{}) string {
	return JoinMessage("A required parameter is missing or doesn't have the right format", reqParam)
}

func JoinMessage(preMessage string, reqParam interface{}) string {
	result, _ := JoinMessageWithCheck(preMessage, reqParam)
	return result
}

func JoinMessageWithCheck(preMessage string, reqParam interface{}) (string, error) {
	result := preMessage
	var err error
	if value, ok := reqParam.(string); ok {
		result += value
		return result, nil
	}
	var info string
	if info, err = ConvJson(reqParam); err != nil {
		return result, err
	} else {
		return result + "," + info, err
	}

}

func ConvJson(anyObject interface{}) (string, error) {
	var result string
	var err error
	if isStruct(anyObject) || isPointer(anyObject) {
		var bytevv []byte
		bytevv, err = json.Marshal(anyObject)
		result = string(bytevv)
		return result, err
	}
	switch vv := anyObject.(type) {
	case []interface{}:
		var bytevv []byte
		bytevv, err = json.Marshal(vv)
		result = string(bytevv)
	default:
		result = ""
		err = errors.New("string,map,struct only is supported")
	}
	return result, err
}

func isStruct(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Struct
}

func isPointer(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Ptr
}
