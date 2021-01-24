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
	"github.com/tempor1s/mosaic/client/sound"
	"github.com/tempor1s/mosaic/client/upload"
)

// FileWatcher is a structure that will keep track and upload all of the files that are in a folder
type FileWatcher struct {
	mu      sync.Mutex        // mutex lock
	started bool              // if the watcher has been started
	sound   bool              // if the file watcher should play sounds on upload, etc
	folder  string            // the folder to watch
	watcher *fsnotify.Watcher // notify file watcher
	player  *sound.Player     // sound player for upload sounds
}

// New will return a new file watcher
func New(folder string) *FileWatcher {
	// may need to be changed to use absolute directory (or add file to users home dir or something)
	s := sound.NewPlayer("./sound/beep.mp3")

	return &FileWatcher{
		folder: folder,
		player: s, // for upload sound
		sound:  true,
	}
}

// Start will start the file watcher
func (f *FileWatcher) Start() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	// dont start the watcher again if its already started
	if f.started {
		log.Println("could not start watcher. watcher already started")
		return errors.New("could not start watcher. watcher already started")
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("failed to create watcher. err:", err)
		return err
	}

	// assign the watcher so we can manage it
	f.watcher = w

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

					// if we can play sound
					if f.sound {
						f.player.Play()
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
	if err := f.watcher.Add(f.folder); err != nil {
		fmt.Println("error creating watcher for that folder. err:", err)
	}

	// the watcher is now started
	f.started = true

	return nil
}

// Stop will cleanly stop the file watcher
func (f *FileWatcher) Stop() {
	f.mu.Lock()
	defer f.mu.Unlock()

	// stop the watcher
	f.watcher.Close()
	// the watcher is now stoped
	f.started = false
}

// ToggleSound will take the current value of the sound state and set it to the opposite
func (f *FileWatcher) ToggleSound() {
	f.mu.Lock()
	defer f.mu.Unlock()

	// set sound to be the opposite of what it is
	f.sound = !f.sound
	log.Println("set sound value to:", f.sound)
}
