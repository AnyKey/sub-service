package user

// Usecase represent the User's usecases.
type Usecase interface {
	Token(string) (*string, error)
}

type GrpcDelivery interface {
	GetToken(string) (*string, error)
}
