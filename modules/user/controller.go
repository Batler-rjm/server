package user

type UserController interface {
	CreateRootUser(data CreateUserDTO) (UserPresenter, error)
}
