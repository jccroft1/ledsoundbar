# ledsoundbar

A Go program that visualizes an audio input on an LED strip. Toggled by HTTP requests. 

## Installation

- Install Golang 
- Install PortAudio 
- Clone this repo to the Golang path
- Install the packages 
- Run the program 
```
sudo apt-get install golang
sudo apt-get install portaudio19-dev
cd $HOME/go/src/
go get github.com/jccroft1/ledsoundbar
cd github.com/jccroft1/ledsoundbar
go get -d -v .
go run main.go > output.txt & 
```
