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

#include <stdlib.h>
#include <iup.h>
#include <iupweb.h>
#include "webbrowser.h"
#include "_cgo_export.h"

#define GO_PREFIX        "_GO_"
#define IUP_COMPLETED_CB "COMPLETED_CB"
#define IUP_ERROR_CB     "ERROR_CB"
#define IUP_NAVIGATE_CB  "NAVIGATE_CB"
#define IUP_NEWWINDOW_CB "NEWWINDOW_CB"

const char* GO_COMPLETED_CB = GO_PREFIX IUP_COMPLETED_CB;
const char* GO_ERROR_CB     = GO_PREFIX IUP_ERROR_CB;
const char* GO_NAVIGATE_CB  = GO_PREFIX IUP_NAVIGATE_CB;
const char* GO_NEWWINDOW_CB = GO_PREFIX IUP_NEWWINDOW_CB;

void goIupSetCompletedFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_COMPLETED_CB, f);
	IupSetCallback(ih, IUP_COMPLETED_CB, (Icallback) goIupCompletedCB);
}

void goIupSetErrorFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_ERROR_CB, f);
	IupSetCallback(ih, IUP_ERROR_CB, (Icallback) goIupErrorCB);
}

void goIupSetNavigateFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_NAVIGATE_CB, f);
	IupSetCallback(ih, IUP_NAVIGATE_CB, (Icallback) goIupNavigateCB);
}

void goIupSetNewWindowFunc(Ihandle *ih, void *f) {
	IupSetCallback(ih, GO_NEWWINDOW_CB, f);
	IupSetCallback(ih, IUP_NEWWINDOW_CB, (Icallback) goIupNewWindowCB);
}

