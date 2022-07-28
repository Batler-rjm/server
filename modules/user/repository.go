package user

type UserRepository interface {
	CreateUser(data CreateUserDTO) (User, error)
}

type Database[Insert any, Retrieve any] interface {
	Insert(table string, data Insert) (Retrieve, error)
}

type UserRepositoryImpl struct {
	db Database[CreateUserDTO, User]
}

func (repository UserRepositoryImpl) CreateUser(data CreateUserDTO) (User, error) {
	created, err := repository.db.Insert("user", data)
	return created, err
}
