package main

import (
	"flag"

	"github.com/marsevilspirit/m_RPC/example"
	"github.com/marsevilspirit/m_RPC/server"
)

var (
	addr = flag.String("addr", "localhost:30000", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterWithName("HelloWorld", new(example.HelloWorld), "")
	s.Serve("tcp", *addr)
}
