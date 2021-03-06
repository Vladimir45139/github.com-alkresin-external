// Copyright 2018 Alexander S.Kresin <alex@kresin.ru>, http://www.kresin.ru
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import egui "github.com/alkresin/external"

func main() {

	if egui.Init("") != 0 {
		return
	}

	pWindow := &egui.Widget{X: 100, Y: 100, W: 400, H: 280, Title: "External"}
	egui.InitMainWindow(pWindow)

	egui.Menu("")
	egui.Menu("File")
	egui.AddMenuItem("Open form", 0, openf, "openf")
	egui.AddMenuSeparator()
	egui.AddMenuItem("Exit", 0, nil, "hwg_EndWindow()")
	egui.EndMenu()
	egui.Menu("Help")
	egui.AddMenuItem("About", 0, nil, "hwg_MsgInfo(hb_version()+chr(10)+chr(13)+hwg_version(),\"About\")")
	egui.EndMenu()
	egui.EndMenu()

	pWindow.Activate()

	egui.Exit()
}

func openf([]string) string {

	egui.OpenForm("forms/testget2.xml")
	return ""
}
