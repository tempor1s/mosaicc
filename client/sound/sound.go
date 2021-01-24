package sound

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Player represents a sound player (so we dont have to decode the sound every time etc)
type Player struct {
	Streamer beep.StreamSeekCloser
}

// NewPlayer will return a new sound player struct
func NewPlayer(name string) *Player {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return &Player{
		Streamer: streamer,
	}
}

// Play will play the sound that is loaded in the player struct
func (p *Player) Play() {
	done := make(chan bool)
	speaker.Play(beep.Seq(p.Streamer, beep.Callback(func() {
		log.Println("done playing beep")
		done <- true
	})))

	<-done
}

// Close will cleanly close down the player
func (p *Player) Close() {
	p.Streamer.Close()
}
