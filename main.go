package main

import (
	"fmt"
	"next/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main(){
	a:= app.New()
	w:= a.NewWindow("Weather App Deb")

	w.Resize(fyne.NewSize(500,500))

	//API
	res, err:= http.Get("http://api.openweathermap.org/data/2.5/weather?q=kolkata&appid=f30e52a9877e7ebb68688c3e5ca68782")
	if err!=nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	

	w.ShowAndRun()

}