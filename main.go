package main

import (
	"context"
	"encoding/binary"
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
		/**/ state_volume := binary.BigEndian.Uint16([]byte{0x02, 0x10}) // 闲暇中，10%容积
		/**/ syn1_syn2 := binary.BigEndian.Uint16([]byte{136, 137}) // SYN = 136+1
		/**/ currentWeight := binary.BigEndian.Uint16([]byte{0x01, 0x10}) // 1.1Kg
		/**/ idCardH := binary.BigEndian.Uint16([]byte{0x01, 0x02}) //
		/**/ idCardL := binary.BigEndian.Uint16([]byte{0x03, 0x04}) // 卡号为1234
		/**/ err1_2 := binary.BigEndian.Uint16([]byte{0x00, 0x00}) // 故障:123
		/**/ err3_ := binary.BigEndian.Uint16([]byte{0x00, 0x00}) // 故障:123

		server.HoldingRegisters = []uint16{
			state_volume,
			syn1_syn2,
			currentWeight,
			idCardH, idCardL,
			err1_2, err3_,
		}
		log.Println("Modbus Server started@ 127.0.0.1:1502")

	}(context.Background())
	defer server.Close()
	signal := <-channel
	context.Background().Done()
	log.Println("Received exit signal: ", signal)
}
