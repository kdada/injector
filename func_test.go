package injector

import (
	"fmt"
	"strconv"
	"testing"
)

func FuncTest(a string) string {
	return a + "-PASS"
}

type FuncStruct struct {
}

func (this *FuncStruct) Test(a string, b int) string {
	return a + strconv.Itoa(b)
}

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
