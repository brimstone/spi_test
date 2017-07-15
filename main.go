package main

import (
	"fmt"

	"golang.org/x/exp/io/spi"
)

func main() {
	dev, err := spi.Open(&spi.Devfs{
		Dev:      "/dev/spidev0.1",
		Mode:     spi.Mode0,
		MaxSpeed: 500000,
	})
	if err != nil {
		panic(err)
	}
	defer dev.Close()

	w := make([]byte, 8)
	w[0] = 8
	w[1] = 11
	w[2] = 3
	w[3] = 0
	r := make([]byte, len(w))
	if err := dev.Tx(w, r); err != nil {
		panic(err)
	}

	fmt.Printf("w: %#v\n", w)
	fmt.Printf("r: %#v\n", r)

}
