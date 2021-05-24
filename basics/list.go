package basics

import "errors"

// https://www.cs.usfca.edu/~galles/visualization/StackArray.html

var ErrFull = errors.New("array Full")
var ErrEmpty = errors.New("array is Empty")
var ErrNeedNew = errors.New("array must be created by NewArray")

// 使用原生的list 即可
type Array struct {
	len int
	top int
	new bool
	arr []interface{}
}

func NewArray(length int) *Array {
	return &Array{
		len: length,
		new: true,
		arr: make([]interface{}, length),
	}
}

// 添加到末尾
func (arr *Array) Push(value interface{}) error {
	arr.check()
	if arr.top == arr.len {
		return ErrFull
	}
	arr.arr[arr.top] = value
	arr.top++
	arr.len++
	return nil
}

// 删除末尾元素
func (arr *Array) Pop() error {
	arr.check()
	if arr.len == 0 {
		return ErrEmpty
	}
	arr.top--
	arr.len--

	arr.arr[arr.top] = nil
	return nil
}

func (arr *Array) check() error {
	if !arr.new {
		return ErrNeedNew
	}
	return nil
}
