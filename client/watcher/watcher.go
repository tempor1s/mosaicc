package watcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"github.com/atotto/clipboard"
	"github.com/fsnotify/fsnotify"
	"github.com/tempor1s/mosaic/client/upload"
)

// FileWatcher is a structure that will keep track and upload all of the files that are in a folder
type FileWatcher struct {
	mu      sync.Mutex        // mutex lock
	Started bool              // if the watcher has been started
	Folder  string            // the folder to watch
	Watcher *fsnotify.Watcher // notify file watcher
}

// New will return a new file watcher
func New(folder string) *FileWatcher {
	return &FileWatcher{
		Folder: folder,
	}
}

// Start will start the file watcher
func (f *FileWatcher) Start() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// dont start the watcher again if its already started
	if f.Started {
		log.Println("could not start watcher. watcher already started")
		return errors.New("could not start watcher. watcher already started")
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("failed to create watcher. err:", err)
		return err
	}

	// assign the watcher so we can manage it
	f.Watcher = w

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				// only do stuff when a new file is created
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file:", event.Name)

					// get the name of the file
					fileName := filepath.Base(event.Name)
					// get the contents of the file
					bytes, err := ioutil.ReadFile(event.Name)
					if err != nil {
						log.Println("failed to read file. err:", err)
						return
					}

					// TODO: upload to both mosaic and imgur
					// upload to imgur and copy url
					url, err := upload.ToImgur(fileName, bytes)
					if err != nil {
						log.Println("failed to upload to imgur. err:", err)
						return
					}

					// write the URL to clipboard
					err = clipboard.WriteAll(url)
					if err != nil {
						log.Println("failed to copy url to clipboard. err:", err)
						return
					}
				}
			case err, ok := <-w.Errors:
				if !ok {
					return
				}

				fmt.Println("watcher err event. err:", err)
			}
		}
	}()

	// start the watcher on the given folder
	if err := f.Watcher.Add(f.Folder); err != nil {
		fmt.Println("error creating watcher for that folder. err:", err)
	}

	f.Started = true

	<-done

	return nil
}

// Stop will cleanly stop the file watcher
func (f *FileWatcher) Stop() {
	f.mu.Lock()
	defer f.mu.Unlock()

	// stop the watcher
	f.Watcher.Close()
	f.Started = false
}
