package data

import (
	"capstoneproject/features/negotiations"

	"gorm.io/gorm"
)

type mysqlNegotiationRepository struct {
	db *gorm.DB
}

func NewNegotiationRepository(conn *gorm.DB) negotiations.Data {
	return &mysqlNegotiationRepository{
		db: conn,
	}
}

func (repo *mysqlNegotiationRepository) SelectNegotiationsByIdUser(idUser, limit, offset int) ([]negotiations.Core, error) {
	var dataNegotiations = []Negotiation{}
	result := repo.db.Preload("User").Preload("House").Where("user_id = ?", idUser).Limit(limit).Offset(offset).Find(&dataNegotiations)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataNegotiations), nil
}

func (repo *mysqlNegotiationRepository) SelectNegotiationsByIdHouse(idHouse, limit, offset int) ([]negotiations.Core, error) {
	var dataNegotiations = []Negotiation{}
	result := repo.db.Preload("User").Preload("House").Where("house_id = ?", idHouse).Limit(limit).Offset(offset).Find(&dataNegotiations)
	if result.Error != nil {
		return nil, result.Error
	}
	return toCoreList(dataNegotiations), nil
}

func (repo *mysqlNegotiationRepository) InsertNewNegotiation(data negotiations.Core) (int, error) {
	var dataNegotiation = fromCore(data)
	result := repo.db.Create(&dataNegotiation)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}

func (repo *mysqlNegotiationRepository) CheckAlreadyNegotiation(idUser, idHouse int) (row int, err error) {
	result := repo.db.Where("user_id = ? AND house_id = ? ", idUser, idHouse).First(&Negotiation{})
	if result.Error != nil {
		return -1, result.Error
	}
	return int(result.RowsAffected), nil
}
