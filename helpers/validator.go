package helpers

import (
	"reflect"
	"strconv"
)

// 判断变量是否为空
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

func IsEmpty(params interface{}) bool {
	var (
		flag          bool = true
		default_value reflect.Value
	)

	r := reflect.ValueOf(params)

	//获取对应类型默认值
	default_value = reflect.Zero(r.Type())
	//由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = false
	}
	return flag
}

// 等价于PHP函数gettype()
func Gettype(variable interface{}) string {
	return reflect.TypeOf(variable).Kind().String()
}

// 等价于PHP函数is_array()
func IsArray(variable interface{}) bool {
	var b bool
	if Gettype(variable) == "array" {
		b = true
	}
	return b
}

// 等价于PHP函数is_bool
func IsBool(variable interface{}) bool {
	var b bool
	if Gettype(variable) == "bool" {
		b = true
	}
	return b
}

// 等价于PHP函数is_double
func IsDouble(variable interface{}) bool {
	var b bool
	variableType := Gettype(variable)
	if variableType == "float64" || variableType == "float32" {
		b = true
	}
	return b
}

// is_double的别名
func IsFloat(variable interface{}) bool {
	return IsDouble(variable)
}

// 等价于PHP函数is_int
func IsInt(variable interface{}) bool {
	var b bool
	variableType := Gettype(variable)
	if variableType == "int" || variableType == "uint" || variableType == "uint64" ||
		variableType == "int64" || variableType == "uint32" || variableType == "int32" ||
		variableType == "int8" || variableType == "uint8" {
		b = true
	}
	return b
}

// is_integer是is_int的别名
func IsInteger(variable interface{}) bool {
	return IsInt(variable)
}

// is_long是is_int的别名
func IsLong(variable interface{}) bool {
	return IsInt(variable)
}

// 等价于PHP函数is_string
func IsString(variable interface{}) bool {
	var b bool
	if Gettype(variable) == "string" {
		b = true
	}
	return b
}

// 等价于PHP函数is_numeric
func IsNumeric(variable interface{}) bool {

	var b bool
	if IsInt(variable) || IsFloat(variable) {
		b = true
	}

	if IsString(variable) {
		// 接口转string
		tmpStr, err := variable.(string)
		if err {
			_, err := strconv.Atoi(tmpStr)
			if err == nil {
				b = true
			}
		}
	}
	return b
}
