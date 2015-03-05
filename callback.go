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

/*Callbacks are stored in slices. The index of the callback is saved in an IUP attribute with prefix 
  _GO_ in front of the action name
*/

// This slice stores all callbacks that have a signature of: func(*Ihandle) int
var callbackfuncs []func(*Ihandle) int

type MapFunc func(*Ihandle) int

//export goIupMapCB
func goIupMapCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_MAP_CB))
	return callbackfuncs[i]((*Ihandle)(h))
}

func SetMapFunc(ih *Ihandle, f MapFunc) {
	C.goIupSetMapFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type UnmapFunc func(*Ihandle) int

//export goIupUnmapCB
func goIupUnmapCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_UNMAP_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetUnmapFunc(ih *Ihandle, f UnmapFunc) {
	C.goIupSetUnmapFunc((*C.Ihandle)(ih),  C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type DestroyFunc func(*Ihandle) int

//export goIupDestroyCB
func goIupDestroyCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_DESTROY_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetDestroyFunc(ih *Ihandle, f DestroyFunc) {
	C.goIupSetDestroyFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type GetFocusFunc func(*Ihandle) int

//export goIupGetFocusCB
func goIupGetFocusCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_GETFOCUS_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetGetFocusFunc(ih *Ihandle, f GetFocusFunc) {
	C.goIupSetGetFocusFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type KillFocusFunc func(*Ihandle) int

//export goIupKillFocusCB
func goIupKillFocusCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_KILLFOCUS_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetKillFocusFunc(ih *Ihandle, f KillFocusFunc) {
	C.goIupSetKillFocusFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type EnterWindowFunc func(*Ihandle) int

//export goIupEnterWindowCB
func goIupEnterWindowCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_ENTERWINDOW_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetEnterWindowFunc(ih *Ihandle, f EnterWindowFunc) {
	C.goIupSetEnterWindowFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type LeaveWindowFunc func(*Ihandle) int

//export goIupLeaveWindowCB
func goIupLeaveWindowCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_LEAVEWINDOW_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetLeaveWindowFunc(ih *Ihandle, f LeaveWindowFunc) {
	C.goIupSetLeaveWindowFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type KAnyFunc func(*Ihandle, int) int
var kanyfuncs []KAnyFunc

//export goIupKAnyCB
func goIupKAnyCB(ih unsafe.Pointer, c int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_K_ANY_CB))
	return kanyfuncs[i]((*Ihandle)(ih), int(c))
}

func SetKAnyFunc(ih *Ihandle, f KAnyFunc) {
	C.goIupSetKAnyFunc((*C.Ihandle)(ih), C.int(len(kanyfuncs)))
	kanyfuncs = append(kanyfuncs, f)
}

type HelpFunc func(*Ihandle) int

//export goIupHelpCB
func goIupHelpCB(ih unsafe.Pointer)int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_HELP_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetHelpFunc(ih *Ihandle, f HelpFunc) {
	C.goIupSetHelpFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type ButtonFunc func(*Ihandle, int, int, int, int, string) int
var buttonfuncs []ButtonFunc

//export goIupButtonCB
func goIupButtonCB(ih unsafe.Pointer, button, pressed, x, y int, status unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_BUTTON_CB))
	goStatus := C.GoString((*C.char)(status))
	return buttonfuncs[i]((*Ihandle)(ih), button, pressed, x, y, goStatus)
}

func SetButtonFunc(ih *Ihandle, f ButtonFunc) {
	C.goIupSetButtonFunc((*C.Ihandle)(ih), C.int(len(buttonfuncs)))
	buttonfuncs = append(buttonfuncs, f)
}

type DropFilesFunc func(*Ihandle, string, int, int, int) int
var dropfilesfuncs []DropFilesFunc

//export goIupDropFilesCB
func goIupDropFilesCB(ih, filename unsafe.Pointer, num, x, y int)int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_DROPFILES_CB))
	goFilename := C.GoString((*C.char)(filename))
	return dropfilesfuncs[i]((*Ihandle)(h), goFilename, int(num), int(x), int(y))
}

func SetDropFilesFunc(ih *Ihandle, f DropFilesFunc) {
	C.goIupSetDropFilesFunc((*C.Ihandle)(ih), C.int(len(dropfilesfuncs)))
	dropfilesfuncs = append(dropfilesfuncs, f)
}

type ActionFunc func(*Ihandle) int

//export goIupActionCB
func goIupActionCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_ACTION))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetActionFunc(ih *Ihandle, f ActionFunc) {	
	C.goIupSetActionFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type ListActionFunc func(ih *Ihandle, text string, item, state int) int
var listactionfuncs []ListActionFunc

//export goIupListActionCB
func goIupListActionCB(ih, text unsafe.Pointer, item, state int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_LIST_ACTION))
	goText := C.GoString((*C.char)(text))
	return listactionfuncs[i]((*Ihandle)(ih), goText, int(item), int(state))
}

