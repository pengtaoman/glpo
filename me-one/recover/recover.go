package main

import (
	"errors"
	"fmt"
)

/**
recover仅在defer中调用
获取panic的值
如果无法处理 可重新panic

https://segmentfault.com/a/1190000021141276

Go 的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起painc异常。
defer是一个面向编译器的声明，他会让编译器做两件事：

编译器会将defer声明编译为runtime.deferproc(fn)，这样运行时，会调用runtime.deferproc，在deferproc中将所有defer挂到goroutine的defer链上；
编译器会在函数return之前（注意，是return之前，而不是return xxx之前，后者不是一条原子指令），增加runtime.deferreturn调用；这样运行时，开始处理前面挂在defer链上的所有defer。

你不应该试图去恢复其他包引起的panic。公有的API应该将函数的运行失败作为error返回，而不是panic。同样的，你也不应该恢复一个由他人开发的函数引起的panic，比如说调用者传入的回调函数，因为你无法确保这样做是安全的。
/////////////////////////////////////////////
///////////////////////////////////////////////
panic() 执行后，后续语句不再执行，会先调用当前协程的 defer 链表。
如果某个 goroutine 的 defer 没有 recover，会终止整个程序（exit(2)），不仅仅是终止当前 goroutine 。
如发现 defer 函数包含 recover, 则会运行 recovery 函数，recovery 会跳转到 deferreturn 。
panic 被 recover 后，会影响到当前函数中的后续语句的执行，但不影响当前 goroutine 的继续执行。
recover() 的作用是捕获异常之后让程序正常往下执行而不会退出。
recover() 必须写在defer语块中才能生效。
recover() 的作用范围仅限于当前的所属 go routine。发生 panic 时只会执行当前协程中的defer函数，其它协程里面的 defer 不会执行。
如果要查找堆栈跟踪，请使用在 Debug 程序包下定义的 PrintStack 函数。
既然你要使用 panic，那为什么要 recover ？你的期望是什么？如果不希望 go die 为什么要用 panic ？
如果实在每个 panic 都想捕获，可以考虑把 panic 这样的事件通知给其他 goroutine 处理。
*/

/**
panic
停止当前函数的执行
一直向上返回，执行每一层的defer
如果没有遇见recover，直接退出
*/
func tryRecover() {
	defer func() {
		//recover 返回的是interface{}, 可以是任何类型
		println("进入defer！！！！")
		r := recover()
		//r.(error) 强制r是error？？？？
		if err, ok := r.(error); ok {
			fmt.Println("Recover错误爆出::", err)
		} else {
			//panic(r)
		}
	}()
	//panic
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	//panic(errors.New("这个panic错误！！！"))
	//panic("非错误的panic！！！！")
	panic(errors.New("ERROR错误的panic！！！！"))
	println("panic 之后的")
}

func main() {

	tryRecover()
}
