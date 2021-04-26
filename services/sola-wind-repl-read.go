package services

import (
	"fmt"
	"strings"
)

func (sw *SolaWind) Read(s string) {
	trimed := strings.TrimSpace(s)
	if sw.IsTransaction {
		if v, isexist := sw.Transactions[trimed]; isexist {
			fmt.Println(v)
		} else {
			fmt.Println("key not found")
		}
	} else {
		if v, isexist := sw.Store[trimed]; isexist {
			fmt.Println(v)
		} else {
			fmt.Println("key not found")
		}
	}

}
