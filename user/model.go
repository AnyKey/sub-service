package user

// Usecase represent the User's usecases.
type Usecase interface {
	Token() error
}

// HttpDelivery represent the User's transport
type HttpDelivery interface {
	GetToken() error
}
