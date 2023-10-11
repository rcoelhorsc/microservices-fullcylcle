package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatedNewClient(t *testing.T) {
	client, err := NewClient("Renato", "rcoelho@email")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Renato", client.Name)
	assert.Equal(t, "rcoelho@email", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Renato", "rcoelho@email")
	err := client.Update("Renato Update", "rcoelho@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "Renato Update", client.Name)
	assert.Equal(t, "rcoelho@email.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Renato", "rcoelho@email")
	err := client.Update("", "rcoelho@email.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Renato", "rcoelho@email")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
