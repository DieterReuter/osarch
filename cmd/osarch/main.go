package main

import (
	"fmt"

	"github.com/dieterreuter/osarch"
)

func main() {
	fmt.Println(osarch.Arch())
	fmt.Println(osarch.OsArch())
}
