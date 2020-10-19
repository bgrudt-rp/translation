package postgres

import (
	"database/sql"
	"fmt"
	"translation/model"
)

//InsertSourceSystemApplication inserts a new
//application for use in qualifying data sources
func InsertSourceSystemApplication(ssa model.SourceSystemApplication) (sql.Result, error) {
	if len(ssa.Name) > 50 {
		return nil, fmt.Errorf("application name too long.  do not exceed 50 characters")
	}
	if len(ssa.Vendor) > 50 {
		return nil, fmt.Errorf("vendor name too long.  do not exceed 50 characters")
	}

	qry := "INSERT INTO source_system_application (application_name, vendor_name, created_user, modified_user) VALUES ('" + ssa.Name + "', '" + ssa.Vendor + "', '" + appUser + "', '" + appUser + "');"
	fmt.Println(qry)
	r, err := db.Exec(qry)
	if err != nil {
		return nil, err
	}
	return r, nil
}
