package helper

import "github.com/axgle/mahonia"

func EncodingAlUrl(response string) string {
	c := mahonia.NewDecoder("GBK").ConvertString(response)
	return mahonia.NewEncoder("UTF-8").ConvertString(c)
}

func EncodingUrl(response string) string {
	c := mahonia.NewDecoder("UTF-8").ConvertString(response)
	return mahonia.NewEncoder("GBK").ConvertString(c)
}

func EncodingUTF8(rawString string) string {
	return mahonia.NewEncoder("UTF-8").ConvertString(rawString)
}

func DecodingUTF8(rawString string) string {
	return mahonia.NewDecoder("UTF-8").ConvertString(rawString)
}

func EncodingGBK(rawString string) string {
	return mahonia.NewEncoder("GBK").ConvertString(rawString)
}

func DecodingGBK(rawString string) string {
	return mahonia.NewDecoder("GBK").ConvertString(rawString)
}
