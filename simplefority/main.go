package main

import "fmt"

type Fruit interface {
	Show()
}

// 工厂类
type AbsuibteFotity interface {
	CreateFotity() Fruit
}
type Apple struct{}

func (r *Apple) Show() {
	fmt.Println("我是生产出来的苹果")
}

type Orange struct{}

func (r *Orange) Show() {
	fmt.Println("我是生产出来的橙子")
}

type Banna struct{}

func (r *Banna) Show() {
	fmt.Println("我是生产出来的香蕉")
}

// 苹果工厂
type AppleFority struct{}

func (r *AppleFority) CreateFotity() Fruit {
	return &Apple{}
}

// 橙子工厂
type OrangeFority struct{}

func (r *OrangeFority) CreateFotity() Fruit {
	return &Orange{}
}

// 香蕉工厂
type BannaFority struct{}

func (r *Banna) CreateFotity() Fruit {
	return &Banna{}
}

func main() {
	//	生产苹果
	apple := &AppleFority{}
	app := apple.CreateFotity()
	app.Show()

	orange := OrangeFority{}
	oranges := orange.CreateFotity()
	oranges.Show()
}
