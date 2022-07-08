package storage

import (
	"bank/internal/entities"
	"errors"
	"sync"
)

type AccountStorage struct {
	sync.RWMutex
	accounts map[uint]*entities.Account
}

func NewAccountStorage() *AccountStorage {
	return &AccountStorage{
		accounts: make(map[uint]*entities.Account, 1000000),
	}
}

func (a *AccountStorage) AddAccount(money uint) {
	var maxId uint = 0
	for id := range a.accounts {
		if id > maxId {
			maxId = id
		}
	}
	a.accounts[maxId+1] = entities.NewAccount(maxId+1, money)
}

func (a *AccountStorage) GetAccount(id uint) (*entities.Account, bool) {
	account, ok := a.accounts[id]
	return account, ok
}

func (a *AccountStorage) AddMoney(id uint, sum uint) error {
	account, ok := a.GetAccount(id)
	if !ok {
		return errors.New("account doesn't exist")
	}
	account.AddMoney(sum)
	return nil
}

func (a *AccountStorage) SubtractMoney(id uint, sum uint) error {
	account, ok := a.GetAccount(id)
	if !ok {
		return errors.New("account doesn't exist")
	}
	err := account.SubtractMoney(sum)
	if err != nil {
		return err
	}

	return nil
}
