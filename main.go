package main

import (
	"fmt"
	"module/common/services"
)

func main() {
	fmt.Println("comeca aqui")
	services.Start()
	fmt.Println("termina aqui")
}
