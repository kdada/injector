package injector

import (
	"fmt"
	"reflect"
)

//注册kind 为 Func的创建器
var funcRegErr = registerProviderCreator(reflect.Func, newfuncInjectorProvider)

type funcParamInjectInfo struct {
	param reflect.Type //参数类型
	name  string       //参数类型名称
	index uint8        //参数序号
}

// 方法注入器提供器
type funcInjectorProvider struct {
	funcInterface interface{}
	funcType      reflect.Type           //函数类型
	funcValue     reflect.Value          //函数反射值
	injectInfo    []*funcParamInjectInfo //函数参数注入信息
}

// CreateInjector创建注入器
func (this *funcInjectorProvider) CreateInjector() (Injector, error) {
	var injector = &funcInjector{injectInfo: this.injectInfo}
	injector.value = this.funcInterface
	injector.mapInfo = make(map[string]ValueGenerator)
	//设置上级值获取器
	injector.SetParentValueGetter(GlobalMapper)
	return injector, nil
}

// newfuncInjectorProvider 创建函数注入器提供器
func newfuncInjectorProvider(value interface{}) (InjectorProvider, error) {
	var provider = &funcInjectorProvider{}
	provider.funcInterface = value
	provider.funcValue = reflect.ValueOf(value)
	provider.funcType = reflect.TypeOf(value)
	var length = provider.funcType.NumIn()
	provider.injectInfo = make([]*funcParamInjectInfo, 0, length)
	for i := 0; i < length; i++ {
		var paramType = provider.funcType.In(i)
		var info = &funcParamInjectInfo{
			param: paramType,
			name:  paramType.String(),
			index: uint8(i),
		}
		provider.injectInfo = append(provider.injectInfo, info)
	}
	return provider, nil
}

// 方法注入器
type funcInjector struct {
	commonInjector
	injectInfo []*funcParamInjectInfo //函数参数注入信息
}

// Inject 启动注入过程并返回注入结果
func (this *funcInjector) Inject() (interface{}, error) {
	var value = reflect.ValueOf(this.value)
	var valueType = reflect.TypeOf(this.value)
	var paramValues = make([]reflect.Value, valueType.NumIn())
	for i := 0; i < len(this.injectInfo); i++ {
		var info = this.injectInfo[i]
		var fieldValue, err = this.ValueByTypeName(info.name)
		if err == nil {
			paramValues[info.index] = reflect.ValueOf(fieldValue)
		} else {
			fmt.Println(err)
		}
	}
	var result = value.Call(paramValues)
	var typeResult = make([]interface{}, len(result))
	for j := 0; j < len(typeResult); j++ {
		typeResult[j] = result[j].Interface()
	}
	return typeResult, nil
}
