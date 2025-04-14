package neatapi

import (
	"fmt"

	"github.com/hejingwen098/neatapi/neatlogic"
)

func main() {
	// Initialize neatClient
	neatClient := neatlogic.NewNeatClient()
	fmt.Println(neatClient)
}
