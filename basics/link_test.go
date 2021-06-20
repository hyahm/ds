package basics

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T) {
	top := NewLinkTop()
	top.Insert(5)
	top.Insert(6)
	top.Insert(7)
	top.Push(5555)
	top.Push(4444)
	top.Push(7777)
	top.Insert(7)
	top.Insert(9)
	top.Print()
	fmt.Println("---------------------")
	top.PopEnd()
	top.PopStart()
	top.Print()
}
