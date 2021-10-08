package main

import (
	"fmt"
	"image/color"
	"image/png"
	"net/http"
)
import "y-captcha/core"

//var cap *core.Captcha

func main()  {
	cap := core.New()

    //We can load font not only from localfile, but also from any []byte slice
	//fontContenrs, err := ioutil.ReadFile("comic.ttf")
	//if err != nil {
	//	panic(err.Error())
	//}
	//err = cap.AddFontFromBytes(fontContenrs)
	//if err != nil {
	//	panic(err.Error())
	//}

	cap.SetSize(128, 64)
	cap.SetDisturbance(core.HIGH)
	//cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cap.Create(4, core.ALL)
		png.Encode(w, img)
		println(str)
	})

	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		base64Img, str := cap.CreateBase64(4, core.ALL)
		println(base64Img)
		println(str)
		htmlImage := "<img src=\"data:image/png;base64," + base64Img + "\" />"
		fmt.Println(htmlImage)
	})


	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.RawQuery
		img := cap.CreateCustom(str)
		png.Encode(w, img)
	})

	http.ListenAndServe(":8085", nil)

}