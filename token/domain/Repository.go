package domain

import (
	"apiSecurity/user/domain"
)

type Repository interface {
	ValidUser(*domain.User) (bool, error)
}
