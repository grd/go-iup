/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-iup.

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.
*/

package iup

// #include "callback.h"
import "C"

import (
	"unsafe"
)

// Common callbacks

type MapFunc func(*Ihandle) int

//export goIupMapCB
func goIupMapCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*MapFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_MAP_CB)))
	return f((*Ihandle)(h))
}

func SetMapFunc(ih *Ihandle, f MapFunc) {
	C.goIupSetMapFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type UnmapFunc func(*Ihandle) int

//export goIupUnmapCB
func goIupUnmapCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*UnmapFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_UNMAP_CB)))
	return f((*Ihandle)(ih))
}

func SetUnmapFunc(ih *Ihandle, f UnmapFunc) {
	C.goIupSetUnmapFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type DestroyFunc func(*Ihandle) int

//export goIupDestroyCB
func goIupDestroyCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*DestroyFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_DESTROY_CB)))
	return f((*Ihandle)(ih))
}

func SetDestroyFunc(ih *Ihandle, f DestroyFunc) {
	C.goIupSetDestroyFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type GetFocusFunc func(*Ihandle) int

//export goIupGetFocusCB
func goIupGetFocusCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*GetFocusFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_GETFOCUS_CB)))
	return f((*Ihandle)(ih))
}

func SetGetFocusFunc(ih *Ihandle, f GetFocusFunc) {
	C.goIupSetGetFocusFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type KillFocusFunc func(*Ihandle) int

//export goIupKillFocusCB
func goIupKillFocusCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*KillFocusFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_KILLFOCUS_CB)))
	return f((*Ihandle)(ih))
}

func SetKillFocusFunc(ih *Ihandle, f KillFocusFunc) {
	C.goIupSetKillFocusFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type EnterWindowFunc func(*Ihandle) int

//export goIupEnterWindowCB
func goIupEnterWindowCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*EnterWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ENTERWINDOW_CB)))
	return f((*Ihandle)(ih))
}

func SetEnterWindowFunc(ih *Ihandle, f EnterWindowFunc) {
	C.goIupSetEnterWindowFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type LeaveWindowFunc func(*Ihandle) int

//export goIupLeaveWindowCB
func goIupLeaveWindowCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*LeaveWindowFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_LEAVEWINDOW_CB)))
	return f((*Ihandle)(ih))
}

func SetLeaveWindowFunc(ih *Ihandle, f LeaveWindowFunc) {
	C.goIupSetLeaveWindowFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type KAnyFunc func(*Ihandle, int) int

//export goIupKAnyCB
func goIupKAnyCB(ih unsafe.Pointer, c int) int {
	h := (*C.Ihandle)(ih)
	f := *(*KAnyFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_K_ANY_CB)))
	return f((*Ihandle)(ih), int(c))
}

func SetKAnyFunc(ih *Ihandle, f KAnyFunc) {
	C.goIupSetKAnyFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type HelpFunc func(*Ihandle) int

//export goIupHelpCB
func goIupHelpCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*HelpFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_HELP_CB)))
	return f((*Ihandle)(ih))
}

func SetHelpFunc(ih *Ihandle, f HelpFunc) {
	C.goIupSetHelpFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type ButtonFunc func(*Ihandle, int, int, int, int, string) int

//export goIupButtonCB
func goIupButtonCB(ih unsafe.Pointer, button, pressed, x, y int, status unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*ButtonFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_BUTTON_CB)))
	goStatus := C.GoString((*C.char)(status))
	return f((*Ihandle)(ih), button, pressed, x, y, goStatus)
}

func SetButtonFunc(ih *Ihandle, f ButtonFunc) {
	C.goIupSetButtonFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type DropFilesFunc func(*Ihandle, string, int, int, int) int

//export goIupDropFilesCB
func goIupDropFilesCB(ih, filename unsafe.Pointer, num, x, y int)int {
	h := (*C.Ihandle)(ih)
	f := *(*DropFilesFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_DROPFILES_CB)))
	goFilename := C.GoString((*C.char)(filename))
	return f((*Ihandle)(h), goFilename, int(num), int(x), int(y))
}

func SetDropFilesFunc(ih *Ihandle, f DropFilesFunc) {
	C.goIupSetDropFilesFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type ActionFunc func(*Ihandle) int

//export goIupActionCB
func goIupActionCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	f := *(*ActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ACTION)))
	return f((*Ihandle)(ih))
}

func SetActionFunc(ih *Ihandle, f ActionFunc) {
	C.goIupSetActionFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type ListActionFunc func(ih *Ihandle, text string, item, state int) int

//export goIupListActionCB
func goIupListActionCB(ih, text unsafe.Pointer, item, state int) int {
	h := (*C.Ihandle)(ih)
	f := *(*ListActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_LIST_ACTION)))
	goText := C.GoString((*C.char)(text))
	return f((*Ihandle)(ih), goText, int(item), int(state))
}

func SetListActionFunc(ih *Ihandle, f ListActionFunc) {
	C.goIupSetListActionFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type CaretFunc func(ih *Ihandle, lin, col, pos int) int

//export goIupCaretCB
func goIupCaretCB(ih unsafe.Pointer, lin, col, pos int) int {
	h := (*C.Ihandle)(ih)
	f := *(*CaretFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_CARET_CB)))
	return f((*Ihandle)(ih), int(lin), int(col), int(pos))
}

