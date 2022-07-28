package user

import "testing"

type DatabaseMock struct{}

func (db DatabaseMock) Insert(table string, data CreateUserDTO) (User, error) {
	return User(data), nil
}

func TestUserRepository(test *testing.T) {
	db := DatabaseMock{}
	repository := UserRepositoryImpl{db: db}
	dto := CreateUserDTO{
		Username: "teste",
		Password: "teste",
	}
	created, err := repository.CreateUser(dto)
	if err != nil {
		test.Error(err)
	}
	if created != User(dto) {
		test.Error("Invalid created user")
	}
}
