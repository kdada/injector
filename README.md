# injector
依赖注入

目标:  
1.结构体注入[完成]  
2.方法参数注入[完成]  


```go
//结构体注入
func TestStructInjector(t *testing.T) {
	//注册要注入的值
	GlobalMapper.Map(12544)
	GlobalMapper.MapTo("xxxx", "")
	GlobalMapper.MapTo(SpString("xxxxsp"), SpString(""))
	//	GlobalMapper.MapToTypeName(&OutputService{}, "injector.SomeService")
	GlobalMapper.MapTo(&OutputService{}, (*SomeService)(nil))
	fmt.Println(GlobalMapper)

	var provider, err = NewInjectorProvider(SomeStruct{})
	if err != nil {
		t.Error(err)
		return
	}
	var injector, _ = provider.CreateInjector()
	var value, _ = injector.Inject()
	var v, ok = value.(*SomeStruct)
	if ok {
		v.Service.Output()
		fmt.Println(v)
	} else {
		t.Error("结构体注入失败")
	}

}


//方法注入
func TestFuncInject(t *testing.T) {
	var provider, _ = NewInjectorProvider(FuncTest)
	var injector, _ = provider.CreateInjector()
	injector.Map("some")
	var resultArr, err = injector.Inject()
	var result = resultArr.([]interface{})
	fmt.Println(result, err)

	if result[0] != "some-PASS" {
		t.Error("FuncTest注入错误")
		return
	}
	var provider2, _ = NewInjectorProvider((new(FuncStruct)).Test)
	var injector2, _ = provider2.CreateInjector()
	injector2.Map("kkkk")
	injector2.Map(2333)
	var result2Arr, err2 = injector2.Inject()
	var result2 = result2Arr.([]interface{})
	fmt.Println(result2, err2)

	if result2[0] != "kkkk2333" {
		t.Error("FuncStruct.Test注入错误")
		return
	}

}
```