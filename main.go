package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tbrandon/mbserver"
)

func main() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGABRT)
	server := mbserver.NewServer()
	go func(ctx context.Context) {
		if err := recover(); err != nil {
			log.Println(err)
		}
		select {
		case <-ctx.Done():
			return
		default:
			{

			}

		}
		err := server.ListenTCP("0.0.0.0:1502")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		server.Debug = true
		server.Coils = []byte{0, 1}
		server.DiscreteInputs = []byte{0, 1}
		server.InputRegisters = []uint16{0, 1}
		server.HoldingRegisters = []uint16{0, 1}

		log.Println("Modbus Server started")

	}(context.Background())
	defer server.Close()
	signal := <-channel
	context.Background().Done()
	log.Println("Received exit signal: ", signal)
}
