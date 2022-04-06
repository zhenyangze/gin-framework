// Package helpers provides ...
package helpers

import gojsonq "github.com/thedevsaddam/gojsonq/v2"

func JsonInit() *gojsonq.JSONQ {
	return gojsonq.New()
}

func JsonFromStr(jsonStr string) *gojsonq.JSONQ {
	return JsonInit().FromString(jsonStr)
}

func JsonFromFile(fileName string) *gojsonq.JSONQ {
	return JsonInit().File(fileName)
}

func JsonGetFromStr(jsonStr string) (*gojsonq.Result, error) {
	return JsonInit().FromString(jsonStr).GetR()
}

func JsonGetFromFile(fileName string) (*gojsonq.Result, error) {
	return JsonInit().File(fileName).GetR()
}

func JsonFindFromStr(jsonStr, key string) (*gojsonq.Result, error) {
	ret, err := JsonInit().FromString(jsonStr).FindR(key)
	return ret, err
}

func JsonFindFromFile(fileName string, key string) (*gojsonq.Result, error) {
	return JsonInit().File(fileName).FindR(key)
}
