package tsm

// #cgo pkg-config: libtsm
// #include <libtsm.h>
import "C"

import (
	"runtime"
)

type Screen struct {
	s *C.struct_tsm_screen
}

func NewScreen() *Screen {
	s := &Screen{}
	if ret := C.tsm_screen_new(&s.s, nil, nil); ret != 0 {
		panic("tsm: cannot create screen")
	}
	runtime.SetFinalizer(s, func(s *Screen) {
		C.tsm_screen_unref(s.s)
	})
	return s
}
