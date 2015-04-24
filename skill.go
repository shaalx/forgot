package forgot

import (
	"reflect"
	"unsafe"
)

func B2S(buf []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	up := unsafe.Pointer(p.Data)
	return up
}
