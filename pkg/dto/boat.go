package dto

type BoatDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateUpdateBoatDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
