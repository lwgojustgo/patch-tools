package xunsafe

import (
	"reflect"
	"unsafe"
)

func TestPanicFunc() {
	a := "test panic"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&a))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	b := *(*[]byte)(unsafe.Pointer(&bh))
	b[0] = 'H'
}
