package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tbrandon/mbserver"
)

var Coils []byte = make([]byte, 1000)
var DiscreteInputs = make([]byte, 1000)
var InputRegisters = make([]uint16, 1000)
var HoldingRegisters = make([]uint16, 1000)

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
		//
		server.Coils = Coils
		server.Coils = Coils
		server.DiscreteInputs = DiscreteInputs
		server.InputRegisters = InputRegisters
		server.HoldingRegisters = HoldingRegisters
		//
		log.Println("Modbus Server started: 0.0.0.0:1502")

	}(context.Background())
	defer server.Close()
	signal := <-channel
	context.Background().Done()
	log.Println("Received exit signal: ", signal)
}
