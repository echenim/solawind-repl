package services

import (
	"strings"
)

func (sw *SolaWind) Write(s string) {

	trimed := strings.TrimSpace(s)
	data := strings.Fields(trimed)
	k := data[0]
	v := string(trimed[len(k):])
	if sw.IsTransaction {
		sw.Transactions[k] = v
	} else {
		sw.Store[k] = v
	}

}
