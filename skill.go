package forgot

import (
	"reflect"
	"unsafe"
)

func B2S(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func S2B(s *string) []byte {
	v := (*reflect.SliceHeader)(unsafe.Pointer(s))
	val := *(*[]byte)(unsafe.Pointer((v)))
	return val
}
