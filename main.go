package main

import (
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//アプリを新規作成
	a := app.New()
	//ダークモードに設定
	a.Settings().SetTheme(theme.DarkTheme())
	//ウィンドウを新規作成
	w := a.NewWindow("Calc")
	//テキスト入力用エントリーを作成
	edit := widget.NewEntry()
	//エントリーを複数入力可能に設定
	edit.MultiLine = true
	//スクロールコンテナ作成
	sc := container.NewScroll(edit)
	//ラベルを設定
	inf := widget.NewLabel("information bar.")

//new file function.
nf := func() {
	dialog.ShowConfirm("Alert", "Create New document?", func(f bool) {
		if f {
			edit.SetText("")
			inf.SetText("create new document.")
		}
	}, w)
}

//open file function
of := func() {
	f := widget.NewEntry()
	dialog.ShowCustomConfirm("open file name.", "OK", "Cancel", f, func(b bool){
		if b {
			fn := f.Text + ".txt"
			ba, er := ioutil.ReadFile(fn)
			if er != nil {
				dialog.ShowError(er, w)
			} else {
				edit.SetText(string(ba))
				inf.SetText("Open from file '" + fn + "'.")
			}
		}
	}, w)
}
//save file function
sf := func() {
	f := widget.NewEntry()
	dialog.ShowCustomConfirm("Save file name.", "OK", "Cancel", f, func(b bool) {
		if b{
			fn := f.Text + ".txt"
			er := ioutil.WriteFile(fn,[]byte(edit.Text),os.ModePerm)
			if er != nil {
				dialog.ShowError(er, w)
				return
			}
			inf.SetText("Save to file '" + fn + "'.")
		}
	}, w)
}

//quit function
qf := func() {
	dialog.ShowConfirm("Alert", "Quit application?", func(f bool){
		if f{
			a.Quit()
		}
	}, w)
}
tf := true

//change theme function
cf := func() {
	if tf {
		a.Settings().SetTheme(theme.LightTheme())
		inf.SetText("change to Light-Theme.")
	} else {
		a.Settings().SetTheme(theme.DarkTheme())
		inf.SetText("change to Dark-Theme.")
	}
	tf = !tf
}

//create manubar function
createMenubar := func() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() {
		nf()
	}),
fyne.NewMenuItem("Open...", func() {
	of()
}),
fyne.NewMenuItem("Save...", func() {
	sf()
}),
fyne.NewMenuItem("Change Theme", func() {
	cf()
}),
fyne.NewMenuItem("Quit", func() {
	qf()
}),
),
fyne.NewMenu("Edit",fyne.NewMenuItem("Cut", func() {
	edit.TypedShortcut(&fyne.ShortcutCut{
		Clipboard: w.Clipboard()})
	inf.SetText("Cut text.")
}),
fyne.NewMenuItem("Copy", func() {
	edit.TypedShortcut(
		&fyne.ShortcutCopy{
			Clipboard: w.Clipboard()})
			inf.SetText("Copy text.")
}),
fyne.NewMenuItem("paste", func() {
	edit.TypedShortcut(
		&fyne.ShortcutPaste{
			Clipboard: w.Clipboard()})
			inf.SetText("Paste text.")
			}),
		),
	)
}

//create toolbar function
createToolbar := func() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(
			theme.DocumentCreateIcon(), func() {
				nf()
			}),
		widget.NewToolbarAction(
			theme.FolderOpenIcon(), func() {
				of()
			}),
		widget.NewToolbarAction(
			theme.DocumentSaveIcon(), func() {
				sf()
			}),
		)
	}

md := createMenubar()
tb := createToolbar()

w.SetMainMenu(md)
w.SetContent(
	fyne.NewContainerWithLayout(
		layout.NewBorderLayout(
			tb, inf, nil, nil,
		),
		tb, inf, sc,
	),
)
w.Resize(fyne.NewSize(500, 500))
w.ShowAndRun()
}