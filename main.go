package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gordonklaus/portaudio"
)

func main() {
	log.Println("starting")

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 512)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	if err != nil {
		panic(err)
	}

	err = stream.Start()
	if err != nil {
		panic(err)
	}

	ch := time.After(time.Second * 10)

	for {
		stream.Read()
		fmt.Print(rms(in), " ")
		select {
		case <-ch:
			break
		default:
		}
	}

	err = stream.Stop()
	if err != nil {
		panic(err)
	}
}

func rms(input []int32) float64 {
	const n = len(input)
	sum := int32(0)
	for _, i := range input {
		sum += i * i
	}
	return float64(sum) / float64(n)
}
