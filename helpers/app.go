package helpers

import (
	"net/url"
	"sort"

	"fyne.io/fyne/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func App() {
	// Create a new window
	myApp := app.New()
	appWindow := myApp.NewWindow("A.C.I.T Traffic Listener")
	icon, _ := fyne.LoadResourceFromPath("icons/sniff.jpg")
	appWindow.SetIcon(icon)
	appWindow.Resize(fyne.NewSize(500, 500))

	// Some text for the window

	label := widget.NewLabel("EXPERIMENTAL Server Change Program for Div2")
	linkURL, _ := url.Parse("https://www.winpcap.org/install/default.htm")
	labelText := "Download WinPCap"
	hyperlink := widget.NewHyperlink(labelText, linkURL)

	labelInfo := widget.NewLabel("Reboot your machine after this.")
	author := widget.NewLabel("Author: alexanderdth")
	
	// Create a map of options
	options := GrabAllDevices()

	// Convert the map of options to a sorted slice
	var optionKeys []string
	for key := range options {
		optionKeys = append(optionKeys, key)
	}
	sort.Strings(optionKeys)

	// Create a slice of option values
	var optionValues []string
	for _, key := range optionKeys {
		optionValues = append(optionValues, options[key])
	}
	deviceInput := widget.NewSelect(optionValues, nil)

	form := widget.NewForm(
		widget.NewFormItem("Network Adapter: ", deviceInput),
	)

	appWindow.SetContent(container.NewVBox(
		label,
		hyperlink,
		labelInfo,
		form,
		author,
	))
	appWindow.Resize(fyne.NewSize(550, 285))
	appWindow.SetFixedSize(true)
	appWindow.CenterOnScreen()
	appWindow.ShowAndRun()
}
