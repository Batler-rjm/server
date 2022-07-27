package user

import "testing"

type UserServiceMock struct{}

func (service UserServiceMock) CreateRootUser(data CreateUserDTO) (User, error) {
	return User(data), nil
}

var TestCreateUserDTO = CreateUserDTO{
	Username: "testUsername",
	Password: "testPassword",
}

func TestCreateRootUser_UserController(test *testing.T) {
	service := UserServiceMock{}
	controller := UserControllerImpl{service}
	user, err := controller.CreateRootUser(TestCreateUserDTO)
	if err != nil {
		test.Error(err)
	}
	if user.Username != TestCreateUserDTO.Username {
		test.Error("Invalid user name")
	}
}
