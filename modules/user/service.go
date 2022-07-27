package user

type UserService interface {
	CreateRootUser(data CreateUserDTO) User
}

type Hasher interface {
	Hash(value string) (string, error)
}

type UserServiceImpl struct {
	repository UserRepository
	hasher     Hasher
}

func (service UserServiceImpl) CreateRootUser(data CreateUserDTO) (User, error) {
	hash, err := service.hasher.Hash(data.Password)
	if err != nil {
		return User{}, err
	}
	user, err := service.repository.CreateUser(CreateUserDTO{
		Username: data.Username,
		Password: hash,
	})
	return user, err
}
