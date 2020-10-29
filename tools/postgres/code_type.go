package postgres

import (
	"database/sql"
	"fmt"
	"translation/model"
)

//InsertCodeType inserts a new code type into the
//translation database, using a custom struct as an input
func InsertCodeType(c model.CodeType) (sql.Result, error) {
	if len(c.Description) > 50 {
		return nil, fmt.Errorf("code description too long.  do not exceed 50 characters")
	}

	qry := "INSERT INTO code_type (description, created_user, modified_user) VALUES ('" + c.Description + "', '" + model.AppUser + "', '" + model.AppUser + "');"
	r, err := model.Db.Exec(qry)
	if err != nil {
		return nil, err
	}
	return r, nil
}
