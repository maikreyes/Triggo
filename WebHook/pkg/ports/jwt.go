package ports

type JWTServices interface {
	CreateJWT() (string, error)
}
