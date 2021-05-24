package copy

import (
	"errors"
	"reflect"

	"github.com/hyahm/golog"
)

var ErrOutValueNotPointer = errors.New("output pointer not a pointer")
var ErrNotSupport = errors.New("not support struct")
var ErrCanNotConvert = errors.New("can not convert struct")
var ErrCanNotSet = errors.New("can not set value")

// 将前面的参数的值复制给后面的参数
func CopyValue(input, output interface{}) error {
	// 后面个输出值必须是指针

	ip := reflect.TypeOf(input)
	rty := reflect.TypeOf(output)

	if rty.Kind() != reflect.Ptr {
		golog.Error(ErrOutValueNotPointer)
		return ErrOutValueNotPointer
	}
	rte := rty.Elem()
	in := reflect.ValueOf(input)
	ot := reflect.ValueOf(&output)
	// if !ot.CanSet() {
	// 	golog.Error(ErrCanNotSet)
	// 	return ErrCanNotSet
	// }
	// 如果值和获取的值都是指针
	if ip.Kind() == reflect.Ptr {
		golog.Info("aaaa")
		// 如果前面的值是指针，
		ipe := ip.Elem()
		if rte.Kind() == ipe.Kind() {
			// 如果2个类型的值是一样的
		}
		// 判断类型是否是一样的
		// if reflect.DeepEqual(ipe, rte) {
		// 	golog.Info("aaaaaa")
		// }
	} else {
		golog.Info("aaaa")
		// 如果前面的值不是指针，
		if rte.Kind() == ip.Kind() {

			// 如果2个类型的值是一样的

			inputKind := rte.Kind()
			switch inputKind {
			case reflect.Bool, reflect.String,
				reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64,
				reflect.Uintptr,
				reflect.Float32,
				reflect.Float64,
				reflect.Complex64,
				reflect.Complex128:
				ot.Set(in.Addr())
				return nil

			case reflect.Array:
			case reflect.Chan:
			case reflect.Func:
			case reflect.Interface:
			case reflect.Map:
			case reflect.Ptr:
			case reflect.Slice:

			case reflect.Struct:
			case reflect.UnsafePointer:
			}
		} else {
			golog.Info("aaaaa")
			outputKind := rte.Kind()
			switch outputKind {
			case reflect.Int:
				golog.Info("aaaaa")
				// 输出是int， 那么下面这些类型都是可以转为int的
				switch ip.Kind() {
				case reflect.Int8:
					ot.Set(in)
				case reflect.Int16:
				case reflect.Int32:
				case reflect.Int64:
				case reflect.Uint:
				case reflect.Uint8:
				case reflect.Uint16:
				case reflect.Uint32:
				case reflect.Uint64:
				case reflect.Uintptr:
				case reflect.Float32:
				case reflect.Float64:
				case reflect.Complex64:
				case reflect.Complex128:
				case reflect.String:
				}
			case reflect.Int8:
			case reflect.Int16:
			case reflect.Int32:
			case reflect.Int64:
			case reflect.Uint:
			case reflect.Uint8:
			case reflect.Uint16:
			case reflect.Uint32:
			case reflect.Uint64:
			case reflect.Uintptr:
			case reflect.Float32:
			case reflect.Float64:
			case reflect.Complex64:
			case reflect.Complex128:
			case reflect.Array:
			case reflect.Chan:
			case reflect.Func:
			case reflect.Interface:
			case reflect.Map:
			case reflect.Ptr:
			case reflect.Slice:
			case reflect.String:
			case reflect.Struct:
			case reflect.UnsafePointer:
			}
		}

		// if reflect.DeepEqual(ip, rte) {
		// 	golog.Info("aaaaaa")
		// }
	}
	// op := rty.Elem()

	return nil
}

// case reflect.Int:
// case reflect.Int8:
// case reflect.Int16:
// case reflect.Int32:
// case reflect.Int64:
// case reflect.Uint:
// case reflect.Uint8:
// case reflect.Uint16:
// case reflect.Uint32:
// case reflect.Uint64:
// case reflect.Uintptr:
// case reflect.Float32:
// case reflect.Float64:
// case reflect.Complex64:
// case reflect.Complex128:
// case reflect.Array:
// case reflect.Chan:
// case reflect.Func:
// case reflect.Interface:
// case reflect.Map:
// case reflect.Ptr:
// case reflect.Slice:
// case reflect.String:
// case reflect.Struct:
// case reflect.UnsafePointer:
