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

/*
#include <stdlib.h>
#include <iup.h>

void _IupSetfAttributeId2(Ihandle *ih, const char *name, int lin, int col, const char *value) {
	IupSetfAttributeId2(ih, name, lin, col, value);
}
*/
import "C"
import (
	"unsafe"
	"fmt"
	"bytes"
)

func StoreAttribute(ih *Ihandle, name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupStoreAttribute((*C.Ihandle)(ih), cName, cValue)
}

func StoreAttributeId(ih *Ihandle, name string, id int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupStoreAttributeId((*C.Ihandle)(ih), cName, C.int(id), cValue)
}

func SetAttribute(ih *Ihandle, name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupSetAttribute(ih.C(), cName, cValue)
}

func SetAttributeId(ih *Ihandle, name string, id int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupSetAttributeId((*C.Ihandle)(ih), cName, C.int(id), cValue)
}

func SetfAttribute(ih *Ihandle, name, format string, args ...interface{}) {
	StoreAttribute(ih, name, fmt.Sprintf(format, args...))
}

func SetfAttributeId(ih *Ihandle, name string, id int, format string, args ...interface{}) {
	StoreAttributeId(ih, name, id, fmt.Sprintf(format, args...))
}

func SetfAttributeId2(ih *Ihandle, name string, lin int, col int, format string, args ...interface{}) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cValue))

	C._IupSetfAttributeId2((*C.Ihandle)(ih), cName, C.int(lin), C.int(col), cValue)
}

func SetAttributes(ih *Ihandle, values string) {
	cValues := C.CString(values)
	defer C.free(unsafe.Pointer(cValues))

	C.IupSetAttributes((*C.Ihandle)(ih), cValues)
}

func ResetAttribute(ih *Ihandle, name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupResetAttribute((*C.Ihandle)(ih), cName)
}

// Warning: handle_name is ignored
func SetAtt(ih *Ihandle, handle_name string, args ...string) *Ihandle {
	attrs := bytes.NewBufferString("")
	for i := 0; i < len(args); i += 2 {
		if i > 0 {
			attrs.WriteString(",")
		}
		attrs.WriteString(fmt.Sprintf("%s=\"%s\"", args[i], args[i+1]))
	}

	SetAttributes(ih, attrs.String())

	return ih
}

// This method does not exist in C Iup. It has been provided as a convience function to allow
// code such as:
//
//     box := iup.Hbox(button1, button2).SetAttrs("GAP", "5", "MARGIN", "8x8")
//
// C Iup provides SetAtt for this purpose but in Go Iup SetAttrs is an easier method to
// accomplish this task due to no necessity of handle_name.
func SetAttrs(ih *Ihandle, args ...string) *Ihandle {
	return SetAtt(ih, "", args...)
}

func SetAttributeHandle(ih *Ihandle, name string, ihNamed *Ihandle) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.IupSetAttributeHandle((*C.Ihandle)(ih), cName, (*C.Ihandle)(ihNamed))
}

func GetAttributeHandle(ih *Ihandle, name string) *Ihandle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return (*Ihandle)(C.IupGetAttributeHandle((*C.Ihandle)(ih), cName))
}

func GetAttribute(ih *Ihandle, name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.GoString(C.IupGetAttribute((*C.Ihandle)(ih), cName))
}

func GetAttributeId(ih *Ihandle, name string, id int) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.GoString(C.IupGetAttributeId((*C.Ihandle)(ih), cName, C.int(id)))
}

func GetIntId(ih *Ihandle, name string, id int) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int(C.IupGetIntId((*C.Ihandle)(ih), cName, C.int(id)))
}

func GetFloatId(ih *Ihandle, name string, id int) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.IupGetFloatId((*C.Ihandle)(ih), cName, C.int(id)))
}

func GetAttributes(ih *Ihandle) string {
	return C.GoString(C.IupGetAttributes((*C.Ihandle)(ih)))
}

func GetFloat(ih *Ihandle, name string) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return float64(C.IupGetFloat((*C.Ihandle)(ih), cName))
}

func GetInt(ih *Ihandle, name string) int64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return int64(C.IupGetInt((*C.Ihandle)(ih), cName))
}

func StoreGlobal(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupStoreGlobal(cName, cValue)
}

func SetGlobal(name string, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.IupSetGlobal(cName, cValue)
}

func GetGlobal(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return C.GoString(C.IupGetGlobal(cName))
}

func GetHandle(name string) *Ihandle {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	return (*Ihandle)(C.IupGetHandle(cName))
}

