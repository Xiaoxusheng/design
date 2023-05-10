package main

import (
	"fmt"
	"sync"
)

/*
单例模式
*/

var once sync.Once

type singelton struct{}

var instance *singelton

// GetInstance 只会生成一个单例
func GetInstance() *singelton {

	once.Do(func() {
		instance = new(singelton)
	})

	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	fmt.Printf("内存地址%p\n", s)
	s.SomeThing()
	s1 := GetInstance()
	fmt.Printf("内存地址%p", s1)
}
