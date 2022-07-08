package storage

import (
	"bank/internal/entities"
	"errors"
	"fmt"
	"sync"
)

type AccountStorage struct {
	sync.RWMutex
	accounts map[uint]*entities.Account
}

func (a *AccountStorage) PrintAccs() {
	for _, acc := range a.accounts {
		fmt.Println(*acc)
	}
}

func NewAccountStorage() *AccountStorage {
	return &AccountStorage{
		accounts: make(map[uint]*entities.Account, 1000000),
	}
}

func (a *AccountStorage) AddAccount(money float32) {
	a.Lock()
	defer a.Unlock()
	var maxId uint = 0
	for id := range a.accounts {
		if id > maxId {
			maxId = id
		}
	}
	a.accounts[maxId+1] = entities.NewAccount(maxId+1, money)
}

func (a *AccountStorage) GetAccount(id uint) (*entities.Account, bool) {
	a.RLock()
	account, ok := a.accounts[id]
	a.RUnlock()
	return account, ok
}

func (a *AccountStorage) AddMoney(id uint, sum float32) error {
	account, ok := a.GetAccount(id)
	if !ok {
		return errors.New("account doesn't exist")
	}
	a.Lock()
	defer a.Unlock()
	account.AddMoney(sum)
	return nil
}

func (a *AccountStorage) SubtractMoney(id uint, sum float32) error {
	account, ok := a.GetAccount(id)
	if !ok {
		return errors.New("account doesn't exist")
	}
	a.Lock()
	defer a.Unlock()
	err := account.SubtractMoney(sum)
	if err != nil {
		return err
	}

	return nil
}
