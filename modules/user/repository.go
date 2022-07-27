package user

type UserRepository interface {
	CreateUser(data CreateUserDTO) (User, error)
}
