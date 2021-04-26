package services

import (
	"fmt"
	"strings"
)

func (sw *SolaWind) Delete(s string) {
	trimed := strings.TrimSpace(s)
	if sw.IsTransaction {
		if _, isexist := sw.Transactions[trimed]; isexist {
			delete(sw.Transactions, trimed)
		} else {
			fmt.Println("key not found")
		}
	} else {
		if _, isexist := sw.Store[trimed]; isexist {
			delete(sw.Store, trimed)
		} else {
			fmt.Println("key not found")
		}
	}
}
