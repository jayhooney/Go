package main

import (
	"fmt"

	"./secret"
)

func main() {
	fmt.Println(secret.GetDBInfo())
}
