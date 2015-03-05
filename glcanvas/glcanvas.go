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

package glcanvas

/*
#cgo LDFLAGS: -liup -liupcontrols -liupgl
#cgo windows LDFLAGS: -lgdi32 -lole32 -lcomdlg32 -lcomctl32

#include <stdlib.h>
#include <iup.h>
#include <iupgl.h>
*/
import "C"
import "github.com/grd/iup"

func Canvas(opts ...interface{}) *iup.Ihandle {
	iup.OpenControlLib()

	ih := (*iup.Ihandle)(C.IupGLCanvas(nil))

	for _, o := range opts {
		switch o.(type) {
		default:
			iup.Decorate(ih, o)
		}
	}

	return ih
}

func MakeCurrent(ih *iup.Ihandle) {
	C.IupGLMakeCurrent((*C.Ihandle)(ih))
}

func IsCurrent(ih *iup.Ihandle) int {
	return int(C.IupGLIsCurrent((*C.Ihandle)(ih)))
}

func SwapBuffers(ih *iup.Ihandle) {
	C.IupGLSwapBuffers((*C.Ihandle)(ih))
}

func Palette(ih *iup.Ihandle, index int, r, g, b float64) {
	C.IupGLPalette((*C.Ihandle)(ih), C.int(index), C.float(r), C.float(g), C.float(b))
}

func UseFont(ih *iup.Ihandle, first, count, list_base int) {
	C.IupGLUseFont((*C.Ihandle)(ih), C.int(first), C.int(count), C.int(list_base))
}

func Wait(gl int) {
	C.IupGLWait(C.int(gl))
}
