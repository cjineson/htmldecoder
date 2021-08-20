package main

import (
	"html"

	"github.com/tadvi/winc"
)

func main() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(1000, 1000)
	mainWindow.SetText("HTML Decoder")

	in := winc.NewEdit(mainWindow)
	in.SetPos(20, 20)
	in.SetSize(460, 870)
	in.SetText("Paste HTML to decode here")

	out := winc.NewEdit(mainWindow)
	out.SetPos(500, 20)
	out.SetSize(460, 870)

	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Decode")
	btn.SetPos(20, 910)
	btn.SetSize(940, 40)
	btn.OnClick().Bind(func(e *winc.Event) {
		out.SetText(html.UnescapeString(in.Text()))
	})

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.EnableSizable(false)
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
