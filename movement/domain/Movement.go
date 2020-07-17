package domain

// Struct of accountant movements
type Movement struct {
	Id          string
	Client      string
	Description string
	Value       float32
	Date        string
}
