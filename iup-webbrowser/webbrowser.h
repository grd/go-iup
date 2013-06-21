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

#include <iup.h>

extern const char* GO_COMPLETED_CB;
extern const char* GO_ERROR_CB;
extern const char* GO_NAVIGATE_CB;
extern const char* GO_NEWWINDOW_CB;

extern void goIupSetCompletedFunc(Ihandle *ih, void *f);
extern void goIupSetErrorFunc(Ihandle *ih, void *f);
extern void goIupSetNavigateFunc(Ihandle *ih, void *f);
extern void goIupSetNewWindowFunc(Ihandle *ih, void *f);

