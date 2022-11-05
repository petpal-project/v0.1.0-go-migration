package models

import "gorm.io/gorm"

type Species struct {
	gorm.Model
	Name         string `json:"name"`
	BinomialName string `json:"binomialName"`
}

func RetrieveSpecies(speciesId uint, DB *gorm.DB) *Species {
	var species *Species
	DB.First(&species, "id = ?", speciesId)
	return species
}
