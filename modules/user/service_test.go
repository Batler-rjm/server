package user

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

type UserRepositoryMock struct{}

func (repository UserRepositoryMock) CreateUser(data CreateUserDTO) (User, error) {
	return User(data), nil
}

type HasherMock struct{}

func (hasher HasherMock) Hash(text string) (string, error) {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

var (
	testPassword = "testPassword"
	testUsername = "testUsername"
)

func TestCreateRootUser(test *testing.T) {
	repository := UserRepositoryMock{}
	hasher := HasherMock{}
	service := UserServiceImpl{repository, hasher}
	user, err := service.CreateRootUser(CreateUserDTO{
		Password: testPassword,
		Username: testUsername,
	})
	if err != nil {
		test.Error(err)
	}
	if user.Username != testUsername {
		test.Error("Error on username")
	}
	if user.Password == testPassword {
		test.Error("Password not hashed")
	}
}
