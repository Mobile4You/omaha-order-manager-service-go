package usecase

import (
	"log"
	"net/http"

	"github.com/arthurstockler/omaha-order-manager-service-go/db"
	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/arthurstockler/omaha-order-manager-service-go/rediscli"
)

var (
	// UseCase interface
	use UseCase
)

// RemoteAPI exported
type RemoteAPI interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	ShowOrder(w http.ResponseWriter, r *http.Request)
	ListOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
	CreateItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	ShowItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	ListTransaction(w http.ResponseWriter, r *http.Request)
}

// Persistence exported
type Persistence interface {
	SaveOrder(models.Order) error
	SaveItem(models.Item) error
}

// UseCase exported
type UseCase struct {
	Persistence
	RemoteAPI
}

// Init exported
func Init() {
	use = UseCase{}
	Start()
}

// SaveOrder exported
func (u *UseCase) SaveOrder(o models.Order) error {
	if o.Status == models.CLOSED {
		db := db.MgoDb{}
		db.Open()
		err := db.Db.C("order").Insert(&o)
		db.Close()
		return err
	}
	re := &rediscli.ORedis{}
	err := re.PutOrder(o)
	return err
}

// SaveItem exported
func (u *UseCase) SaveItem(i models.Item) error {
	log.Print("metodo real SaveItem")
	return nil
}
