package copy

import (
	"testing"

	"github.com/hyahm/golog"
)

type A struct {
	Age  int
	Name string
}

func TestString(t *testing.T) {
	defer golog.Sync()
	var a int8 = 77
	var b int
	err := CopyValue(a, &b)
	if err != nil {
		t.Fatal(err)
	}
	golog.Info(b)
}

func TestInt(t *testing.T) {
	defer golog.Sync()
	var a int32 = 1<<31 - 1
	var b int32 = -1 << 31
	golog.Info(a)
	golog.Info(b)
}
