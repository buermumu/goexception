// 异常处理包
package exception

import (
	_ "fmt"
)

// 自定义异常处理方法,
type Type_Catch_Handler func(message interface{}, args ...interface{})

// 异常类
type Exception struct {
	message interface{}
}

// get Instance
func New() *Exception {
	return new(Exception)
}

// Try run user process
func (this *Exception) Try(process func()) {
	defer func() {
		if err := recover(); err != nil {
			this.message = err
		}
	}()
	process()
}

// Throw process exception message
func (this *Exception) Throw(message string) {
	panic(message)
}

// Catch
func (this *Exception) Catch(handler Type_Catch_Handler, args ...interface{}) {
	handler(this.message, args...)
}
