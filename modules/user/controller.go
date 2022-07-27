package user

type UserController interface {
	CreateRootUser(data CreateRootUser) UserPresenter
}
