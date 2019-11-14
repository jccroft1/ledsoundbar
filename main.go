package main

import (
	"log"

	"github.com/gordonklaus/portaudio"
)

func main() {
	log.Println("starting")

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	if err != nil {
		panic(err)
	}

	err := stream.Start()
	if err != nil {
		panic(err)
	}

	for {
		stream.Read()
		log.Println(average(in))
	}

	err := stream.Stop()
	if err != nil {
		panic(err)
	}
}

func average(input []int) float64 {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return float64(sum) / float64(len(input))
}
