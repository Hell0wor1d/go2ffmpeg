/*
Author Kev7n
Email root@kev7n.com
https://github.com/Hell0wor1d/go2ffmpeg

MIT
*/
package ffmpeg

import (
	//"bytes"
	//	"encoding/binary"
	"io"
	"log"
	"os"
	"os/exec"
)

type FFmpeg struct {
	cmd *exec.Cmd
}

func init() {
	// check if ffmpeg file exists.
	f, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal("ffmpeg dose not install.", err)
	}
	log.Println("ffmpeg path: ", f)
}

func NewFFmpeg(args []string) *FFmpeg {
	cmd := exec.Command("ffmpeg", args...) // "-i", filename, "-f", "s15le", "-")
	//	stdout, err := cmd.StdoutPipe()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	stdin, err := cmd.StdinPipe()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	stderr, err := cmd.StderrPipe()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if err := cmd.Start(); err != nil {
	//		log.Fatal(err)
	//	}
	return &FFmpeg{cmd}
}

func (f *FFmpeg) Write() error {
	return nil
}

func (f *FFmpeg) Read(outWriter io.Writer) error {
	stdout, err := f.cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	stderr, err := f.cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = f.cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer f.wait()
	//go func() {
	//	for n, err = stdout.Read(buffer); err == nil; n, err = stdout.Read(buffer) {
	//		log.Printf("%s\n", buffer[:n])
	//	}
	//}()
	go io.Copy(outWriter, stdout)
	go io.Copy(os.Stdout, stderr)
	return err
}

func (f *FFmpeg) wait() error {
	return f.cmd.Wait()
}

func (f *FFmpeg) Close() error {
	return nil
}

func (e *FFmpeg) ProcessAudio(out [][]int16) {
	//	// int16 takes 2 bytes
	//	bufferSize := len(out[0]) * 4
	//	var pack = make([]byte, bufferSize)
	//	if _, err := e.Out.Read(pack); err != nil {
	//		log.Fatal(err)
	//	}
	//	n := make([]int16, len(out[0])*2)
	//	for i := range n {
	//		var x int16
	//		buf := bytes.NewBuffer(pack[2*i : 2*(i+1)])
	//		binary.Read(buf, binary.LittleEndian, &x)
	//		n[i] = x
	//	}
	//
	//	for i := range out[0] {
	//		out[0][i] = n[2*i]
	//		out[1][i] = n[2*i+1]
	//	}
}
