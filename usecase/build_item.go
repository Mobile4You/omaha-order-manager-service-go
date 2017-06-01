package usecase

import (
	"strings"
	"time"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
	"github.com/satori/go.uuid"
)

// compila um item
func buildItem(i *models.Item) {
	if len(strings.TrimSpace(i.UUID)) == 0 {
		i.UUID = uuid.NewV4().String()
		i.CreatedAt = time.Now()
	}
	i.UpdatedAt = time.Now()
}
