package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/originDoris/weakscope/internal/probe"
)

func main() {
	addr := flag.String("target", "", "探测目标，host:port 形式")
	flag.Parse()
	results := probe.Probe(context.Background(), probe.NetDialer{}, *addr)
	for _, result := range results {
		marshal, _ := json.Marshal(result)
		fmt.Println(string(marshal))
	}
}
