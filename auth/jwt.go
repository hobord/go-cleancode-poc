package auth

type jwtPayload struct {
	exp   int
	sub   string // user id
	name  string
	email string
	photo string
	roles []string
}
