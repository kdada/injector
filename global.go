package injector

import (
	"reflect"
	"sync"
)

// 存储全局注入信息
var GlobalMapper InjectorValueMapper = &commonInjector{mapInfo: make(map[string]ValueGenerator)}

type ProviderCreator func(interface{}) (InjectorProvider, error)

var (
	creatorsMu sync.Mutex                               //生成器互斥锁
	creators   = make(map[reflect.Kind]ProviderCreator) //注入器生成器
)

// NewInjectorProvider 创建一个InjectorProvider
//  valueType:被注入的类型
func NewInjectorProvider(valueType interface{}) (InjectorProvider, error) {
	var vt = reflect.TypeOf(valueType)
	var vtKind = vt.Kind()
	var creator, ok = creators[vtKind]
	if !ok {
		return nil, InjectorErrorNoCreator.Format(vtKind.String()).Error()
	}
	return creator(valueType)
}

// RegisterProviderCreator 注册InjectorProvider创建器
func RegisterProviderCreator(kind reflect.Kind, creator ProviderCreator) {
	creatorsMu.Lock()
	defer creatorsMu.Unlock()
	creators[kind] = creator
}
