package user

// Usecase represent the User's usecases.
type Usecase interface {
	Token() error
}

// Repository represent the User's repository contract.
type Repository interface {
	GetOption() error
}

// HttpDelivery represent the User's transport
type HttpDelivery interface {
	GetToken() error
}
