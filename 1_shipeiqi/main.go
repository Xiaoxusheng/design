package main

import "fmt"

type kill interface {
	fight()
}

// 无关类
type a struct{}

func (a *a) flay() {
	fmt.Println("降龙十八掌")
}

// 适配器
type att struct {
	a *a
}

func (ac *att) fight() {
	ac.a.flay()
}

func Newatt(a *a) *att {
	return &att{a}
}

type b struct {
	k kill
}

func (b *b) chag() {
	b.k.fight()
}

func Newb(k kill) *b {
	return &b{k}
}

func main() {
	ac := Newb(Newatt(&a{}))
	ac.chag()

}
