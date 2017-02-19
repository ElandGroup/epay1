package helper

import (
	"encoding/json"
	. "epaygo/core/common"
	"fmt"
	"reflect"
	"strconv"

	"errors"
)

//private,don't change public
func MessageString(resourceKey int, params ...interface{}) string {
	return fmt.Sprintf(MessageMap[resourceKey], params...)
}
func NewApiError(resourceKey int, params ...interface{}) (apiError *APIError) {
	return NewApiErrorWithDetails(resourceKey, "", params...)
}

func NewApiMessage(resourceKey int, params ...interface{}) *APIResult {
	return NewApiMessageWithDetails(resourceKey, "", params...)
}

//success:false,details:null,message:10012 +params
func CheckRequestFormat(params string) *APIResult {
	return NewApiMessage(10012, params)
}

//success:false,details:detail,message:10001
func SystemErrorMessage(detail string) *APIResult {
	return NewApiMessageWithDetails(10001, detail)
}

func NewApiErrorWithDetails(resourceKey int, details string, params ...interface{}) (apiError *APIError) {
	return &APIError{Code: resourceKey, Message: MessageString(resourceKey, params...), Details: details}
}

func NewApiMessageWithDetails(resourceKey int, details string, params ...interface{}) *APIResult {
	return &APIResult{Success: false, Error: *NewApiErrorWithDetails(resourceKey, details, params...)}
}

func ConvJson(anyObject interface{}) (result string, err error) {

	val := reflect.ValueOf(anyObject).Elem()
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		result = val.String()
	case reflect.Slice, reflect.Map, reflect.Struct, reflect.Array:
		var bytevv []byte
		bytevv, err = json.Marshal(val)
		result = string(bytevv)
	default:
		result = ""
		err = errors.New("Type is not recognized")
	}
	return
}
