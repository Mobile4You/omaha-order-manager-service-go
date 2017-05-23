package usecase

import (
	"testing"

	"github.com/arthurstockler/omaha-order-manager-service-go/models"
)

func TestBuildItem(t *testing.T) {

	i := models.Item{}
	buildItem(&i)
	if len(i.UUID.Hex()) == 0 {
		t.Error("Não gerou o UUID corretamente")
	}

	if i.CreatedAt.IsZero() || i.UpdatedAt.IsZero() {
		t.Error("Value of data fields, can not be null")
	}
}

func BenchmarkBuildItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := models.Item{}
		buildItem(&i)
	}
}
