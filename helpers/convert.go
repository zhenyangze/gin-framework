package helpers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// ValueInterfaceToString interface转string，非map[string]interface{}
func ValueInterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// ValueInterfaceToInt interface转int，map[string]interface{}
func ValueInterfaceToInt(_value interface{}) int64 {
	return StringToInt(ValueInterfaceToString(_value))
}

// MapInterfaceToString interface转string，针对map[string]interface{}的某个键
func MapInterfaceToString(_map map[string]interface{}, _key string) string {
	value := _map[_key].(string)
	return value
}

// ArrayInterfaceToString interface转string，准对一维数组[]string{}或[]int{}
func ArrayInterfaceToString(_array interface{}) string {
	value := fmt.Sprintf("%v", _array)
	return value
}

// StringToInt string转int
func StringToInt(_str string) int64 {
	_int, err := strconv.ParseInt(_str, 10, 64) // string转int
	if err != nil {                             // 报错则默认返回0
		_int = 0
		//fmt.Println("格式转换错误，默认为0。")
		//fmt.Println(err)
	}
	return _int
}

// IntToString int转string
func IntToString(_int int64) string {
	_str := strconv.FormatInt(_int, 10)
	return _str
}

func ByteToString(_byte []byte) string {
	return string(_byte)
}

func StringToByte(_str string) []byte {
	return []byte(_str)
}

func ByteToStringFast(_byte []byte) string {
	return *(*string)(unsafe.Pointer(&_byte))
}

func StringToByteFast(_str string) (_byte []byte) {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&_str))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&_byte))
	bs.Data = ss.Data
	bs.Len = ss.Len
	bs.Cap = ss.Len
	return _byte
}

func Int32ToInt64(_int int32) int64 {
	return int64(_int)
}

func Int64ToInt32(_int int64) int32 {
	return int32(_int)
}

// StringToFloat string转float
func StringToFloat(_str string) float64 {
	_float, err := strconv.ParseFloat(_str, 64) // string转int
	if err != nil {                             // 报错则默认返回0
		_float = 0.0
		//fmt.Println("格式转换错误，默认为0。")
		//fmt.Println(err)
	}
	return _float
}

// FloatToString float转string
func FloatToString(_float float64) string {
	_str := strconv.FormatFloat(_float, 'e', 10, 64)
	return _str
}

func MapInterfaceToJson(_map map[string]interface{}) []byte {
	_json, _ := json.Marshal(_map)
	return _json
}

func InterfaceToString(_array interface{}) string {
	key := ArrayInterfaceToString(_array)
	key = EncodeURL(key)
	return key
}

func JsonEncode(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// 需要判断返回值，否则会出问题
func JsonDecode(_string string, _type interface{}) error {
	return json.Unmarshal([]byte(_string), &_type)
}
