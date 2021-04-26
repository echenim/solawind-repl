package services

import (
	"fmt"
	"os"
)

func (sw *SolaWind) Quit() {
	fmt.Println("Exiting ...")
	os.Exit(0)
}
