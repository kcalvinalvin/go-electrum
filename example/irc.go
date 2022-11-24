package main

import (
	"log"

	"github.com/kcalvinalvin/go-electrum/irc"
)

func main() {
	log.Println(irc.FindElectrumServers())
}
