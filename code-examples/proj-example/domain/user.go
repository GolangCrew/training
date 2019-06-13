package domain

type UserRole string

const (
	DefaultUser UserRole = "User"
	Admin       UserRole = "Admin"
)

type User struct {
	ID       string
	Login    string
	Password string
	Role     UserRole
}
