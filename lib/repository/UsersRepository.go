package repository

import (
	"database/sql"
	"fmt"

	"github.com/forge-io/backend/lib/models"
	"github.com/google/uuid"
)

type UserRepository struct{
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository{
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetAllUsers()([]models.User, error) {
	query := "SELECT user_id, user_name, user_age, user_email, user_password, user_phone, user_address FROM users_schema.table_users"
	rows, err := ur.connection.Query(query)
	if (err != nil) {
		fmt.Println(err)
		return []models.User{}, err
	}
	var userList []models.User
	var userObj models.User

	for rows.Next(){
		err = rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Age,
			&userObj.Email,
			&userObj.Password,
			&userObj.Phone,
			&userObj.Address)

			if (err != nil){
				fmt.Println(err)
				return []models.User{}, err
			}

			userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (ur *UserRepository) CreateUser(user models.User) (uuid.UUID, error) {

	var user_id uuid.UUID

	query, err := ur.connection.Prepare("INSERT INTO users_schema.table_users" +
	"(user_name, user_age, user_email, user_password, user_phone, user_address)" +
	"VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id")
	if err != nil {
		fmt.Println(err)
		return user_id, err
	}
	err = query.QueryRow(user.Name, user.Age, user.Email, user.Password, user.Phone, user.Phone).Scan(&user_id)
	if err != nil {
		fmt.Println(err)
		return user_id, err
	}

	query.Close()
	return user_id, nil
}