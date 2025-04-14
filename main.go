package neatapi

import (
	"fmt"

	"github.com/hejingwen098/neatapi/neatlogic"
)

func main() {
	// Initialize neatClient
	neatClient := neatlogic.NewNeatClient()
	fmt.Println(neatClient)
	// cientity, err := neatClient.GetCientity(1380287128805376, 1391331125501952)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(cientity))

	// searchcientity, err := neatClient.SearchCientity(1380287128805376, "PD_sjxt-xc")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(len(searchcientity))

	// allcientity, err := neatClient.GetAllCientity(1380287128805376)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(len(allcientity))
}
