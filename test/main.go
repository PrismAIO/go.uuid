package main

import (
	"fmt"

	"github.com/PrismAIO/go.uuid"
)

func main() {
	uuid, _ := uuid.NewV1()

	fmt.Println(uuid.String())
}