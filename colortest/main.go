package main

import (
    "github.com/huyngo878/img_mod/Colors"
    "github.com/huyngo878/img_mod/GetPic"
    "github.com/huyngo878/img_mod/Grayscale"
    "github.com/huyngo878/img_mod/Text"
)

func main() {
	GetPic.GetPic()  
	Colors.Colors("downloaded_image.jpg") 
	Grayscale.Grayscale("downloaded_image.jpg")
	Text.Text()
}

