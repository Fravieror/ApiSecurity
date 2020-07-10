package application

import (
	helperDomain "apiSecurity/helper/domain"
	"apiSecurity/token/domain"
	userDomain "apiSecurity/user/domain"
)

type Bussines struct {
	repository domain.Repository
}

func NewBussines(rep domain.Repository) domain.Service {
	return &Bussines{
		repository: rep,
	}
}

func (b *Bussines) GenerateToken(user *userDomain.User) *helperDomain.Response {
	var response = helperDomain.NewResponse()
	ok, err := b.repository.ValidUser(user)
	if err != nil {
		response.Errors = "Error internal server"
		response.Code = "500"
	}
	if !ok {
		response.Errors = "User invalid"
		response.Code = "500"
	}

	return response
}
