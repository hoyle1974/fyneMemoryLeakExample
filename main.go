package main

import (
	_ "net/http/pprof"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	leakyApp := app.New()
	window := leakyApp.NewWindow("Memory Leak")

	container := container.NewBorder(nil, nil, nil, nil, nil)
	window.Resize(fyne.NewSize(512, 512))

	window.SetContent(container)

	// Simply update an image with the same image
	go func() {
		var oldImage *canvas.Image

		for {
			for t := 0; t < 10000; t++ {
				newImage := canvas.NewImageFromResource(theme.FyneLogo())
				newImage.FillMode = canvas.ImageFillContain

				container.Remove(oldImage)
				container.Add(newImage)

				oldImage = newImage

				// If I remove this line, the memory leak stops.  I'm pretty sure it has something to do
				// with the driver cache
				container.Refresh()

				// If I remove this the application seems to crash after running awhile
				time.Sleep(time.Millisecond)
			}
		}
	}()

	window.ShowAndRun()
}
