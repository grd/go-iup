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
#include "callback.h"
#include "_cgo_export.h"


const char *GO_MAP_CB = GO_PREFIX IUP_MAP_CB;
const char *GO_UNMAP_CB = GO_PREFIX IUP_UNMAP_CB;
const char *GO_DESTROY_CB = GO_PREFIX IUP_DESTROY_CB;
const char *GO_GETFOCUS_CB = GO_PREFIX IUP_GETFOCUS_CB;
const char *GO_KILLFOCUS_CB = GO_PREFIX IUP_KILLFOCUS_CB;
const char *GO_ENTERWINDOW_CB = GO_PREFIX IUP_ENTERWINDOW_CB;
const char *GO_LEAVEWINDOW_CB = GO_PREFIX IUP_LEAVEWINDOW_CB;
const char *GO_K_ANY_CB = GO_PREFIX IUP_K_ANY;
const char *GO_HELP_CB = GO_PREFIX IUP_HELP_CB;
const char *GO_BUTTON_CB = GO_PREFIX IUP_BUTTON_CB;
const char *GO_DROPFILES_CB = GO_PREFIX IUP_DROPFILES_CB;
const char *GO_ACTION = GO_PREFIX IUP_ACTION;
const char *GO_LIST_ACTION = GO_PREFIX IUP_ACTION;
const char *GO_CARET_CB = GO_PREFIX IUP_CARET_CB;
const char *GO_DBLCLICK_CB = GO_PREFIX IUP_DBLCLICK_CB;
const char *GO_EDIT_CB = GO_PREFIX IUP_EDIT_CB;
const char *GO_MOTION_CB = GO_PREFIX IUP_MOTION_CB;
const char *GO_MULTISELECT_CB = GO_PREFIX IUP_MULTISELECT_CB;
const char *GO_VALUECHANGED_CB = GO_PREFIX IUP_VALUECHANGED_CB;
const char *GO_TABCHANGE_CB = GO_PREFIX IUP_TABCHANGE_CB;
const char *GO_TABCHANGEPOS_CB = GO_PREFIX IUP_TABCHANGEPOS_CB;
const char *GO_SPIN_CB = GO_PREFIX IUP_SPIN_CB;
const char *GO_SHOW_CB = GO_PREFIX IUP_SHOW_CB;
const char *GO_TRAYCLICK_CB = GO_PREFIX IUP_TRAYCLICK_CB;
const char *GO_IDLE_ACTION = GO_PREFIX IUP_IDLE_ACTION;

void goIupSetMapFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_MAP_CB, f);
	IupSetCallback(ih, IUP_MAP_CB, (Icallback) goIupMapCB);
}

void goIupSetUnmapFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_UNMAP_CB, f);
	IupSetCallback(ih, IUP_UNMAP_CB, (Icallback) goIupUnmapCB);
}

void goIupSetDestroyFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_DESTROY_CB, f);
	IupSetCallback(ih, IUP_DESTROY_CB, (Icallback) goIupDestroyCB);
}

void goIupSetGetFocusFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_GETFOCUS_CB, f);
	IupSetCallback(ih, IUP_GETFOCUS_CB, (Icallback) goIupGetFocusCB);
}

void goIupSetKillFocusFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_KILLFOCUS_CB, f);
	IupSetCallback(ih, IUP_KILLFOCUS_CB, (Icallback) goIupKillFocusCB);
}

void goIupSetEnterWindowFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_ENTERWINDOW_CB, f);
	IupSetCallback(ih, IUP_ENTERWINDOW_CB, (Icallback) goIupEnterWindowCB);
}

void goIupSetLeaveWindowFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_LEAVEWINDOW_CB, f);
	IupSetCallback(ih, IUP_LEAVEWINDOW_CB, (Icallback) goIupLeaveWindowCB);
}

void goIupSetKAnyFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_K_ANY_CB, f);
	IupSetCallback(ih, IUP_K_ANY, (Icallback) goIupKAnyCB);
}

void goIupSetHelpFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_HELP_CB, f);
	IupSetCallback(ih, IUP_HELP_CB, (Icallback) goIupHelpCB);
}

void goIupSetButtonFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_BUTTON_CB, f);
	IupSetCallback(ih, IUP_BUTTON_CB, (Icallback) goIupButtonCB);
}

void goIupSetDropFilesFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_DROPFILES_CB, f);
	IupSetCallback(ih, IUP_DROPFILES_CB, (Icallback) goIupDropFilesCB);
}

void goIupSetActionFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupActionCB);
}

void goSetFunc(Ihandle *ih, char *goName, void *gof, char *cName, void *cf) {
	IupSetCallback(ih, goName, gof);
	IupSetCallback(ih, cName, cf);
}

void goIupSetListActionFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_LIST_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupListActionCB);
}

void goIupSetCaretFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_CARET_CB, f);
	IupSetCallback(ih, IUP_CARET_CB, (Icallback) goIupCaretCB);
}

void goIupSetDblclickFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_DBLCLICK_CB, f);
	IupSetCallback(ih, IUP_DBLCLICK_CB, (Icallback) goIupDblclickCB);
}

void goIupSetEditFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_EDIT_CB, f);
	IupSetCallback(ih, IUP_EDIT_CB, (Icallback) goIupEditCB);
}

void goIupSetMotionFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_MOTION_CB, f);
	IupSetCallback(ih, IUP_MOTION_CB, (Icallback) goIupMotionCB);
}

void goIupSetMultiselectFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_MULTISELECT_CB, f);
	IupSetCallback(ih, IUP_MULTISELECT_CB, (Icallback) goIupMultiselectCB);
}

void goIupSetValueChangedFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_VALUECHANGED_CB, f);
	IupSetCallback(ih, IUP_VALUECHANGED_CB, (Icallback) goIupValueChangedCB);
}

void goIupSetTextActionFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupTextActionCB);
}

void goIupSetToggleActionFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_ACTION, f);
	IupSetCallback(ih, IUP_ACTION, (Icallback) goIupToggleActionCB);
}

void goIupSetTabChangeFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_TABCHANGE_CB, f);
	IupSetCallback(ih, IUP_TABCHANGE_CB, (Icallback) goIupTabChangeCB);
}

void goIupSetTabChangePosFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_TABCHANGEPOS_CB, f);
	IupSetCallback(ih, IUP_TABCHANGEPOS_CB, (Icallback) goIupTabChangePosCB);
}

void goIupSetSpinFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_SPIN_CB, f);
	IupSetCallback(ih, IUP_SPIN_CB, (Icallback) goIupSpinCB);
}

void goIupSetShowFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_SHOW_CB, f);
	IupSetCallback(ih, IUP_SHOW_CB, (Icallback) goIupShowCB);
}

void goIupSetTrayClickFunc(Ihandle *ih, int f) {
	IupSetInt(ih, GO_TRAYCLICK_CB, f);
	IupSetCallback(ih, IUP_TRAYCLICK_CB, (Icallback) goIupTrayClickCB);
}

void goIupSetIdleFunc() {
	IupSetFunction(IUP_IDLE_ACTION, (Icallback) goIupIdleCB);
}
