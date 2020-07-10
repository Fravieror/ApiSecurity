package infrastructure

import (
	helperInfrastructure "apiSecurity/helper/infrastructure"
	userDomain "apiSecurity/user/domain"
	"database/sql"
	"fmt"
	"log"
	"os"
)

// MySQLRepository estructura del repositorio.
type MySQLRepository struct {
	Db *sql.DB
}

// NewMySQLRepository nueva instancia del repositorio.
func NewMySQLRepository() *MySQLRepository {
	repo := &MySQLRepository{}
	db, err := sql.Open("mysql", os.Getenv("USERMYSQL")+":"+os.Getenv("PASSMYSQL")+"@tcp("+os.Getenv("IPMYSQL")+")/"+os.Getenv("DBMYSQL"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	repo.Db = db
	return repo
}

// ValidUser Valida que el usuario exista.
func (m *MySQLRepository) ValidUser(user *userDomain.User) (bool, error) {
	var chiper helperInfrastructure.CipherAES
	salt, err := m.GetUserSalt(user)
	if err != nil {
		log.Println(err)
		return false, err
	}
	password := chiper.Encrypt(user.Password + salt + os.Getenv("PEPERCYPHER"))
	query := fmt.Sprintf("SELECT * FROM user where email=%q, password=%q", user.Email, string(password))
	err = m.Db.QueryRow(query).Scan(&user.Password)
	if err != nil {
		return false, err
	}
	if user.Password != string(password) {
		return false, nil
	}
	return true, nil
}

// GetUserSalt obtiene el salt de un usuario.
func (m *MySQLRepository) GetUserSalt(user *userDomain.User) (string, error) {
	err := m.Db.QueryRow("SELECT salt FROM user where email=%q", user.Email).Scan(&user.Salt)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return user.Salt, nil
}
