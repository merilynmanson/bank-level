package entities

import "errors"

type Account struct {
	Id    uint    `json:"id"`
	Money float32 `json:"money"`
}

func NewAccount(id uint, money float32) *Account {
	return &Account{
		Id:    id,
		Money: money,
	}
}

func (a *Account) GetId() uint {
	return a.Id
}

func (a *Account) GetMoney() float32 {
	return a.Money
}

func (a *Account) AddMoney(sum float32) {
	a.Money += sum
}

func (a *Account) SubtractMoney(sum float32) error {
	if sum < 0 {
		return errors.New("sum to subtract should be a positive number")
	}
	if a.Money < sum {
		return errors.New("account's balance should be greater or equal to the sum to subtract")
	}

	a.Money -= sum
	return nil
}
