package wx

import (
	"encoding/xml"
	"epaygo/helper"
	"epaygo/helper/cryptoHelper"
	"epaygo/wx/wxConst"
	"sort"

	"strings"

	"encoding/json"

	"errors"
)

type WxPayData struct {
	xmlMap helper.XmlMap
}

func NewWxPayData() *WxPayData {
	wxPayData := WxPayData{}
	wxPayData.xmlMap = make(helper.XmlMap)
	return &wxPayData
}

func (w *WxPayData) SetValue(key string, value string) {
	w.xmlMap[key] = value
}

func (w *WxPayData) GetValue(key string) string {
	return w.xmlMap[key]
}
func (w *WxPayData) RemoveKey(key string) {
	delete(w.xmlMap, key)
}

func (w *WxPayData) IsSet(key string) bool {
	_, ok := w.xmlMap[key]
	return !ok
}

func (w *WxPayData) ToXml() string {
	if len(w.xmlMap) == 0 {
		return ""
	}
	if x, err := xml.MarshalIndent(w.xmlMap, "", " "); err == nil {
		return string(x)
	}
	return ""
}

func (w *WxPayData) ToUrl() string {
	return w.sortedUrlWX(&(w.xmlMap))
}
func (w *WxPayData) ToJson() string {

	v, _ := json.Marshal(w.xmlMap)

	return string(v)
}

func (w *WxPayData) MakeSign(key string) (sign string) {
	str := w.ToUrl()
	str += "&key=" + key
	sign, _ = cryptoHelper.GetMD5Hash(str)
	return
}

func (w *WxPayData) CheckSign(key string) (isSign bool, err error) {
	if !w.IsSet(wxConst.Sign) {
		return false, errors.New("10001")
	} else if len(strings.TrimSpace(w.GetValue(wxConst.Sign))) == 0 {
		return false, errors.New("10002")
	} else if w.GetValue(wxConst.Sign) == w.MakeSign(key) {
		return true, nil
	}
	return false, errors.New("10002")

}

func (w *WxPayData) FromXml(xml string, key string) (xmlMap helper.XmlMap, errKey error) {

	if xmlMap, err := helper.XmlToWxMap(xml); err == nil {
		w.xmlMap = xmlMap
	} else {
		return nil, err
	}
	return w.xmlMap, nil

}

func (w *WxPayData) sortedUrlWX(m *helper.XmlMap) string {
	if len(*m) == 0 {
		return ""
	}
	sk := w.sortedKeysWX(m)
	var sortedData string
	for _, k := range sk {
		sortedData += "&" + k + "=" + (*m)[k]
	}
	return sortedData[1:]
}
func (w *WxPayData) sortedKeysWX(m *helper.XmlMap) []string {
	sortedKeys := make([]string, len(*m))
	i := 0
	for key := range *m {
		sortedKeys[i] = key
		i++
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}
