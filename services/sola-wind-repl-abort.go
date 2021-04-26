package services

func (sw *SolaWind) Abort() {
	mp := make(map[string]string)
	sw.Transactions = mp
}
