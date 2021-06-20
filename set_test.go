package ds

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add(6)
	set.Add(8)
	set.Add(2)
	set.Add(2)
	set.Add(2)
	set.Add(3)
	fmt.Println(set.Get())
}