func SetCaretFunc(ih *Ihandle, f CaretFunc) {
	C.goIupSetCaretFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type DblclickFunc func(ih *Ihandle, item int, text string) int

//export goIupDblclickCB
func goIupDblclickCB(ih unsafe.Pointer, item int, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*DblclickFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_DBLCLICK_CB)))
	goText := C.GoString((*C.char)(text))
	return f((*Ihandle)(ih), int(item), goText)
}

func SetDblclickFunc(ih *Ihandle, f DblclickFunc) {
	C.goIupSetDblclickFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type EditFunc func(ih *Ihandle, item int, text string) int

//export goIupEditCB
func goIupEditCB(ih unsafe.Pointer, item int, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*EditFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_EDIT_CB)))
	goText := C.GoString((*C.char)(text))
	return f((*Ihandle)(ih), item, goText)
}

func SetEditFunc(ih *Ihandle, f EditFunc) {
	C.goIupSetEditFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type MotionFunc func(ih *Ihandle, x, y int, status string) int

//export goIupMotionCB
func goIupMotionCB(ih unsafe.Pointer, x, y int, status unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*MotionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_MOTION_CB)))
	goStatus := C.GoString((*C.char)(status))
	return f((*Ihandle)(ih), x, y, goStatus)
}

func SetMotionFunc(ih *Ihandle, f MotionFunc) {
	C.goIupSetMotionFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type MultiselectFunc func(ih *Ihandle, text string) int

//export goIupMultiselectCB
func goIupMultiselectCB(ih, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*MultiselectFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_MULTISELECT_CB)))
	goText := C.GoString((*C.char)(text))
	return f((*Ihandle)(ih), goText)
}

func SetMultiselectFunc(ih *Ihandle, f MultiselectFunc) {
	C.goIupSetMultiselectFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type ValueChangedFunc func(ih *Ihandle) int

//export goIupValueChangedCB
func goIupValueChangedCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*ValueChangedFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_VALUECHANGED_CB)))
	return f((*Ihandle)(ih))
}

func SetValueChangedFunc(ih *Ihandle, f ValueChangedFunc) {
	C.goIupSetValueChangedFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type TextActionFunc func(ih *Ihandle, ch int, newValue string) int

//export goIupTextActionCB
func goIupTextActionCB(ih unsafe.Pointer, ch int, newValue unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*TextActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ACTION)))
	goNewValue := C.GoString((*C.char)(newValue))
	return f((*Ihandle)(ih), ch, goNewValue)
}

func SetTextActionFunc(ih *Ihandle, f TextActionFunc) {
	C.goIupSetTextActionFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type ToggleActionFunc func(ih *Ihandle, state int) int

//export goIupToggleActionCB
func goIupToggleActionCB(ih unsafe.Pointer, state int) int {
	h := (*C.Ihandle)(ih)
	f := *(*ToggleActionFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_ACTION)))
	return f((*Ihandle)(ih), state)
}

func SetToggleActionFunc(ih *Ihandle, f ToggleActionFunc) {
	C.goIupSetToggleActionFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type TabChangeFunc func(ih, new_tab, old_tab *Ihandle) int

//export goIupTabChangeCB
func goIupTabChangeCB(ih, new_tab, old_tab unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	f := *(*TabChangeFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_TABCHANGE_CB)))
	return f((*Ihandle)(ih), (*Ihandle)(new_tab), (*Ihandle)(old_tab))
}

func SetTabChangeFunc(ih *Ihandle, f TabChangeFunc) {
	C.goIupSetTabChangeFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type TabChangePosFunc func(ih *Ihandle, new_pos, old_pos int) int

//export goIupTabChangePosCB
func goIupTabChangePosCB(ih unsafe.Pointer, new_pos, old_pos int) int {
	h := (*C.Ihandle)(ih)
	f := *(*TabChangePosFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_TABCHANGEPOS_CB)))
	return f((*Ihandle)(ih), new_pos, old_pos)
}

func SetTabChangePosFunc(ih *Ihandle, f TabChangePosFunc) {
	C.goIupSetTabChangePosFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type SpinFunc func(ih *Ihandle, inc int) int

//export goIupSpinCB
func goIupSpinCB(ih unsafe.Pointer, inc int) int {
	h := (*C.Ihandle)(ih)
	f := *(*SpinFunc)(unsafe.Pointer(C.IupGetAttribute(h, C.GO_SPIN_CB)))
	return f((*Ihandle)(h), inc)
}

func SetSpinFunc(ih *Ihandle, f SpinFunc) {
	C.goIupSetSpinFunc((*C.Ihandle)(ih), unsafe.Pointer(&f))
}

type IdleFunc func() int

//export goIupIdleCB
func goIupIdleCB() int {
	f := *(*IdleFunc)(unsafe.Pointer(C.IupGetFunction(C.GO_IDLE_ACTION)))
	return f()
}

// The user idle callback function can use a select on input channels and a time.After() to process
// external events without taking too much CPU (idle loop can behave almost like an infinite loop otherwise)
// Example:
// func idleFunc() int {
//     select {
//     case cmd := <-idleChan:
//         // process cmd
//     case <-time.After(time.Duration(150 * time.Millisecond)):
//     }
//     return iup.DEFAULT
// }
func SetIdleFunc(f IdleFunc) {
	C.goIupSetIdleFunc(unsafe.Pointer(&f))
}
