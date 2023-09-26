package main

import (
	"PasswordGenerator/funcs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func main() {
	// This code creates a new Fyne application and window. It also sets the window's icon
	myApp := app.New()
	myWindow := myApp.NewWindow("Password Generator")
	r, _ := fyne.LoadResourceFromURLString("https://raw.githubusercontent.com/Hatef-PR/Password-Generator/main/icon.png")
	myWindow.SetIcon(r)

	myWindow.Resize(fyne.NewSize(400, 300))

	// This code creates a drop-down menu with the possible password lengths
	lengthDropdown := widget.NewSelect([]string{"8", "12", "16", "20"}, nil)
	lengthDropdown.PlaceHolder = "Select length"

	// These two labels are used to display the generated password to the user
	passwordLabel := widget.NewLabel("Generated Password:")
	passwordValue := widget.NewLabel("")

	// These three buttons are used to copy the generated password to the clipboard, clear the password, and generate a new password, respectively
	copyButton := widget.NewButton("Copy", func() {
		clipboard.WriteAll(passwordValue.Text)
		dialog.ShowInformation("Password Copied", "Password copied to clipboard", myWindow)
	})

	clearButton := widget.NewButton("Clear", func() {
		lengthDropdown.SetSelected("")
		passwordValue.SetText("")
	})

	generateButton := widget.NewButton("Generate Password", func() {
		selectedLength := lengthDropdown.Selected
		if selectedLength == "" {
			return
		}
		length := funcs.Atoi(selectedLength)
		password := funcs.PassGen(length)
		passwordValue.SetText(password)
	})

	// This code creates a horizontal box containing the copy and clear buttons
	buttons := container.NewHBox(
		copyButton,
		layout.NewSpacer(),
		clearButton,
	)

	// This code creates an about tab with information about the application
	aboutTab := container.NewVBox(
		widget.NewLabel("Password Generator"),
		widget.NewLabel("Version: 1.0.0"),
		widget.NewLabel("Developer: Hatef PourRajabi"),
		widget.NewLabel("Email: hatef.pr@gmail.com"),
	)

	// This code creates an app tabs widget with two tabs: Main and About
	// The Main tab contains the password length drop-down menu, generate button, password labels, and buttons
	// The About tab contains information about the application
	tabs := container.NewAppTabs(
		container.NewTabItem("Main", container.NewVBox(
			widget.NewLabel("Password Length:"),
			lengthDropdown,
			generateButton,
			passwordLabel,
			passwordValue,
			buttons,
		)),
		container.NewTabItem("About", aboutTab),
	)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
