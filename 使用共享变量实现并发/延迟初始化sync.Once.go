package main

import (
	"fmt"
	"image"
	"sync"
)

func main() {
	fmt.Println("延迟初始化 sync.Once")
}

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades":     loadIcon("spades.png"),
		"hearts.png": loadIcon("heart.png"),
	}
}
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() //一次性的初始化
	}
	return icons[name]
} //并发调用 Icon 时这个模式是不安全的

//在填充数据之前把一个空 map 赋给 icons
func loadIcons1() {
	icons = make(map[string]image.Image)
	icons["spades"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("heart.png")
} //一个 goroutine 发现 icons 不是空并不意味着变量的初始化已经完成

//保证所有 goroutine 都能观察到 loadIcons 效果最简单的正确方法就是用一个互斥锁来做同步
var mu sync.Mutex //保护 icons
var icons map[string]image.Image

//并发安全
func Icons(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons1()
	}
	return icons[name]
}

//采用互斥锁访问 icons 的额外代价就是两个 goroutine 不能并发访问这个变量，
//即使在变量已经安全完成初始化且不再更改的情况下，也会造成这个后果，使用一个可以并发读的锁可以改善
var mu1 sync.RWMutex
var icons1 map[string]image.Image

func Icon(name string) image.Image {
	mu1.RLock()
	if icons1 != nil {
		icon := icons1[name]
	}
	mu1.RUnlock()

	//获取互斥锁
	mu1.Lock()
	if icons1 == nil { //必须重新检测 nil 值
		loadIcons1()
	}
	icon := icons1[name]
	mu1.Unlock()
	return icon
}

//sync.Once   唯一方法 Do 以初始化函数作为它的参数
var loadIconsOnce sync.Once
var icons map[string]image.Image

//并发安全
func Icon1(name string) image.Image {
	loadIconsOnce.Do(loadIcons1)
	return icons[name]
}

//每次调用 Do(loadIcons1) 时会先锁定互斥量并检测里面的布尔变量
//在第一次调用时，这个布尔变量为假，Do 会调用 loadIcons1 然后把变量设为 真
