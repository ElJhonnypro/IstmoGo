package userUseModels

type User struct {
	ID        string
	Name      string
	Phone     string
	Role      string
	RID       *string
	RIDPhoto  *string
	CarID     *string
	Email     string
	Password  string
	Location  string
	Birthdate *string
	CreatedAt string
}
