package usecase

import "github.com/arthurstockler/omaha-order-manager-service-go/caching"

var (
	// UseCase interface
	use   UseCase
	cache caching.RedisCache
)

// UseCase exported
type UseCase struct {
	Persistence
	RemoteAPI
}

// Init exported
func Init() {
	cache = caching.RedisCache{}
	use = UseCase{}
	ServerStart()
}
