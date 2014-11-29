package ffmpeg

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os/exec"
)

type ffmpeg struct {
	in  io.ReadCloser
	Cmd *exec.Cmd
}

func NewFfmpeg(filename string) *ffmpeg {
	cmd := exec.Command("ffmpeg", "-i", filename, "-f", "s16le", "-")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return &ffmpeg{stdout, cmd}
}

func (f *ffmpeg) Close() error {
	return f.in.Close()
}

func (e *ffmpeg) ProcessAudio(_, out [][]int16) {
	// int16 takes 2 bytes
	bufferSize := len(out[0]) * 4
	var pack = make([]byte, bufferSize)
	if _, err := e.in.Read(pack); err != nil {
		log.Fatal(err)
	}
	n := make([]int16, len(out[0])*2)
	for i := range n {
		var x int16
		buf := bytes.NewBuffer(pack[2*i : 2*(i+1)])
		binary.Read(buf, binary.LittleEndian, &x)
		n[i] = x
	}

	for i := range out[0] {
		out[0][i] = n[2*i]
		out[1][i] = n[2*i+1]
	}
}
