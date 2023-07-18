package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"os"
)

type soundDemo struct {
	audioContext *audio.Context
	soundPlayer  *audio.Player
	counter      int
}

func (demo *soundDemo) Update() error {
	demo.counter += 1
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		demo.soundPlayer.Rewind()
		demo.soundPlayer.Play()
		demo.counter = 0
	}
	return nil
}

func (s soundDemo) Draw(screen *ebiten.Image) {
	if s.counter >= 20 {
		screen.Fill(colornames.Crimson)
	} else {
		screen.Fill(colornames.Deepskyblue)
	}
}

func (s soundDemo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//TODO implement me
	return outsideWidth, outsideHeight
}

const (
	WINDOW_WIDTH      = 512
	WINDOW_HEIGHT     = 512
	SOUND_SAMPLE_RATE = 48000
)

func main() {
	soundContext := audio.NewContext(SOUND_SAMPLE_RATE)
	soundGame := soundDemo{
		audioContext: soundContext,
		soundPlayer:  LoadWav("Thunder1.wav", soundContext),
		counter:      20,
	}
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Demo Space Scroller")
	err := ebiten.RunGame(&soundGame)
	if err != nil {

	}
}

func LoadWav(name string, context *audio.Context) *audio.Player {
	thunderFile, err := os.Open(name)
	if err != nil {
		fmt.Println("Error Loading sound: ", err)
	}
	thunderSound, err := wav.DecodeWithoutResampling(thunderFile)
	if err != nil {
		fmt.Println("Error interpreting sound file: ", err)
	}
	soundPlayer, err := context.NewPlayer(thunderSound)
	if err != nil {
		fmt.Println("Couldn't create sound player: ", err)
	}
	return soundPlayer
}
