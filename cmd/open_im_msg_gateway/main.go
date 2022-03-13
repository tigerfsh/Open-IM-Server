package main

import (
	"Open_IM/internal/msg_gateway/gate"
	"flag"
)

func main() {
	rpcPort := flag.Int("rpc_port", 10400, "rpc listening port")
	wsPort := flag.Int("ws_port", 17778, "ws listening port")
	flag.Parse()

	//等价的实现
	ch := make(chan int)
	gate.Init(*rpcPort, *wsPort)
	gate.Run()
	ch <- 1

	// var wg sync.WaitGroup
	// wg.Add(1)
	// gate.Init(*rpcPort, *wsPort)
	// gate.Run()
	// wg.Wait()
}
