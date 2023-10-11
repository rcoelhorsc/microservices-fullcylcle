package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Renato Coelho", "rcoelho@email")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditedAccount(t *testing.T) {
	client, _ := NewClient("Renato Coelho", "rcoelho@email")
	account := NewAccount(client)
	account.Credit(100)
	assert.NotNil(t, account)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitedAccount(t *testing.T) {
	client, _ := NewClient("Renato Coelho", "rcoelho@email")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(45)
	assert.NotNil(t, account)
	assert.Equal(t, float64(55), account.Balance)
}
