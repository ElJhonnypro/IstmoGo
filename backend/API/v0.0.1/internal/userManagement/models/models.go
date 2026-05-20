package userUseModels

type User struct {
	ID           string
	Name         string
	Phone        string
	Role         string
	RID          *string
	RIDPhoto     *string
	CarID        *string
	Email        string
	Password     string
	LocationLat  float64
	LocationLong float64
	Birthdate    *string
	CreatedAt    string
}
