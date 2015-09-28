package injector

// 方法注入器提供器
type funcInjectorProvider struct {
}

// CreateInjector创建注入器
func (this *funcInjectorProvider) CreateInjector() (Injector, error) {
	return nil, nil
}

// 方法注入器
type funcInjector struct {
	commonInjector
}

// Inject 启动注入过程并返回注入结果
func (this *funcInjector) Inject() (interface{}, error) {
	return nil, InjectorErrorInvalidInjector.Error()
}
