package converter

import (
	"os"
	"os/exec"
)

func CreateMp4(basePath, image, sound string) (string, error) {
	//ffmpeg -loop 1 -y -i test.png -i y2mate.com\ -\ saufen_junge_giJ7O2GIopA.mp3 -shortest -strict -2 out.mp4
	cmd := exec.Command(
		"ffmpeg",
		"-loop", "1",
		"-y",
		"-i", basePath+image,
		"-i", basePath+sound,
		"-shortest",
		"-strict", " -2",
		basePath+"out.mp4")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return basePath + "out.mp4", cmd.Run()
}
