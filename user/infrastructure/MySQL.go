package infrastructure

import (
	helperInfrastructure "apiSecurity/helper/infrastructure"
	userDomain "apiSecurity/user/domain"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

// MySQL struct to refernece mySql.
type MySQL struct {
	Db *sql.DB
}

// NewMySQL create new instances of repository.
func NewMySQL() *MySQL {
	repo := &MySQL{}
	db, err := sql.Open("mysql", os.Getenv("USERMYSQL")+":"+os.Getenv("PASSMYSQL")+"@tcp("+os.Getenv("IPMYSQL")+")/"+os.Getenv("DBMYSQL"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	repo.Db = db
	return repo
}

// SaveUser Save users.
func (m *MySQL) SaveUser(user *userDomain.User) error {
	repo := NewMySQL()
	var chiper helperInfrastructure.CipherAES
	var randomGenerator helperInfrastructure.CryptoSource
	salt := randomGenerator.GenRandom()
	password := chiper.Encrypt(user.Password + strconv.Itoa(salt) + os.Getenv("PEPERCYPHER"))
	query := fmt.Sprintf("INSERT INTO users VALUES ( %q, %q, %q)", user.Email, salt, string(password))
	insert, err := repo.Db.Query(query)
	defer insert.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
