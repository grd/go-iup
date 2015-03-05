/*
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-  If not, see <http://www.gnu.org/licenses/>.
*/

package pplot

/*
#cgo LDFLAGS: -liup -liupcontrols -liupcd -liup_pplot -lcd
#cgo windows LDFLAGS: -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
#include <iupcontrols.h>
#include <iup_pplot.h>
*/
import "C"

import (
	"unsafe"
	"github.com/grd/iup"
)

/*
 Duplicate utility functions until C types can be exported in Go
*/

func stringArrayToC(strs []string) []*C.char {
	max := len(strs)
	result := make([]*C.char, max+1)

	for k, v := range strs {
		result[k] = C.CString(v)
	}
	result[max] = nil

	return result
}

func freeCStringArray(strs []*C.char) {
	for _, v := range strs {
		if v != nil {
			C.free(unsafe.Pointer(v))
		}
	}
}

func float64ArrayToC(nums []float64) []C.float {
	result := make([]C.float, len(nums))

	for k, v := range nums {
		result[k] = C.float(v)
	}

	return result
}

/*
 Actual library code
*/

var libOpened = false

func Open(opts ...interface{}) *iup.Ihandle {
	iup.OpenControlLib()

	if libOpened == false {
		C.IupPPlotOpen()
		libOpened = true
	}

	ih := (*iup.Ihandle)(C.IupPPlot())

	for _, o := range opts {
		switch o.(type) {
		default:
			iup.Decorate(ih, o)
		}
	}

	return ih
}

func Begin(ih *iup.Ihandle, strXdata int) {
	C.IupPPlotBegin((*C.Ihandle)(ih), C.int(strXdata))
}

func Add(ih *iup.Ihandle, x, y float64) {
	C.IupPPlotAdd((*C.Ihandle)(ih), C.float(x), C.float(y))
}

func AddString(ih *iup.Ihandle, x string, y float64) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))

	C.IupPPlotAddStr((*C.Ihandle)(ih), cX, C.float(y))
}

func End(ih *iup.Ihandle) {
	C.IupPPlotEnd((*C.Ihandle)(ih))
}

func Insert(ih *iup.Ihandle, index, sample_index int, x, y float64) {
	C.IupPPlotInsert((*C.Ihandle)(ih), C.int(index), C.int(sample_index), C.float(x), C.float(y))
}

func InsertString(ih *iup.Ihandle, index, sample_index int, x string, y float64) {
	cX := C.CString(x)
	defer C.free(unsafe.Pointer(cX))

	C.IupPPlotInsertStr((*C.Ihandle)(ih), C.int(index), C.int(sample_index), cX, C.float(y))
}

func InsertPoints(ih *iup.Ihandle, index, sample_index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)

	C.IupPPlotInsertPoints((*C.Ihandle)(ih), C.int(index), C.int(sample_index), &cX[0], &cY[0],
		C.int(count))
}

// Differs from InsertPoints as `count' is determined automatically in this case
func InsertStringPoints(ih *iup.Ihandle, index, sample_index int, x []string, y []float64) {
	count := len(x)

	cX := stringArrayToC(x)
	defer freeCStringArray(cX)

	cY := float64ArrayToC(y)

	C.IupPPlotInsertStrPoints((*C.Ihandle)(ih), C.int(index), C.int(sample_index), &cX[0], &cY[0], C.int(count))
}

func AddPoints(ih *iup.Ihandle, index int, x, y []float64) {
	count := len(x)
	cX := float64ArrayToC(x)
	cY := float64ArrayToC(y)

	C.IupPPlotAddPoints((*C.Ihandle)(ih), C.int(index), &cX[0], &cY[0], C.int(count))
}

func AddStringPoints(ih *iup.Ihandle, index int, x []string, y []float64) {
	count := len(x)

	cX := stringArrayToC(x)
	defer freeCStringArray(cX)

	cY := float64ArrayToC(y)

	C.IupPPlotAddStrPoints((*C.Ihandle)(ih), C.int(index), &cX[0], &cY[0], C.int(count))
}

func Transform(ih *iup.Ihandle, x, y float64) (int, int) {
	cIX := new(C.int)
	cIY := new(C.int)

	C.IupPPlotTransform((*C.Ihandle)(ih), C.float(x), C.float(y), cIX, cIY)

	return int(*cIX), int(*cIY)
}
