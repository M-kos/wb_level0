package repositories

import (
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/models"
)

type addressRepository struct {
	db *db.PostgresDB
}

func NewAddressRepository(db *db.PostgresDB) *addressRepository {
	return &addressRepository{db: db}
}

func (ar *addressRepository) Create(address models.AddressDbModel) (int, error) {

}

func (ar *addressRepository) Find(id int) (models.Address, error) {}
