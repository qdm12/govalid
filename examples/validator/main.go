package main

import (
	"fmt"

	"github.com/qdm12/govalid"
	"github.com/qdm12/govalid/address"
)

func main() {
	const s = ":8000"
	const uid = 1000

	err := govalid.ValidateAddress(s, address.OptionListening(uid))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Address", s, "is valid")
}
