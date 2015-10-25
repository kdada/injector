package injector

import "reflect"

func init() {
	//注册kind 为 Func的创建器
	RegisterProviderCreator(reflect.Func, newfuncInjectorProvider)
	//注册kind 为 Struct的创建器
	RegisterProviderCreator(reflect.Struct, newstructInjectorProvider)
}
