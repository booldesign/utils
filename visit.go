package utils

import "sync/atomic"

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/31 15:40
 * @Desc: 原子计数器
 */

//记录访问次数
var Visits uint64

//加1
func Increment() uint64 {
	//原子级的加一操作
	return atomic.AddUint64(&Visits, 1)
}

//减一
func Decrement() uint64 {
	// 原子级的减一操作
	//^uint64(0) 先将 0 转换成 uint64，再
	// 按位取反。
	// 另外，用 x 减去一个带符号的正常数值 c
	// 用 AddUint64(&x, ^uint64(c-1))
	return atomic.AddUint64(&Visits, ^uint64(0))
}
