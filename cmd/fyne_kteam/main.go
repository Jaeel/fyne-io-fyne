package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"net/http"
)

var topWindow fyne.Window
func main() {
	a := app.NewWithID("io.fyne.demo")
	w := a.NewWindow("Fyne Demo")
	topWindow = w
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(640, 460))
	go httpServer()
	w.ShowAndRun()
}

func httpServer(){
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		topWindow.Show()
		topWindow.Maximize()
		topWindow.CenterOnScreen()
	})

	log.Println("Starting v2 httpserver")
	log.Println(http.ListenAndServe(":9999", mux))
}
