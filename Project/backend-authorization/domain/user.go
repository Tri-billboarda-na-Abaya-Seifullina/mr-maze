package domain

const (
	REGISTER = "register"
)

type User struct {
	Login    string
	Password string
	Id       int
}
