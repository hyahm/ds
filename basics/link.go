package basics

import (
	"fmt"
	"sync"
)

// 链表
type Link struct {
	Value int
	Prev  *Link
	Next  *Link
	Top   *Top
}

type Top struct {
	Length int
	First  *Link
	End    *Link
	Mu     *sync.RWMutex
}

func NewLinkTop() *Top {
	return &Top{Mu: &sync.RWMutex{}}
}

// 添加到末尾
func (l *Top) Push(value int) *Top {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	// top的长度+1
	new := &Link{
		Value: value,
		Top:   l,
	}
	if l.Length == 0 {
		l.End = new
		l.First = new

	} else {
		new.Prev = l.End
		l.End.Next = new
		l.End = new
	}
	l.Length++
	return l
}

// 添加到开头
func (l *Top) Insert(value int) *Top {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	// top的长度+1
	new := &Link{
		Value: value,
		Top:   l,
		Next:  l.First,
	}
	if l.Length == 0 {
		l.End = new
		l.First = new

	} else {
		l.First.Prev = nil
		l.First = new
		new.Top = l
	}
	l.Length++
	return l
}

// 删除末尾
func (l *Top) PopEnd() (*Top, bool) {
	// top的长度+1
	l.Mu.Lock()
	defer l.Mu.Unlock()
	if l.Length == 0 {
		return l, false
	}
	if l.Length == 1 {
		l.End = nil
		l.First = nil

	} else {
		l.End = l.End.Prev
		l.End.Next = nil
	}
	l.Length--
	return l, true
}

// 添加到开头
func (l *Top) PopStart() (*Top, bool) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	// top的长度+1
	if l.Length == 0 {
		return l, false
	}
	if l.Length == 1 {
		l.End = nil
		l.First = nil

	} else {
		l.First = l.First.Next
		l.First.Prev = nil
	}
	l.Length++
	return l, true
}

func (l *Top) Print() {
	l.Mu.RLock()
	defer l.Mu.RUnlock()
	// top的长度+1
	if l.Length == 0 {
		return
	}
	start := l.First
	for start != nil {
		fmt.Println(start.Value)
		start = start.Next
	}
}
