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
	t.Log(b)

}
