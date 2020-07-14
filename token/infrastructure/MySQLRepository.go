package infrastructure

import (
	helperInfrastructure "apiSecurity/helper/infrastructure"
	userDomain "apiSecurity/user/domain"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLRepository estructura del repositorio.
type MySQLRepository struct {
	Db *sql.DB
}

// NewMySQLRepository nueva instancia del repositorio.
func NewMySQLRepository() *MySQLRepository {
	repo := &MySQLRepository{}
	db, err := sql.Open("mysql", os.Getenv("USERMYSQL")+":"+os.Getenv("PASSMYSQL")+"@tcp("+os.Getenv("IPMYSQL")+")/"+os.Getenv("DBMYSQLUSERS"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	repo.Db = db
	return repo
}

// ValidUser Valida que el usuario exista.
func (m *MySQLRepository) ValidUser(user *userDomain.User) (bool, error) {
	repo := NewMySQLRepository()
	var chiper helperInfrastructure.CipherAES
	salt, err := repo.GetUserSalt(user)
	if err != nil {
		log.Println(err)
		return false, err
	}
	password := chiper.Encrypt(user.Password + salt + os.Getenv("PEPERCYPHER"))
	query := fmt.Sprintf("SELECT Password FROM user where email=%q", user.Email)
	err = repo.Db.QueryRow(query).Scan(&user.Password)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if user.Password != string(password) {
		return true, nil //cambiar
	}
	return true, nil
}

// GetUserSalt obtiene el salt de un usuario.
func (m *MySQLRepository) GetUserSalt(user *userDomain.User) (string, error) {
	err := m.Db.QueryRow(fmt.Sprintf("SELECT salt FROM user where email=%q", user.Email)).Scan(&user.Salt)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return user.Salt, nil
}
