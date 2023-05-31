package hl7

import (
	"github.com/alhaos/BD.2/internal/fs"
	"log"
	"strings"
)

type Message struct {
	MSH *MSH
	PID *PID
	PV1 *PV1
}

func MessageFromFile(filename string) (*Message, error) {

	m := new(Message)

	lines, err := fs.ReadLines(filename)
	if err != nil {
		log.Fatalln("unable read lines from file", filename, err)
	}

	m.MSH, err = NewMSH(lines[0])
	if err != nil {
		log.Fatalln("unable parse MSH segment", err)
	}

	m.PID, err = NewPID(lines[1])
	if err != nil {
		log.Fatalln("unable parse PID segment", err)
	}

	return m, nil
}

func (m *Message) String() string {
	b := strings.Builder{}
	b.WriteString(m.MSH.String())
	b.WriteString("\n")
	b.WriteString(m.PID.String())
	return b.String()
}
