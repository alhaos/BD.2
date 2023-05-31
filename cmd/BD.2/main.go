package main

import (
	"fmt"
	"github.com/alhaos/BD.2/internal/hl7"
	"log"
)

const filename = `D:\repository\BD\test\arch\1681409804645.dat`

func main() {
	m, err := hl7.MessageFromFile(filename)
	if err != nil {
		log.Fatalln("unable parse file", filename, err)
	}

	fmt.Println(m.String())
}
