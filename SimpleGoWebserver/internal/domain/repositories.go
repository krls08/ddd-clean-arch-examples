package domain

//go:generate moq -out "../adapters/repositories/user/mock.gen.go" -pkg userrepositories . UserRepository:MockUserRepository
type UserRepository interface {
	Create(user *User) error
	Get(id uint64) (*User, error)
	// Update accept id of user, get it and pass in updFunc.
	// updFunc should update the user and return the updated user
	Update(
		id uint64, updFunc func(*User) (*User, error),
	) error
	Delete(id uint64) error
}

//go:generate moq -out "../adapters/repositories/btc/mock.gen.go" -pkg btcrepositories . BTCRepository:MockBTCRepository
type BTCRepository interface {
	Get() BTCPrice
	SetPrice(price USD) error
}
