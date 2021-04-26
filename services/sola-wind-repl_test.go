package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfStartTransactionIsNotEnabled(t *testing.T) {
	sw := SolaWind{}
	sw.IsTransaction = false
	sw.Store = map[string]string{}
	sw.Transactions = map[string]string{}
	assert.Falsef(t, sw.IsTransaction, "Transaction is not enabled!")
	sw.Start()
	assert.False(t, sw.IsTransaction, "Transaction is enabled by the command Start!")
}

func TestIfTransactionIsEnabled(t *testing.T) {
	sw := SolaWind{}
	sw.IsTransaction = false
	sw.Store = map[string]string{}
	sw.Transactions = map[string]string{}
	assert.True(t, sw.IsTransaction, "Transaction is enabled by the command Start!")
}

func TestIfTransactionIsEnabledByStartCmd(t *testing.T) {
	sw := SolaWind{}
	sw.IsTransaction = false
	sw.Store = map[string]string{}
	sw.Transactions = map[string]string{}
	sw.Start()
	assert.False(t, sw.IsTransaction, "Transaction is enabled by the command Start!")
}

func TestCommitIfTransactionIsOn(t *testing.T) {
	sw := SolaWind{}
	sw.IsTransaction = true
	sw.Store = map[string]string{}
	sw.Transactions = map[string]string{}
	if assert.True(t, sw.IsTransaction) {
		sw.Commit()
	}
}

func TestCommitIfTransactionIsNotOn(t *testing.T) {
	sw := SolaWind{}
	sw.IsTransaction = false
	sw.Store = map[string]string{}
	sw.Transactions = map[string]string{}
	if assert.True(t, sw.IsTransaction, "There is no transaction to commit") {
		sw.Commit()
	}
}
