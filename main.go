package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gordonklaus/portaudio"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	log.Println("starting")

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	if err != nil {
		panic(err)
	}

	err = stream.Start()
	if err != nil {
		panic(err)
	}

	for {
		stream.Read()
		log.Println(average(in))
		select {
		case <-sig:
			break
		default:
		}
	}

	err = stream.Stop()
	if err != nil {
		panic(err)
	}
}

func average(input []int32) float64 {
	sum := int32(0)
	for _, i := range input {
		sum += i
	}
	return float64(sum) / float64(len(input))
}
