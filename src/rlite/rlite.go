package rlite

/*
#cgo CFLAGS: -std=gnu99
#cgo LDFLAGS: -lhirlite
#include <hirlite/hirlite.h>
// I don't know how to cast in go
static inline rliteReply* ptor_pos(rliteReply **p, size_t pos) {
    rliteReply **element = p;
    return element[pos];
}
static inline rliteReply* ptor(void *p) {
    return (rliteReply*)p;
}
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

type Conn struct {
	db *C.rliteContext
}

func cStr(s string) *C.char {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return (*C.char)(unsafe.Pointer(h.Data))
}

func Open(name string) (*Conn, error) {
	name += "\x00"

	db := C.rliteConnect(cStr(name), 0)
	if db == nil {
		return nil, errors.New("Unable to open database")
	}
	return &Conn{db: db}, nil
}

func Close(c *Conn) {
	C.rliteFree(c.db)
}

func GetReply(reply *C.rliteReply) (interface{}, error) {
	if reply._type == C.RLITE_REPLY_ERROR {
		return nil, errors.New(C.GoStringN(reply.str, reply.len))
	}
	if reply._type == C.RLITE_REPLY_STRING || reply._type == C.RLITE_REPLY_STATUS {
		return C.GoStringN(reply.str, reply.len), nil
	}
	if reply._type == C.RLITE_REPLY_INTEGER {
		return int(reply.integer), nil
	}
	if reply._type == C.RLITE_REPLY_NIL {
		return nil, nil
	}
	if reply._type == C.RLITE_REPLY_ARRAY {
		arr := make([]interface{}, reply.elements)
		for i := C.size_t(0); i < reply.elements; i++ {
			// TODO: what if the array has an error?
			arr[i], _ = GetReply(C.ptor_pos(reply.element, i))
		}
		return arr, nil
	}
	return nil, errors.New(fmt.Sprintf("Unknown type %d", reply._type))
}

func Command(c *Conn, list []string) (interface{}, error) {
	var b *C.char
	argvSize := unsafe.Sizeof(b)
	argv := C.malloc(C.size_t(len(list)) * C.size_t(argvSize))
	defer C.free(argv)

	var d *C.size_t
	argvlenSize := unsafe.Sizeof(d)
	argvlen := C.malloc(C.size_t(len(list)) * C.size_t(argvlenSize))

	for i := 0; i < len(list); i++ {
		b := StringToBytes(list[i])
		element := (**C.char)(unsafe.Pointer(uintptr(argv) + uintptr(i)*argvSize))
		*element = (*C.char)(unsafe.Pointer(&b[0]))
		elementlen := (*C.size_t)(unsafe.Pointer(uintptr(argvlen) + uintptr(i)*argvSize))
		*elementlen = C.size_t(len(b))
	}
	p := C.rliteCommandArgv(c.db, C.int(len(list)), (**C.char)(argv), (*C.size_t)(argvlen))
	r := C.ptor(p)
	defer C.rliteFreeReplyObject(p)
	return GetReply(r)
}