func SetListActionFunc(ih *Ihandle, f ListActionFunc) {
	C.goIupSetListActionFunc((*C.Ihandle)(ih), C.int(len(listactionfuncs)))
	listactionfuncs = append(listactionfuncs, f)
}

type CaretFunc func(ih *Ihandle, lin, col, pos int) int
var caretfuncs []CaretFunc

//export goIupCaretCB
func goIupCaretCB(ih unsafe.Pointer, lin, col, pos int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_CARET_CB))
	return caretfuncs[i]((*Ihandle)(ih), int(lin), int(col), int(pos))
}

func SetCaretFunc(ih *Ihandle, f CaretFunc) {
	C.goIupSetCaretFunc((*C.Ihandle)(ih), C.int(len(caretfuncs)))
	caretfuncs = append(caretfuncs, f)
}

type DblclickFunc func(ih *Ihandle, item int, text string) int
var dblclickfuncs []DblclickFunc

//export goIupDblclickCB
func goIupDblclickCB(ih unsafe.Pointer, item int, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_DBLCLICK_CB))
	goText := C.GoString((*C.char)(text))
	return dblclickfuncs[i]((*Ihandle)(ih), int(item), goText)
}

func SetDblclickFunc(ih *Ihandle, f DblclickFunc) {
	C.goIupSetDblclickFunc((*C.Ihandle)(ih), C.int(len(dblclickfuncs)))
	dblclickfuncs = append(dblclickfuncs, f)
}

type EditFunc func(ih *Ihandle, item int, text string) int
var editfuncs []EditFunc

//export goIupEditCB
func goIupEditCB(ih unsafe.Pointer, item int, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_EDIT_CB))
	goText := C.GoString((*C.char)(text))
	return editfuncs[i]((*Ihandle)(ih), item, goText)
}

func SetEditFunc(ih *Ihandle, f EditFunc) {
	C.goIupSetEditFunc((*C.Ihandle)(ih), C.int(len(editfuncs)))
	editfuncs = append(editfuncs, f)
}

type MotionFunc func(ih *Ihandle, x, y int, status string) int
var motionfuncs []MotionFunc

//export goIupMotionCB
func goIupMotionCB(ih unsafe.Pointer, x, y int, status unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_MOTION_CB))
	goStatus := C.GoString((*C.char)(status))
	return motionfuncs[i]((*Ihandle)(ih), x, y, goStatus)
}

func SetMotionFunc(ih *Ihandle, f MotionFunc) {
	C.goIupSetMotionFunc((*C.Ihandle)(ih), C.int(len(motionfuncs)))
	motionfuncs = append(motionfuncs, f)
}

type MultiselectFunc func(ih *Ihandle, text string) int
var multiselectfuncs []MultiselectFunc

//export goIupMultiselectCB
func goIupMultiselectCB(ih, text unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_MULTISELECT_CB))
	goText := C.GoString((*C.char)(text))
	return multiselectfuncs[i]((*Ihandle)(ih), goText)
}

func SetMultiselectFunc(ih *Ihandle, f MultiselectFunc) {
	C.goIupSetMultiselectFunc((*C.Ihandle)(ih), C.int(len(multiselectfuncs)))
	multiselectfuncs = append(multiselectfuncs, f)
}

type ValueChangedFunc func(ih *Ihandle) int

//export goIupValueChangedCB
func goIupValueChangedCB(ih unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_VALUECHANGED_CB))
	return callbackfuncs[i]((*Ihandle)(ih))
}

func SetValueChangedFunc(ih *Ihandle, f ValueChangedFunc) {
	C.goIupSetValueChangedFunc((*C.Ihandle)(ih), C.int(len(callbackfuncs)))
	callbackfuncs = append(callbackfuncs, f)
}

type TextActionFunc func(ih *Ihandle, ch int, newValue string) int
var textactionfuncs []TextActionFunc

//export goIupTextActionCB
func goIupTextActionCB(ih unsafe.Pointer, ch int, newValue unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_ACTION))
	goNewValue := C.GoString((*C.char)(newValue))
	return textactionfuncs[i]((*Ihandle)(ih), ch, goNewValue)
}

func SetTextActionFunc(ih *Ihandle, f TextActionFunc) {
	C.goIupSetTextActionFunc((*C.Ihandle)(ih), C.int(len(textactionfuncs)))
	textactionfuncs = append(textactionfuncs, f)
}

type ToggleActionFunc func(ih *Ihandle, state int) int
var toggleactionfuncs []ToggleActionFunc

//export goIupToggleActionCB
func goIupToggleActionCB(ih unsafe.Pointer, state int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_ACTION))
	return toggleactionfuncs[i]((*Ihandle)(ih), state)
}

func SetToggleActionFunc(ih *Ihandle, f ToggleActionFunc) {
	C.goIupSetToggleActionFunc((*C.Ihandle)(ih), C.int(len(toggleactionfuncs)))
	toggleactionfuncs = append(toggleactionfuncs, f)
}

type TabChangeFunc func(ih, new_tab, old_tab *Ihandle) int
var tabchangefuncs []TabChangeFunc

//export goIupTabChangeCB
func goIupTabChangeCB(ih, new_tab, old_tab unsafe.Pointer) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_TABCHANGE_CB))
	return tabchangefuncs[i]((*Ihandle)(ih), (*Ihandle)(new_tab), (*Ihandle)(old_tab))
}

func SetTabChangeFunc(ih *Ihandle, f TabChangeFunc) {
	C.goIupSetTabChangeFunc((*C.Ihandle)(ih), C.int(len(tabchangefuncs)))
	tabchangefuncs = append(tabchangefuncs, f)
}

type TabChangePosFunc func(ih *Ihandle, new_pos, old_pos int) int
var tabchangeposfuncs []TabChangePosFunc

//export goIupTabChangePosCB
func goIupTabChangePosCB(ih unsafe.Pointer, new_pos, old_pos int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_TABCHANGEPOS_CB))
	return tabchangeposfuncs[i]((*Ihandle)(ih), new_pos, old_pos)
}

func SetTabChangePosFunc(ih *Ihandle, f TabChangePosFunc) {
	C.goIupSetTabChangePosFunc((*C.Ihandle)(ih), C.int(len(tabchangeposfuncs)))
	tabchangeposfuncs = append(tabchangeposfuncs, f)
}

type SpinFunc func(ih *Ihandle, inc int) int
var spinfuncs []SpinFunc

//export goIupSpinCB
func goIupSpinCB(ih unsafe.Pointer, inc int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_SPIN_CB))
	return spinfuncs[i]((*Ihandle)(h), inc)
}

func SetSpinFunc(ih *Ihandle, f SpinFunc) {
	C.goIupSetSpinFunc((*C.Ihandle)(ih), C.int(len(spinfuncs)))
	spinfuncs = append(spinfuncs, f)
}

type ShowFunc func(ih *Ihandle, state int) int
var showfuncs []ShowFunc

//export goIupShowCB
func goIupShowCB(ih unsafe.Pointer, state int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_SHOW_CB))
	return showfuncs[i]((*Ihandle)(h), state)
}

func SetShowFunc(ih *Ihandle, f ShowFunc) {
	C.goIupSetShowFunc((*C.Ihandle)(ih), C.int(len(showfuncs)))
	showfuncs = append(showfuncs, f)
}

type TrayClickFunc func(ih *Ihandle, but, pressed, dclick int) int
var trayclickfuncs []TrayClickFunc

//export goIupTrayClickCB
func goIupTrayClickCB(ih unsafe.Pointer, but, pressed, dclick int) int {
	h := (*C.Ihandle)(ih)
	i := int(C.IupGetInt(h, C.GO_TRAYCLICK_CB))
	return trayclickfuncs[i]((*Ihandle)(h), but, pressed, dclick)
}

func SetTrayClickFunc(ih *Ihandle, f TrayClickFunc) {
	C.goIupSetTrayClickFunc((*C.Ihandle)(ih), C.int(len(trayclickfuncs)))
	trayclickfuncs = append(trayclickfuncs, f)
}

type IdleFunc func() int
var idlefunc IdleFunc

//export goIupIdleCB
func goIupIdleCB() int {
	return idlefunc()
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
	idlefunc = f
	C.goIupSetIdleFunc()
}
