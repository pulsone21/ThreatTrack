package stores

import (
	"ContentManagement/api/models/user"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	iS := &UserStore{
		DB: db,
	}
	iS.createTable()
	return iS
}

func (s *UserStore) createTable() {
	fmt.Println("Try to create user.User Table")
	query, err := LoadSQL("user/CreateTable.sql")
	if err != nil {
		panic(err.Error())
	}
	if _, err := s.DB.Exec(query); err != nil {
		panic(err.Error())
	}
	fmt.Println("No Issues found on table creation")
}

func (s *UserStore) GetAllUser() ([]user.User, error) {
	var users []user.User
	query, err := LoadSQL("user/GetAll.sql")
	if err != nil {
		return nil, err
	}
	res, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		var usr user.User
		fmt.Sprintln(res)
		err := res.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Created_at, &usr.Fullanme)
		if err != nil {
			return nil, err
		}
		users = append(users, usr)
	}
	return users, nil
}

func (s *UserStore) GetUserByID(id string) (*user.User, error) {
	var usr user.User
	query, err := LoadSQL("user/GetById.sql")
	if err != nil {
		return nil, err
	}
	err = s.DB.QueryRow(query, id).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Created_at)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (s *UserStore) CreateUser(user *user.User) error {
	query, err := LoadSQL("user/Create.sql")
	if err != nil {
		return err
	}
	_, err = s.DB.Exec(query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Created_at,
		user.Fullanme)
	return err
}

func (s *UserStore) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("not implemented")
}

func (s *UserStore) DeleteUser(id string) error {
	query, err1 := LoadSQL("user/Delete.sql")
	_, err := s.DB.Exec(query, id)
	return errors.Join(err1, err)
}
