package services

func (sw *SolaWind) Commit() {
	for k, v := range sw.Transactions {
		sw.Store[k] = v
	}
	mp := make(map[string]string)
	sw.Transactions = mp
	sw.IsTransaction = false
}
