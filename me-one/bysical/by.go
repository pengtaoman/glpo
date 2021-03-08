package bysical

import (
	"fmt"
	"strconv"
)

type Bysical struct {
	Speed int
	Name  string
}

func (by Bysical) Run(speed int) string {
	by.Speed = speed
	// return "dfasdfasdfasdf======================"
	return fmt.Sprintf(by.Name + " 's Speed is :: " + strconv.Itoa(by.Speed))
}
