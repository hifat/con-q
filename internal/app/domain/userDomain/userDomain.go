package userDomain

type IUserRepo interface {
	Exists(col string, expected string) (bool, error)
}
