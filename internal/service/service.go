package service

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/flosch/pongo2/v6"
)

//go:embed service.impl
var serviceImpl string

func Execute(v interface{}) (to string, err error) {
	return execute(serviceImpl, v)
}

func execute(from string, v interface{}) (to string, err error) {
	tpl, err := pongo2.FromString(from)
	if err != nil {
		return
	}

	var ctx pongo2.Context

	// 如果v是map[string]interface{}，直接使用
	if ctx, ok := v.(map[string]interface{}); ok {
		to, err = tpl.Execute(ctx)
		return
	}

	// 如果v是JSON字符串，解析JSON
	if jsonStr, ok := v.(string); ok {
		err = json.Unmarshal([]byte(jsonStr), &ctx)
		if err != nil {
			return
		}
		to, err = tpl.Execute(ctx)
		return
	}

	// 如果v是struct，转换为map
	vValue := reflect.ValueOf(v)
	if vValue.Kind() == reflect.Struct {
		ctx = make(pongo2.Context)
		vType := vValue.Type()
		for i := 0; i < vValue.NumField(); i++ {
			ctx[vType.Field(i).Name] = vValue.Field(i).Interface()
		}
		to, err = tpl.Execute(ctx)
		return
	}

	err = fmt.Errorf("unsupported type for template execution")
	return
}
