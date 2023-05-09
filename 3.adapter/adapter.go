package main

import "fmt"

/*
适配器模式
*/

// 抽象层
type V5 interface {
	UseV5()
}

type Iphone struct {
	v V5
}

func NewIphone(v V5) *Iphone {
	return &Iphone{v: v}
}

func (v Iphone) charge() {
	fmt.Println("phone开始充电")
	v.v.UseV5()
}

type Use220v struct {
}

func (use Use220v) Use220v() {
	fmt.Println("使用220v电压")
}

// 适配器
type adapter struct {
	v *Use220v
}

func (r adapter) UseV5() {
	fmt.Println("使用适配器充电")
	r.v.Use220v()
}

func NewAdapter(v *Use220v) *adapter {
	return &adapter{v}
}

func main() {
	phone := NewIphone(NewAdapter(&Use220v{}))
	phone.charge()
}
