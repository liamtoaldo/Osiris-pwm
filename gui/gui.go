package gui

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

//Call calls the main gui application
func Call() {
	a := app.New()
	w := a.NewWindow("Osiris Password Manager")
	w.SetTitle("Osiris Password Manager")

	//usable variables
	spacer := layout.NewSpacer()
	spacer.Move(fyne.NewPos(1, 1))

	//menu with settings, edit and help
	settingsItem := fyne.NewMenuItem("Settings", func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(&fyne.ShortcutCut{
			Clipboard: w.Clipboard(),
		}, w)
	})
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(&fyne.ShortcutCopy{
			Clipboard: w.Clipboard(),
		}, w)
	})
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(&fyne.ShortcutPaste{
			Clipboard: w.Clipboard(),
		}, w)
	})
	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://github.com/Gyro7/Osiris-pwm")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Request feature", func() {
			u, _ := url.Parse("https://github.com/Gyro7/Osiris-pwm/pulls")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Report a bug", func() {
			u, _ := url.Parse("https://github.com/Gyro7/Osiris-pwm/issues/new")
			_ = a.OpenURL(u)
		}))
	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to the first menu
		fyne.NewMenu("File", fyne.NewMenuItemSeparator(), settingsItem),
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator()),
		helpMenu,
	)
	w.SetMainMenu(mainMenu)
	w.SetMaster()

	//title := widget.NewLabelWithStyle("Hello Fyne!", fyne.TextAlignCenter, fyne.TextStyle{})
	title := canvas.NewText("Osiris Password Manager", color.RGBA{R: uint8(128), G: uint8(0), B: uint8(128), A: uint8(1)})
	title.TextSize = 34
	title.TextStyle.Bold = true
	title.Alignment = fyne.TextAlignCenter
	description := canvas.NewText("Questo è come funziona il porcodio di password manager blablabla", color.RGBA{R: uint8(80), G: uint8(0), B: uint8(128), A: uint8(1)})
	description.TextSize = 20
	description.Alignment = fyne.TextAlignCenter

	//set content and run
	w.SetContent(container.NewVBox(title, spacer, description, spacer))
	w.ShowAndRun()
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
