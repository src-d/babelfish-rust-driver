package main

import (
	_ "github.com/bblfsh/rust-driver/driver/impl"
	"github.com/bblfsh/rust-driver/driver/normalizer"

	"gopkg.in/bblfsh/sdk.v2/driver/server"
)

func main() {
	server.Run(normalizer.Transforms)
}
