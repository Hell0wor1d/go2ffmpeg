package main

import (
	//"code.google.com/p/portaudio-go/portaudio"
	"fmt"
	"github.com/Hell0wor1d/go2ffmpeg/ffmpeg"
	"log"
	"os"
)

func init() {
	fmt.Printf("init app.")
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage:\n \t./ffmpeg audio_file")
		os.Exit(0)
	}
	inputfile := os.Args[1]

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	chk := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	finfo, err := os.Stat(inputfile)
	chk(err)
	if !finfo.Mode().IsRegular() {
		panic("target is not a file.")
	}
	args := []string{
		"-i", inputfile,
		"-f", "mp3",
		"-",
	}
	ff := ffmpeg.NewFFmpeg(args)
	if ff == nil {
		panic("ffmpeg can not be nil.")
	}
	out, err := os.Create("/Users/kevin/filename.mp3")
	defer out.Close()
	chk(err)
	ff.Read(out)
}
