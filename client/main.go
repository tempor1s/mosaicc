package main

import (
	"fmt"
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/tempor1s/mosaic/client/watcher"
)

// State represents some of the internal app state that that the local client needs
type State struct {
	Watcher *watcher.FileWatcher // the folder watcher for auto-uploads
}

func main() {
	w := watcher.New("/Users/ben/dev/go/mosaic/client/test")

	d := State{
		Watcher: w,
	}

	onExit := func() {
		// cleanup the file watcher
		d.Watcher.Stop()
	}

	systray.Run(d.onReady, onExit)
}

func (s *State) onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTooltip("Manage the Mosaic screenshot tool.")

	mUploadMosaic := systray.AddMenuItemCheckbox("Upload to mosaic", "Upload the image to mosaic", false)
	mUploadImgur := systray.AddMenuItemCheckbox("Upload to imgur", "Upload the image to imgur", true)

	// functionality for imgur menu button
	go func() {
		for {
			<-mUploadImgur.ClickedCh
			if mUploadImgur.Checked() {
				mUploadImgur.Uncheck()
			} else {
				mUploadImgur.Check()

				// uncheck mosaic since we are now uploading to imgur (dont want to upload to both)
				if mUploadMosaic.Checked() {
					mUploadMosaic.Uncheck()
				}
			}
		}
	}()

	// functionality for mosaic menu button
	go func() {
		for {
			<-mUploadMosaic.ClickedCh
			if mUploadMosaic.Checked() {
				mUploadMosaic.Uncheck()
			} else {
				mUploadMosaic.Check()

				// uncheck mosaic since we are now uploading to imgur (dont want to upload to both)
				if mUploadImgur.Checked() {
					mUploadImgur.Uncheck()
				}
			}
		}
	}()

	systray.AddSeparator()

	mUploadFolder := systray.AddMenuItemCheckbox("Upload from folder", "Upload all the files that are created in a folder.", false)
	go func() {
		for {
			<-mUploadFolder.ClickedCh
			if mUploadFolder.Checked() {
				log.Println("Stopping watcher...")

				mUploadFolder.Uncheck()
				go s.Watcher.Stop()
			} else {
				log.Println("Starting watcher...")

				mUploadFolder.Check()
				go s.Watcher.Start()
			}
		}
	}()

	// add seperator
	systray.AddSeparator()

	// all the user to toggle image upload sound
	mUploadSound := systray.AddMenuItemCheckbox("Sound", "Play a sound when the image is uploaded", true)
	go func() {
		for {
			<-mUploadSound.ClickedCh

			if mUploadSound.Checked() {
				log.Println("Turning off notification sound.")

				mUploadSound.Uncheck()
				go s.Watcher.ToggleSound()
			} else {
				log.Println("Turning on notification sound.")

				mUploadSound.Check()
				go s.Watcher.ToggleSound()
			}

		}
	}()

	// menu option to quit the application
	mQuit := systray.AddMenuItem("Quit Mosaic", "Quit Mosaic")
	// listen for close event (clicked on "Quit")
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("requesting quit")
		systray.Quit()
		fmt.Println("quit")
	}()

}
