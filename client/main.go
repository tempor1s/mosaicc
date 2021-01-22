package main

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	onExit := func() {
		// TODO: cleanup
	}

	systray.Run(onReady, onExit)
}

func onReady() {
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

	mUploadClipboard := systray.AddMenuItemCheckbox("Upload from clipboard", "Upload files that are copied from your clipboard", true)
	go func() {
		for {
			<-mUploadClipboard.ClickedCh
			if mUploadClipboard.Checked() {
				mUploadClipboard.Uncheck()
			} else {
				mUploadClipboard.Check()
			}
		}
	}()

	mUploadFolder := systray.AddMenuItemCheckbox("Upload from folder", "Upload all the files that are created in a folder.", false)
	go func() {
		for {
			<-mUploadFolder.ClickedCh
			if mUploadFolder.Checked() {
				mUploadFolder.Uncheck()
			} else {
				mUploadFolder.Check()
			}
		}
	}()

	// add seperator
	systray.AddSeparator()
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
