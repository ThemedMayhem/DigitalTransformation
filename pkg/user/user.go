package user

type user struct {
	ID       uint32
	Login    string
	Password string
	Role     string
}

type UserRepo interface {
	Authorize(login, password string) (*user, error)
	AddUser(login, password, role string) error
}
