package neatapi

import (
	"fmt"

	"github.com/hejingwen098/neatapi/auth"
)

func main() {
	token, err := auth.Login()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
