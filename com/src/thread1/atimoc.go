package thread1

import (
	"fmt"
	"sync/atomic"
)

/*
原子操作，也是协程加锁的一种方式
我们除了sync.Mutex这种方式可以保证我们多个协程不会出现冲突外
我们还可以使用原则操作，导入sync/atomic这个包，通过atomic.AddInt32/64(&变量,变化量)
*/

/*
//方式一：利用Mutex互斥锁实现同步
var lock sync.Mutex
var mony int = 100

func ADD3() {
	lock.Lock()
	mony++
	lock.Unlock()
}
func Sub3() {
	lock.Lock()
	mony--
	lock.Unlock()
}
*/

// 方式二：通过原子操作实现协程同步 导入“sync/atomic”
// cas compare and swap:比较并且交换，就是在添加之前当前这个值要和以前那个值进行比较一下是否一样，一样才进行添加，否则认为添加失败
func ADD3(my *int32) {
	atomic.AddInt32(my, 1)
	fmt.Println("版本提交v2")
	fmt.Println("版本提交v3")
}
func Sub3(my *int32) {
	atomic.AddInt32(my, -1)
}
