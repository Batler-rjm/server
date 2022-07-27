package user

type UserController interface {
	CreateRootUser(data CreateUserDTO) (UserPresenter, error)
}

type UserControllerImpl struct {
	service UserService
}

func (controller UserControllerImpl) CreateRootUser(data CreateUserDTO) (UserPresenter, error) {
	user, err := controller.service.CreateRootUser(data)
	return UserPresenter{
		Username: user.Username,
	}, err
}
