package gui

import (
	"bufio"
	"image/color"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	//"github.com/d-tsuji/clipboard"
)

const (
	width  = 1024
	heigth = 768
)

//slices of buttons
var (
	services  []*widget.Entry
	usernames []*widget.Entry
	passwords []*widget.Entry

	isLocked bool = false
	//TODO add delete button
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
	description := canvas.NewText("Questo è come funziona il di password manager blablabla", color.RGBA{R: uint8(80), G: uint8(0), B: uint8(128), A: uint8(1)})
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

	//legend with service canvas etc
	lock := widget.NewButtonWithIcon("Block entries", fyne.NewStaticResource("lock", getIcon("gui/lock.png")), func() { handleAllLocks() })
	legend := container.NewAdaptiveGrid(5, serviceTitle, usernameTitle, passwordTitle, lock)
	//mainView of the app
	mainView := container.NewVBox(
		title,
		description,
		layout.NewSpacer(),
		legend,
	)
	//masterview of the app (so that it has a layout.NewSpacer after the last element)
	masterView := container.NewVSplit(container.NewVScroll(mainView), layout.NewSpacer())

	addButton := widget.NewButtonWithIcon("", fyne.NewStaticResource("add", getIcon("gui/add.png")), func() {
		services = append(services, widget.NewEntry())
		usernames = append(usernames, widget.NewEntry())
		passwords = append(passwords, widget.NewEntry())
		w := a.NewWindow("Create new entry")
		title := canvas.NewText("Create your new entry here", color.RGBA{R: uint8(80), G: uint8(0), B: uint8(128), A: uint8(1)})
		title.TextSize = 22
		service := widget.NewEntry()
		service.PlaceHolder = "Insert the service here"
		username := widget.NewEntry()
		username.PlaceHolder = "Insert your username here"
		password := widget.NewPasswordEntry()
		password.PlaceHolder = "Insert your password here"
		cancel := widget.NewButton("Cancel", func() {
			w.Hide()
		})
		confirm := widget.NewButton("Confirm", func() {
			services[getFreePosition()-1].SetText(service.Text)
			usernames[getFreePosition()-1].SetText(username.Text)
			passwords[getFreePosition()-1].SetText(password.Text)
			handleLock(services[getFreePosition()-1],
				usernames[getFreePosition()-1], passwords[getFreePosition()-1])

			updateView(mainView)
			w.Hide()
		})
		w.SetContent(container.NewVBox(
			container.NewHBox(layout.NewSpacer(), title, layout.NewSpacer()),
			layout.NewSpacer(),
			service,
			layout.NewSpacer(),
			username,
			layout.NewSpacer(),
			password,
			layout.NewSpacer(),
			container.NewHBox(
				cancel,
				confirm,
			),
		))
		w.Resize(fyne.NewSize(480, 480))
		w.CenterOnScreen()
		w.Show()
	})

	//call the retrieve data function to pick data from the file
	//retrieveData(a, w)
	legend.Add(addButton)
	//updateView(mainView)

	//set content and run
	w.SetContent(
		masterView,
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

// func retrieveData(a fyne.App, w fyne.Window) {
// 	//for i := 0; i < len(services); i++ {
// 	//services[i] = file.Text ...
// 	services = append(services, widget.NewButton("Minecraft", func() {}))
// 	usernames = append(usernames, widget.NewButton("tiziobe435", func() {}))
// 	passwords = append(passwords, widget.NewButton("daudhoasdhd2346°ç*èòpè", func() {}))
// 	edits = append(edits, widget.NewButton("Edit", func() {
// 		w := a.NewWindow("Edit Entry")
// 		title := canvas.NewText("Edit your values here", color.RGBA{R: uint8(80), G: uint8(0), B: uint8(128), A: uint8(1)})
// 		title.TextSize = 22
// 		service := widget.NewEntry()
// 		service.Text = services[getFreePosition()-1].Text
// 		service.PlaceHolder = "Insert the service here"
// 		username := widget.NewEntry()
// 		username.Text = usernames[getFreePosition()-1].Text
// 		username.PlaceHolder = "Insert your username here"
// 		password := widget.NewPasswordEntry()
// 		password.PlaceHolder = "Insert your password here"
// 		service.Text = services[getFreePosition()-1].Text

// 		cancel := widget.NewButton("Cancel", func() {
// 			w.Hide()
// 		})
// 		confirm := widget.NewButton("Confirm", func() {
// 			services[getFreePosition()-1].SetText(service.Text)
// 			usernames[getFreePosition()-1].SetText(username.Text)
// 			passwords[getFreePosition()-1].SetText(password.Text)
// 			w.Hide()
// 		})
// 		w.SetContent(container.NewVBox(
// 			container.NewHBox(layout.NewSpacer(), title, layout.NewSpacer()),
// 			layout.NewSpacer(),
// 			service,
// 			layout.NewSpacer(),
// 			username,
// 			layout.NewSpacer(),
// 			password,
// 			layout.NewSpacer(),
// 			container.NewHBox(
// 				cancel,
// 				confirm,
// 			),
// 		))
// 		w.Resize(fyne.NewSize(480, 480))
// 		w.CenterOnScreen()
// 		w.Show()
// 	}))
// 	copies = append(copies, widget.NewButton("Copy", func() {
// 		clipboard.Set(passwords[getFreePosition()].Text)
// 	}))
// 	//}
// }

func getFreePosition() int {
	return len(services)
}

func updateView(mainView *fyne.Container) {
	mainView.Add(container.NewAdaptiveGrid(5, services[len(services)-1], usernames[len(services)-1], passwords[len(services)-1]))

}

func getIcon(path string) []byte {
	iconFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(iconFile)
	Icon, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return Icon
}

func handleAllLocks() {
	if isLocked {
		for _, x := range services {
			x.Enable()
		}
		for _, x := range usernames {
			x.Enable()
		}
		for _, x := range passwords {
			x.Enable()
		}
		isLocked = false
		return
	}
	for _, x := range services {
		x.Disable()
	}
	for _, x := range usernames {
		x.Disable()

	}
	for _, x := range passwords {
		x.Disable()
	}
	isLocked = true
}

func handleLock(service *widget.Entry, username *widget.Entry, password *widget.Entry) {
	if isLocked {
		service.Disable()
		username.Disable()
		password.Disable()
		return
	}
	service.Enable()
	username.Enable()
	password.Enable()
}
