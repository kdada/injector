package injector

import (
	"fmt"
	"reflect"
)

type structFieldInjectInfo struct {
	field reflect.StructField //字段类型
	name  string              //字段类型名称
}

// 结构体注入器提供器
type structInjectorProvider struct {
	valueType  reflect.Type             //被注入的结构体的类型
	injectInfo []*structFieldInjectInfo //注入信息
}

// newstructInjectorProvider 创建结构体注入器提供器
func newstructInjectorProvider(value interface{}) (InjectorProvider, error) {
	var provider = &structInjectorProvider{}
	provider.valueType = reflect.TypeOf(value)
	provider.injectInfo = getAllFieldInfo(provider.valueType)
	return provider, nil
}

// getAllFieldInfo 获取字段信息
func getAllFieldInfo(structType reflect.Type) []*structFieldInjectInfo {
	if structType.Kind() != reflect.Struct {
		return make([]*structFieldInjectInfo, 0)
	}
	var fields = make([]*structFieldInjectInfo, 0, structType.NumField())
	for i := 0; i < structType.NumField(); i++ {
		var fieldInfo = structType.Field(i)
		if fieldInfo.Anonymous {
			fields = append(fields, getAllFieldInfo(fieldInfo.Type)...)
		} else {
			var injectTag = fieldInfo.Tag.Get("inject")
			if injectTag == "" {
				injectTag = string(fieldInfo.Tag)
			}
			if injectTag == "inject" {
				var info = &structFieldInjectInfo{
					field: fieldInfo,
					name:  fieldInfo.Type.String(),
				}
				fields = append(fields, info)
			}
		}
	}
	return fields
}

// CreateInjector 创建注入器
func (this *structInjectorProvider) CreateInjector() (Injector, error) {
	var injector = &structInjector{injectInfo: this.injectInfo}
	injector.value = reflect.New(this.valueType).Interface()
	injector.mapInfo = make(map[string]ValueGenerator)
	//设置上级值获取器
	injector.SetParentValueGetter(GlobalMapper)
	return injector, nil
}

// 结构体注入器
type structInjector struct {
	commonInjector
	injectInfo []*structFieldInjectInfo //注入信息
}

// Inject 启动注入过程并返回注入结果
func (this *structInjector) Inject() (interface{}, error) {
	var value = reflect.ValueOf(this.value)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	for i := 0; i < len(this.injectInfo); i++ {
		var info = this.injectInfo[i]
		var fieldValue, err = this.ValueByTypeName(info.name)
		if err == nil {
			value.FieldByIndex(info.field.Index).Set(reflect.ValueOf(fieldValue))
		} else {
			fmt.Println(err)
		}
	}
	return this.value, nil
}
