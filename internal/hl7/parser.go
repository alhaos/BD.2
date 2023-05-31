package hl7

import (
	"errors"
	"fmt"
	"github.com/alhaos/BD.2/internal/fs"
)

func Parse(filename string) (interface{}, error) {

	lines, err := fs.ReadLines(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %s", filename, err.Error())
	}

	msh, err := NewMSH(lines[0])
	if err != nil {
		return nil, fmt.Errorf("unable parse MSH segment from line [%s] %s", lines[0], err.Error())
	}

	pid, err := NewPID(lines[1])
	if err != nil {
		return nil, fmt.Errorf("unable parce PID segment from line [%s] %s", lines[1], err.Error())
	}

	pv1, err := NewPV1(lines[2])

	return nil, errors.New("")
}
