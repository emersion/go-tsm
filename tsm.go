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

func (s *Screen) Width() uint {
	return uint(C.tsm_screen_get_width(s.s))
}

func (s *Screen) Height() uint {
	return uint(C.tsm_screen_get_height(s.s))
}

func (s *Screen) Resize(w, h uint) {
	if ret := C.tsm_screen_resize(s.s, C.uint(w), C.uint(h)); ret != 0 {
		panic("tsm: cannot resize screen")
	}
}

func (s *Screen) SetMargins(top, bottom uint) {
	if ret := C.tsm_screen_resize(s.s, C.uint(top), C.uint(bottom)); ret != 0 {
		panic("tsm: cannot set screen margins")
	}
}

func (s *Screen) SetMaxScrollback(max uint) {
	C.tsm_screen_set_max_sb(s.s, C.uint(max))
}

func (s *Screen) ClearScrollback() {
	C.tsm_screen_clear_sb(s.s)
}

func (s *Screen) ScrollbackUp(num uint) {
	C.tsm_screen_sb_up(s.s, C.uint(num))
}

func (s *Screen) ScrollbackDown(num uint) {
	C.tsm_screen_sb_down(s.s, C.uint(num))
}

func (s *Screen) ScrollbackPageUp(num uint) {
	C.tsm_screen_sb_page_up(s.s, C.uint(num))
}

func (s *Screen) ScrollbackPageDown(num uint) {
	C.tsm_screen_sb_page_down(s.s, C.uint(num))
}

func (s *Screen) ScrollbackReset() {
	C.tsm_screen_sb_reset(s.s)
}
