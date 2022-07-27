package user

type UserService interface {
	CreateRootUser(data CreateRootUser) User
}
