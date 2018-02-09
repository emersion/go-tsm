package tsm

/*
#cgo pkg-config: libtsm

#include <libtsm.h>

void set_screen_attr_bitfield(struct tsm_screen_attr *attr, bool bold,
		bool underline, bool inverse, bool protect, bool blink) {
	attr->bold = bold;
	attr->underline = underline;
	attr->inverse = inverse;
	attr->protect = protect;
	attr->blink = blink;
}
*/
import "C"

import (
	"runtime"
)

type ScreenAttr struct {
	FCCode, BCCode int8 // foreground, background color codes (<0 for rgb)
	FR, FG, FB, BR, BG, BB uint8 // foreground/background red/green/blue
	Bold, Underline, Inverse, Protect, Blink bool
}

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

func (s *Screen) SetDefAttr(attr *ScreenAttr) {
	cattr := C.struct_tsm_screen_attr{
		fccode: C.int8_t(attr.FCCode),
		bccode: C.int8_t(attr.BCCode),
		fr: C.uint8_t(attr.FR),
		fg: C.uint8_t(attr.FG),
		fb: C.uint8_t(attr.FB),
		br: C.uint8_t(attr.BR),
		bg: C.uint8_t(attr.BG),
		bb: C.uint8_t(attr.BB),
	}
	C.set_screen_attr_bitfield(&cattr, C.bool(attr.Bold), C.bool(attr.Underline), C.bool(attr.Inverse), C.bool(attr.Protect), C.bool(attr.Blink))

	C.tsm_screen_set_def_attr(s.s, &cattr)
}

func (s *Screen) Reset() {
	C.tsm_screen_reset(s.s)
}

func (s *Screen) SetFlags(flags uint) {
	C.tsm_screen_set_flags(s.s, C.uint(flags))
}

func (s *Screen) ResetFlags(flags uint) {
	C.tsm_screen_reset_flags(s.s, C.uint(flags))
}

func (s *Screen) Flags() uint {
	return uint(C.tsm_screen_get_flags(s.s))
}
