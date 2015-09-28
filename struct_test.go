package injector

import (
	"fmt"
	"testing"
)

type SomeService interface {
	Output()
}

type OutputService struct {
}

func (this *OutputService) Output() {
	fmt.Println("测试测试1")
}

type SpString string

type SomeStruct struct {
	Value   int         `inject`
	Service SomeService `inject`
	Command SpString    `inject`
	Value2  string      `inject`
}

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
		fmt.Println("1-", err)
		return
	}
	var injector, _ = provider.CreateInjector()
	var value, _ = injector.Inject()
	var v, ok = value.(*SomeStruct)
	if ok {
		v.Service.Output()
		fmt.Println(v)
	} else {
		fmt.Println("注入失败")
	}

}
