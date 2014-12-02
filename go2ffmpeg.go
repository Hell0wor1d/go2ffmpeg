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
	//	chk := func(err error) {
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//	framePerBuffer := 2048
	args := []string{
		"-i", inputfile,
		"-f", "mp3",
		"-",
	}
	ff := ffmpeg.NewFFmpeg(args)
	if ff == nil {
		log.Fatal("ffmpeg can not be nill.")
	}
	//defer ff.Close()
	//buffer := make([]byte, 1024)
	out, err := os.Create("/Users/kevin/filename.mp3")
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}
	ff.Read(out)
	//	portaudio.Initialize()
	//	defer portaudio.Terminate()
	//	stream, err := portaudio.OpenDefaultStream(0, 2, 44100, framePerBuffer, ff)
	//chk(err)
	//defer stream.Close()
	//chk(stream.Start())
	//if err := ff.Wait(); err != nil {
	//		log.Fatal(err)
	//	}
	//	chk(stream.Stop())
}
