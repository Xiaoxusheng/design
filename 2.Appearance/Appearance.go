package main

import "fmt"

/*
外观模式
*/

type Tv struct {
}

func (t *Tv) open() {
	fmt.Println("打开电视")
}

func (t *Tv) watch() {
	fmt.Println("我在看电视")
}

func (t *Tv) offTv() {
	fmt.Println("关掉电视")
}

type sound struct {
}

func (s *sound) open() {
	fmt.Println("打开音响")
}

func (s *sound) off() {
	fmt.Println("关掉音响")
}

type Projector struct {
}

func (p *Projector) open() {
	fmt.Println("打开投影仪")
}
func (p *Projector) off() {
	fmt.Println("关闭投影仪")
}

type FamilyCinema struct {
	t Tv
	s sound
	p Projector
}

func (f *FamilyCinema) look() {
	f.t.open()
	f.s.open()
	f.p.open()
}

func (f *FamilyCinema) off() {
	f.t.offTv()
	f.s.off()
	f.p.off()
}

func main() {
	family := &FamilyCinema{}
	family.look()
	fmt.Println("看完电视")
	family.off()
}
