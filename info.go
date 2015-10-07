package injector

import (
	"errors"
	"fmt"
)

// 注入错误信息
type InjectorError string

// 错误码
const (
	InjectorErrorInvalidInjector InjectorError = "E10010:InjectorErrorInvalidInjector,无效的注入器"

	InjectorErrorNoCreator         InjectorError = "E30000:InjectorErrorNoCreator,不存在指定Kind(%s)的创建器"
	InjectorErrorNoType            InjectorError = "E30010:InjectorErrorNoType,找不到指定类型(%s)的值"
	InjectorErrorInvalidStructType InjectorError = "E30020:InjectorErrorInvalidStructType,类型(%s)不能转换为结构体类型"
)

// Format 格式化错误信息并生成新的错误信息
func (this InjectorError) Format(data ...interface{}) InjectorError {
	return InjectorError(fmt.Sprintf(string(this), data...))
}

// Error 生成error类型
func (this InjectorError) Error() error {
	return errors.New(string(this))
}
