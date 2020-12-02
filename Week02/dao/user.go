package dao

import (
	"database/sql"

	"github.com/GO-000/week02/model"
	"github.com/pkg/errors"
)

// GetUser is get all user from db
func GetUser() (usr []model.User, e error) {
	db, err := sql.Open("mysql", "")
	if err != nil {
		return nil, errors.Wrapf(err, "open db failed")
	}

	rows, err := db.Query("")
	if err != nil {
		return nil, errors.Wrapf(err, "query db failed")
	}

	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.ID, &u.Name)
		usr = append(usr, u)
	}
	return usr, nil
}
