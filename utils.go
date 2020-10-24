package xorm_go2sky_hook

import (
	"reflect"
	"regexp"
	"unsafe"
)

var (
	regIdNumber = regexp.MustCompile(`([1-9])([0-7]\d{4})(19[0-9][0-9]|20[0-3][0-9])\d{4}(\d{3})(\d|X|x)`)
	regPhone    = regexp.MustCompile("(1\\d{2})(\\d{4})(\\d{4})")
)

const (
	replaceIdNumberStr = "$1$2********$4$5"
	replacePhoneStr    = "$1****$3"
	tempStr            = "#@$12$236+96"
)

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
