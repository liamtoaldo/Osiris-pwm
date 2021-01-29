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
	"fyne.io/fyne/v2/widget"
	"github.com/d-tsuji/clipboard"
)

const (
	width  = 1024
	heigth = 768
)

//slices of buttons
var (
	services  [99]*widget.Button
	usernames [99]*widget.Button
	passwords [99]*widget.Button
	edits     [99]*widget.Button
	copies    [99]*widget.Button
)

//SetTextSize modifies the text size with a variadic function to make it easier
func SetTextSize(size float32, text ...*canvas.Text) {
	for i := 0; i < len(text); i++ {
		text[i].TextSize = size
	}
}

//Call calls the main gui application
func Call() {
	a := app.New()
	w := a.NewWindow("Osiris Password Manager")
	w.SetTitle("Osiris Password Manager")

	//usable variables
	spacer := layout.NewSpacer()
	purple := color.RGBA{R: uint8(128), G: uint8(0), B: uint8(128), A: uint8(1)}

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

	//title and description
	title := canvas.NewText("Osiris Password Manager", purple)
	title.TextSize = 34
	title.TextStyle.Bold = true
	title.Alignment = fyne.TextAlignCenter
	description := canvas.NewText("Questo è come funziona il porcodio di password manager blablabla", color.RGBA{R: uint8(80), G: uint8(0), B: uint8(128), A: uint8(1)})
	description.TextSize = 22
	description.Alignment = fyne.TextAlignCenter

	//table with usernames, values and passwords
	serviceTitle := canvas.NewText("Services", purple)
	serviceTitle.Alignment = fyne.TextAlignCenter
	usernameTitle := canvas.NewText("Usernames", purple)
	usernameTitle.Alignment = fyne.TextAlignCenter
	passwordTitle := canvas.NewText("Passwords", purple)
	passwordTitle.Alignment = fyne.TextAlignCenter
	SetTextSize(20, serviceTitle, usernameTitle, passwordTitle)

	//call the retrieve data function to pick data from the file
	retrieveData()

	//set content and run
	w.SetContent(
		container.NewVBox(
			title,
			description,
			spacer,
			container.NewAdaptiveGrid(5, serviceTitle, usernameTitle, passwordTitle),
			container.NewAdaptiveGrid(5, services[0], usernames[0], passwords[0], edits[0], copies[0]),
			spacer,
		),
	)
	w.Resize(fyne.NewSize(width, heigth))
	w.CenterOnScreen()
	w.ShowAndRun()

}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

func retrieveData() {
	//for i := 0; i < len(services); i++ {
	//services[i] = file.Text ...
	services[0] = widget.NewButton("Minecraft", func() {})
	usernames[0] = widget.NewButton("tiziobe435", func() {})
	passwords[0] = widget.NewButton("daudhoasdhd2346°ç*èòpè", func() {})
	edits[0] = widget.NewButton("Edit", func() {

	})
	copies[0] = widget.NewButton("Copy", func() {
		clipboard.Set(passwords[0].Text)
	})
	//}
}
