package models

type RequestBody struct {
	PetId     uint `json:"petId"`
	SpeciesId uint `json:"speciesId"`
}
