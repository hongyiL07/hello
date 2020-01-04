package main

import (
	"fmt"
	"sync"
	)

//互斥锁模式应用非常广泛，所以 sync 包有一个单独的 Mutex 类型来支持这种模式，
//它的 Lock 方法用于获取令牌（token，此过程也称为 上锁），Unlock 方法用于释放令牌

//互斥锁保护共享变量

//在 Lock 和 Unlock 之间的代码，可以自由的读取和修改共享变量，这一部分称为 临界区域
//重要的是：goroutine 在使用完成后就应当释放锁，另外，需要包括函数的所有分支，特别是错误分支
//
//这种函数 互斥锁 变量的组合方式称为 监控模式

func main()  {
	fmt.Println("互斥锁 sync.Mutex")
}

/*func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}*/  //Unlock 和 return 语句已经读完 balance 变量之后执行，所以Balance函数就是并发安全的

//成功时余额减少指定数量，并返回true；如果余额不足无法完成交易，恢复余额返回 false
func Withdraw(amount int ) bool {
	Deposits(-amount)
	if Balance()<0{
		Deposit(amount)
		return false
	}
	return true
}  //Withdraw 存在的问题在于不是原子操作：它包含三个串行的操作，每个操作都申请并释放了互斥锁。但对于整个序列没有上锁
//无法对一个上锁的互斥量再上锁，会导致死锁，函数将被卡住





//读写 互斥锁   sync.RWMutex
//允许只读操作可以并发执行，但写操作需要获得完全独享的访问权限 称为多读单写锁
var mu sync.RWMutex
var balance int

func Balance() int {
	mu.RLock()   //读锁
	defer mu.RUnlock()
	return balance
}//Balance函数现在可以调用 Rlock 和 RUnlock 方法来分别获取和释放一个读锁（也称为 共享锁）