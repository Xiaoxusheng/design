package main

import "fmt"

/*
装饰器模式
*/
type Phone interface {
	show()
}

// 装饰
type Decorator struct {
	Phone Phone
}

type HuaWei struct {
}

func (h *HuaWei) show() {
	fmt.Println("huawei手机")
}

type Xiaomi struct {
}

func (x *Xiaomi) show() {
	fmt.Println("Xiaomi手机")
}

type MoDecorator struct {
	Decorator
}

func (m *MoDecorator) show() {
	m.Phone.show()
	fmt.Println("手机贴膜了")
}

func NewMoDecorator(h Phone) Phone {
	return &MoDecorator{Decorator{h}}
}

type keDecorator struct {
	Decorator
}

func (h *keDecorator) show() {
	h.Phone.show()
	fmt.Println("手机套壳了")
}

func NewKeDecorator(h Phone) Phone {
	return &keDecorator{Decorator{h}}
}

func main() {
	h := &HuaWei{}
	h.show()

	h1 := NewKeDecorator(h)
	h1.show()

	h2 := NewMoDecorator(h1)
	h2.show()
}
