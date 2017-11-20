package signal

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func WaitingSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("get signal! release all resource, ready to exit.")
		os.Exit(0)
	}()
	fmt.Println("waiting for signal.")
	select {}
}
