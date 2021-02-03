package gui

import (
	"Osiris-pwm/crypt"
	"bufio"

	//"image/color"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	//"github.com/atotto/clipboard"
	"github.com/d-tsuji/clipboard"
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
	grids     []*fyne.Container

	isLocked bool = false
)

//Call calls the main gui application
func Call() {
	a := app.New()
	w := a.NewWindow("Osiris Password Manager")
	w.SetTitle("Osiris Password Manager")

	//usable variables
	//mainColor := color.RGBA{R: uint8(128), G: uint8(0), B: uint8(128), A: uint8(1)}
	mainColor := theme.PrimaryColor()
	mainView := container.NewVBox()

	//menu with settings, edit and help
	settingsItem := fyne.NewMenuItem("Settings", func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
		dialog.ShowInformation("Message to the end user", "The colors only apply at the restart of the application", w)
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
	deleteItem := fyne.NewMenuItem("Delete empties", func() {
		for i := 0; i < len(grids); i++ {
			if services[i].Text == "" && usernames[i].Text == "" && passwords[i].Text == "" {
				updateView(mainView, false, i)
			}
		}
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
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), deleteItem),
		helpMenu,
	)
	w.SetMainMenu(mainMenu)
	w.SetMaster()

	//title and description
	title := canvas.NewText("Osiris Password Manager", theme.PrimaryColor())
	title.TextSize = 34
	title.TextStyle.Bold = true
	title.Alignment = fyne.TextAlignCenter
	description := canvas.NewText("Questo è come funziona il di password manager blablabla", theme.PrimaryColor())
	description.TextSize = 22
	description.Alignment = fyne.TextAlignCenter
	description.TextStyle = fyne.TextStyle{Italic: true}

	//table with usernames, values and passwords
	serviceTitle := canvas.NewText("Services", mainColor)
	serviceTitle.Alignment = fyne.TextAlignCenter
	usernameTitle := canvas.NewText("Usernames", mainColor)
	usernameTitle.Alignment = fyne.TextAlignCenter
	passwordTitle := canvas.NewText("Passwords", mainColor)
	passwordTitle.Alignment = fyne.TextAlignCenter
	setTextSize(20, serviceTitle, usernameTitle, passwordTitle)

	//legend with service canvas etc
	lock := widget.NewButtonWithIcon("Block entries", fyne.NewStaticResource("lock", getIcon("gui/lock.png")), func() { handleAllLocks() })
	legend := container.NewAdaptiveGrid(5, serviceTitle, usernameTitle, passwordTitle, lock)
	//mainView of the app
	mainView = container.NewVBox(
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
		title := canvas.NewText("Create your new entry here", mainColor)
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
			handleLock(services[getFreePosition()-1], usernames[getFreePosition()-1], passwords[getFreePosition()-1])
			w.Hide()
			updateView(mainView, true, len(services)-1)
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

	loginHandler(a, w)
	keyHandler(a, w)
	toSHandler(a, w)
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

//SetTextSize modifies the text size with a variadic function to make it easier
func setTextSize(size float32, text ...*canvas.Text) {
	for i := 0; i < len(text); i++ {
		text[i].TextSize = size
	}
}

//TODO
// func retrieveData(a fyne.App, w fyne.Window) {
// 	//for i := 0; i < len(services); i++ {
// 	//services[i] = file.Text ...

func getFreePosition() int {
	return len(services)
}

func updateView(mainView *fyne.Container, add bool, i int) {
	if add {
		grids = append(grids, container.NewAdaptiveGrid(5, services[i], usernames[i], passwords[i]))

		//FIXME Sometimes, you need to move the cursor over the black split line in order for the grid to get added to the view
		go mainView.Add(grids[i])
	} else {
		mainView.Remove(grids[i])
	}

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

func enableDisableAll(items []*widget.Entry, enable bool) {
	if enable {
		for _, x := range items {
			x.Enable()
		}
	} else {
		for _, x := range items {
			x.Disable()
		}
	}
}
func handleAllLocks() {
	if isLocked {
		enableDisableAll(services, true)
		enableDisableAll(usernames, true)
		enableDisableAll(passwords, true)
		isLocked = false
		return
	}
	enableDisableAll(services, false)
	enableDisableAll(usernames, false)
	enableDisableAll(passwords, false)
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

func remove(slice []*fyne.Container, s int) []*fyne.Container {
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}
	return slice
}

func toSHandler(a fyne.App, w fyne.Window) {
	//show dialog which says to accept terms of service
	data, err := ioutil.ReadFile("data/termsAccepted.txt")
	if err != nil {
		os.Create("data/termsAccepted.txt")
		data, err = ioutil.ReadFile("data/termsAccepted.txt")
	}

	if string(data) == "" {
		agreement := dialog.NewConfirm("Terms of Service", "Osiris password manager is provided as is and without guarantees of any kind,\n by accepting you confirm that I, Gyro7, take no responsibilities on your data\n and I am not going to help you in case of lost/stolen data\n", func(accepted bool) {
			f := "data/termsAccepted.txt"
			if accepted == true {
				crypt.EncryptStringInFile(crypt.GetToSKey(), "true", f)
			} else {
				os.Remove(f)
				a.Quit()
				if err != nil {
					println(err)
				}
			}
		}, w)
		agreement.SetConfirmText("I agree")
		agreement.SetDismissText("No, close the app")
		agreement.Show()
	}
}

func keyHandler(a fyne.App, w fyne.Window) {
	//show dialog which shows the key (only for the first time)
	data, err := ioutil.ReadFile("data/keyShown.txt")
	if err != nil {
		os.Create("data/keyShown.txt")
		data, err = ioutil.ReadFile("data/keyShown.txt")
	}

	keyLabel := canvas.NewText("", theme.PrimaryColor())
	keyLabel.Text = crypt.GenerateKey()
	keyLabel.TextStyle.Monospace = true
	keyLabel.Alignment = fyne.TextAlignCenter

	//code to handle the autoCompletion file
	autoCompletion := widget.NewRadioGroup([]string{"Enable Auto-Completion", "Don't enable it"}, func(choice string) {
		f := "data/autoCompletion.txt"
		if choice == "Enable Auto-Completion" {
			//at every change recreate and rewrite the file to avoid the file being full of unused data
			os.Create(f)
			crypt.EncryptStringInFile(crypt.GetToSKey(), "true", f)
		} else {
			os.Remove(f)
		}
	})

	//code to handle the keyShown file
	if string(data) == "" {
		showkey := dialog.NewCustomConfirm(" Here is your key, copy it and save it somewhere, because you'll be asked it \n every time you open Osiris if you don't enable auto-completion.", "Confirm", "Cancel", container.NewVBox(keyLabel, widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
			clipboard.Set((keyLabel.Text))
		}), autoCompletion), func(accepted bool) {
			f := "data/keyShown.txt"
			if accepted {
				crypt.EncryptStringInFile(crypt.GetToSKey(), "true", f)
				os.Create("data/masterKey.txt")
				crypt.EncryptStringInFile(crypt.GetGlobalKey(), keyLabel.Text, "data/masterKey.txt")
			} else {
				os.Remove(f)
				a.Quit()
				if err != nil {
					println(err)
				}
			}
		}, w)
		showkey.Show()
	}
}

func loginHandler(a fyne.App, w fyne.Window) {
	//checks if it is the first time the application is opened and returns, because I don't want the user to write the key the first time he enters in the application
	tmp, err := ioutil.ReadFile("data/keyShown.txt")
	if err != nil || string(tmp) == "" {
		return
	}

	_, err = ioutil.ReadFile("data/masterKey.txt")
	if err != nil {
		os.Create("data/masterKey.txt")
		crypt.EncryptStringInFile(crypt.GetGlobalKey(), "nonzo", "data/masterKey.txt")
	}
	KEY := crypt.DecryptStringFromFile(crypt.GetGlobalKey(), "data/masterKey.txt")
	//show dialog which says to accept terms of service
	data, err := ioutil.ReadFile("data/autoCompletion.txt")
	if err != nil {
		os.Create("data/autoCompletion.txt")
		data, err = ioutil.ReadFile("data/autoCompletion.txt")
	}
	//the entry for the key in the dialog interface
	keyEntry := widget.NewPasswordEntry()

	//what happens if the auto completion is not enabled
	if string(data) == "" {
		agreement := dialog.NewCustomConfirm("  Authentication - Insert your given key here  ", "Confirm", "Cancel", keyEntry, func(confirmed bool) {
			if !confirmed {
				a.Quit()
			}
		}, w)

		agreement.Show()
		agreement.SetOnClosed(func() {
			if keyEntry.Text != KEY {
				agreement.Show()
			}
		})
	}

	//what happens if the auto completion is enabled
	if string(data) != "" {
		//autocomplete the entry with the key
		keyEntry.Text = KEY
		agreement := dialog.NewCustomConfirm("  Authentication - Insert your given key here  ", "Confirm", "Cancel", keyEntry, func(confirmed bool) {
			if !confirmed {
				a.Quit()
			}
		}, w)

		agreement.Show()
		agreement.SetOnClosed(func() {
			if keyEntry.Text != KEY {
				agreement.Show()
			}
		})
	}
}
