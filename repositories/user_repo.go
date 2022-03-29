package repositories

import (
	"database/sql"
	"errors"
	"log"
	"simple-upload-file/models"
	"strings"
)

type IUserRepository interface {
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User, id string) (*models.User, error)
	DeleteUser(id string) (int64, error)
}

type userRepo struct {
	db *sql.DB
}

// Constructor
func NewUserRepo(db *sql.DB) IUserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) GetUser(id string) (*models.User, error) {
	var user models.User

	queryGetUser := `SELECT ID,firstname,lastname,age,picture FROM user WHERE ID = ? AND is_active = 1`

	row := u.db.QueryRow(queryGetUser, id)

	errScan := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Age, &user.Picture)
	if errScan != nil {
		log.Println(errScan.Error())
		return nil, errors.New("error scan get user")
	}

	return &user, nil
}

func (u *userRepo) CreateUser(user *models.User) (*models.User, error) {

	queryInsertUser := `INSERT INTO user (ID,firstname,lastname,age,picture) VALUES (?,?,?,?,?)`
	result, err := u.db.Exec(queryInsertUser, &user.ID, &user.Firstname, &user.Lastname, &user.Age, &user.Picture)
	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = string(rune(id))

	return user, nil
}

func (u *userRepo) UpdateUser(user *models.User, id string) (*models.User, error) {
	// query update
	queryUpdateUser := `UPDATE user SET firstname = ?,lastname = ?,age = ?,picture = ? WHERE ID = ?`
	result, err := u.db.Exec(queryUpdateUser, &user.Firstname, &user.Lastname, &user.Age, &user.Picture, id)
	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepo) DeleteUser(id string) (int64, error) {
	query := `UPDATE user SET is_active=0 WHERE ID = ?`
	result, err := u.db.Exec(query, id)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	newID, _ := result.RowsAffected()
	return newID, nil
}
