package play

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	vlc "github.com/adrg/libvlc-go"
)

const basePath = "files/"
const vidName = "out.mp4"

func init() {
	// Initialize libvlc. Additional command line arguments can be passed in
	// to libvlc by specifying them in the Init function.
	if err := vlc.Init("--quiet"); err != nil {
		log.Fatal("Could not initialize VLC!")
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		log.Println("Releasing VLC")
		vlc.Release()
	}()
}

func PlayMeme(name string) error {

	// Create a new player
	player, err := vlc.NewPlayer()
	if err != nil {
		return err
	}

	defer func() {
		player.Stop()
		player.Release()
	}()

	// Add a media file from path or from URL.
	// Set player media from path:
	// media, err := player.LoadMediaFromPath("localpath/test.mp4")
	// Set player media from URL:
	media, err := player.LoadMediaFromPath(basePath + name + "/" + vidName)
	if err != nil {
		return err
	}
	defer media.Release()

	// Play
	err = player.Play()
	if err != nil {
		return err
	}
	err = player.SetFullScreen(true)
	if err != nil {
		return err
	}
	// Wait some amount of time for the media to start playing.
	// Depends on the version of libvlc. From my tests, libvlc 3.X does not
	// need this delay.
	// TODO: Implement proper callbacks for getting the state of the media.
	 time.Sleep(1 * time.Second)

	// If the media played is a live stream the length will be 0
	length, err := player.MediaLength()
	if err != nil || length == 0 {
		length = 1000 * 60
	}

	time.Sleep(time.Duration(length) * time.Millisecond)
	return nil
}
