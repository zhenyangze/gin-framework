package helpers

import "reflect"

func InvokeMethod(object interface{}, methodName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
}

func InvokeAttr(object interface{}, attrName string) interface{} {
	return reflect.ValueOf(object).Elem().FieldByName(attrName)
}
