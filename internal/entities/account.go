package entities

import "errors"

type Account struct {
	Id    uint `json:"id"`
	Money uint `json:"money"` // Stored in kopeks
}

func NewAccount(id uint, money uint) *Account {
	return &Account{
		Id:    id,
		Money: money,
	}
}

func (a *Account) GetId() uint {
	return a.Id
}

func (a *Account) GetMoney() uint {
	return a.Money
}

func (a *Account) AddMoney(sum uint) {
	a.Money += sum
}

func (a *Account) SubtractMoney(sum uint) error {
	if a.Money < sum {
		return errors.New("account's balance should be greater or equal to the sum to subtract")
	}

	a.Money -= sum
	return nil
}
