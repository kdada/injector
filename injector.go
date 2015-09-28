// 默认情况下,InjectorProvider能够解析结构体和方法的结构,并将结构信息存储在内部,
// 在需要注入的时候,通过CreateInjector生成一个特定的注入器,
// 特定的注入器包含一个指向指定注入结构(结构体或方法参数)的新实例
// 特定的注入器可以额外设置临时使用的注入值
// 特定的注入器通过调用Inject方法生成注入后产生的结果
// 结果可能是结构体指针或函数返回结果数组
package injector

import "reflect"

// InjectorProvider 用于解析并存储被注入的结构体或方法
// 并生成相应的Injector
type InjectorProvider interface {
	// CreateInjector 创建一个新的Injector
	CreateInjector() (Injector, error)
}

// ValueGenerator 注入器映射方法
// 使用该方法返回指定类型的值
type ValueGenerator interface {
	// Type 值的实际类型
	Type() reflect.Type
	// Value 获取值
	Value() interface{}
}

// 值设置器
type InjectorValueSetter interface {
	// Map 映射值到值类型
	Map(value interface{})
	// MapTo 映射值到指定的类型
	MapTo(value interface{}, typeValue interface{})
	// MapGeneratorTo 映射ValueGenerator的Value为指定的类型
	MapGeneratorTo(value ValueGenerator, typeValue interface{})
	// MapToTypeName 映射值到指定的类型
	// typeName:有包名的则格式为 包名.类型名 ,没有包名的格式为 类型名
	MapToTypeName(value interface{}, typeName string)
	// MapGeneratorToTypeName 映射ValueGenerator的Value到指定的类型
	// typeName:有包名的则格式为 包名.类型名 ,没有包名的格式为 类型名
	MapGeneratorToTypeName(value ValueGenerator, typeName string)
}

// 值获取器
type InjectorValueGetter interface {
	// Value 根据类型获取指定的值
	Value(typeValue interface{}) (interface{}, error)
	// ValueByTypeName 根据类型完整名称获取指定的值
	ValueByTypeName(typeName string) (interface{}, error)
	// valueGenerator 根据类型完整名称获取指定的值生成器
	valueGenerator(typeName string) (ValueGenerator, bool)
}

type InjectorValueMapper interface {
	InjectorValueSetter
	InjectorValueGetter
}

// Injector 注入器
type Injector interface {
	InjectorValueMapper
	// SetInjectedValue 设置被注入的对象,必须是指针
	// value:被注入的对象,可以是结构体或方法
	SetInjectedValue(value interface{})
	// SetParentValueGetter 设置上级值获取器
	SetParentValueGetter(getter InjectorValueGetter)
	// Inject 启动注入过程并返回注入结果
	Inject() (interface{}, error)
}
