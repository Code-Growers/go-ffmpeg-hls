package main

import (
	"github.com/Code-Growers/hls_test/hls"
	"log"
)

func main() {
	ffmpegPath := "/usr/bin/ffmpeg"
	srcPath := "./src.mp4"
	targetPath := "./hls_output"
	resOptions := []string{"480p", "720p", "1080p"}
	variants, _ := hls.GenerateHLSVariant(resOptions, "")
	hls.GeneratePlaylist(variants, targetPath, "")

	for _, res := range resOptions {
		if err := hls.GenerateHLS(ffmpegPath, srcPath, targetPath, res); err != nil {
			log.Println(err)
		}
	}
}
