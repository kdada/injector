package injector

import (
	"reflect"
	"sync"
)

//// 公共注入器提供器
//type commonInjectorProvider struct {
//}

//// CreateInjector创建注入器
//func (this *commonInjectorProvider) CreateInjector() (Injector, error) {
//	return nil, nil
//}

// 公共值生成器
type commonValueGenerator struct {
	value interface{}
}

// Type 获取值的真实类型
func (this *commonValueGenerator) Type() reflect.Type {
	return reflect.TypeOf(this.value)
}

// Value 获取值
func (this *commonValueGenerator) Value() interface{} {
	return this.value
}

// 公共注入器
type commonInjector struct {
	value   interface{}               //被注入的值
	parent  InjectorValueGetter       //值获取器,当当前mapInfo不存在指定类型值时继续查找parent的值
	infoMu  sync.Mutex                //注入信息互斥锁
	mapInfo map[string]ValueGenerator //注入映射信息
}

// Map 映射值到值类型
func (this *commonInjector) Map(value interface{}) {
	this.MapTo(value, value)
}

// MapTo 映射值到指定的类型
// 接口指针类型的值将被转换为接口类型的值
func (this *commonInjector) MapTo(value interface{}, typeValue interface{}) {
	var mapValue = reflect.ValueOf(value)
	if mapValue.Kind() == reflect.Ptr && mapValue.Elem().Kind() == reflect.Interface {
		//如果类型是接口指针,那么转换为接口类型
		//接口指针仅仅能用于获取接口的Type信息
		mapValue = mapValue.Elem()
	}
	var vg = &commonValueGenerator{mapValue.Interface()}
	this.MapGeneratorTo(vg, typeValue)
}

// MapGeneratorTo 映射ValueGenerator的Value为指定的类型
// 接口指针类型的值将被转换为接口类型的值
func (this *commonInjector) MapGeneratorTo(value ValueGenerator, typeValue interface{}) {
	var mapType = reflect.TypeOf(typeValue)
	if mapType.Kind() == reflect.Ptr && mapType.Elem().Kind() == reflect.Interface {
		//如果类型是接口指针,那么转换为接口类型
		//接口指针仅仅能用于获取接口的Type信息
		mapType = mapType.Elem()
	}
	this.MapGeneratorToTypeName(value, mapType.String())
}

// MapToTypeName 映射值到指定的包名.类型名
// 接口指针类型的值将被转换为接口类型的值
func (this *commonInjector) MapToTypeName(value interface{}, typeName string) {
	var mapValue = reflect.ValueOf(value)
	if mapValue.Kind() == reflect.Ptr && mapValue.Elem().Kind() == reflect.Interface {
		//如果类型是接口指针,那么转换为接口类型
		//接口指针仅仅能用于获取接口的Type信息
		mapValue = mapValue.Elem()
	}
	var vg = &commonValueGenerator{mapValue.Interface()}
	this.MapGeneratorToTypeName(vg, typeName)
}

// MapGeneratorToTypeName 映射ValueGenerator的Value到指定的包名.类型名
func (this *commonInjector) MapGeneratorToTypeName(value ValueGenerator, typeName string) {
	this.infoMu.Lock()
	defer this.infoMu.Unlock()
	this.mapInfo[typeName] = value
}

// Value 根据类型获取指定的值
func (this *commonInjector) Value(typeValue interface{}) (interface{}, error) {
	var vt = reflect.TypeOf(typeValue)
	return this.ValueByTypeName(vt.String())
}

// ValueByTypeName 根据类型完整名称获取指定的值
func (this *commonInjector) ValueByTypeName(typeName string) (interface{}, error) {
	var generator, ok = this.valueGenerator(typeName)
	if ok {
		return generator.Value(), nil
	}
	return nil, InjectorErrorNoType.Format(typeName).Error()
}

// valueGenerator 根据类型完整名称获取指定的值生成器
func (this *commonInjector) valueGenerator(typeName string) (ValueGenerator, bool) {
	var generator, ok = this.mapInfo[typeName]
	if !ok && this.parent != nil {
		generator, ok = this.parent.valueGenerator(typeName)
	}
	return generator, ok
}

// SetInjectedValue 设置被注入的对象,必须是指针
// value:被注入的对象,可以是结构体或方法
func (this *commonInjector) SetInjectedValue(value interface{}) {
	this.value = value
}

// SetParentValueGetter 设置上级值获取器
func (this *commonInjector) SetParentValueGetter(getter InjectorValueGetter) {
	this.parent = getter
}

// Inject 启动注入过程并返回注入结果
func (this *commonInjector) Inject() (interface{}, error) {
	return nil, InjectorErrorInvalidInjector.Error()
}
