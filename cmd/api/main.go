package main

import (
	"cadigo-api/protocol"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		panic(err)
	}
}
