package converter

import (
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func CreateMp4(basePath, image, sound string) (string, error) {
	var cmd *exec.Cmd
	mimeType, err := getMimeType(basePath + image)
	if err != nil {
		return "", err
	}
	if strings.Contains(mimeType, "gif") {
		cmd = getGifCmd(basePath, image, sound)
	} else {
		cmd = getImgCmd(basePath, image, sound)
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return basePath + "out.mp4", cmd.Run()
}

func getImgCmd(basePath, image, sound string) *exec.Cmd {
	//ffmpeg -loop 1 -y -i test.png -i y2mate.com\ -\ saufen_junge_giJ7O2GIopA.mp3 -shortest -strict -2 out.mp4
	return exec.Command(
		"ffmpeg",
		"-loop", "1",
		"-y",
		"-i", basePath+image,
		"-i", basePath+sound,
		"-shortest",
		"-strict", " -2",
		basePath+"out.mp4")
}

func getGifCmd(basePath, image, sound string) *exec.Cmd {
	// ffmpeg -i bla.mp3 -i tenor.gif -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -strict -2 -c:v libx264 -threads 4 -c:a aac -b:a 192k -pix_fmt yuv420p out.mp4
	return exec.Command(
		"ffmpeg",
		"-i", basePath+sound,
		"-i", basePath+image,
		"-vf", "scale=trunc(iw/2)*2:trunc(ih/2)*2",
		"-strict", " -2",
		"-c:v", "libx264",
		"-threads", "4",
		"-c:a", "aac",
		"-b:a", "192k",
		"-pix_fmt", "yuv420p",
		"-y",
		basePath+"out.mp4")
}

func getMimeType(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	return http.DetectContentType(buffer), nil

}
