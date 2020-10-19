package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"translation/model"
)

//InsertStandardCode inserts a new code value based on
func InsertStandardCode(sc model.StandardCode) (sql.Result, error) {
	if len(sc.Code) > 50 {
		return nil, fmt.Errorf("standard code too long.  do not exceed 50 characters")
	}
	if len(sc.Description) > 50 {
		return nil, fmt.Errorf("standard description too long.  do not exceed 50 characters")
	}

	qry := `SELECT code_type_id FROM code_type WHERE description = $1;`

	row := db.QueryRow(qry, sc.CodeType.Description)

	var id int
	var qry2 string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("code type not found.  please verify")
	case nil:
		qry2 = "INSERT INTO standard_code (code, description, type_id, created_user, modified_user) VALUES ('" + sc.Code + "', '" + sc.Description + "', '" + strconv.Itoa(id) + "', '" + appUser + "', '" + appUser + "');"
	default:
		return nil, fmt.Errorf("error finding code type  please verify")
	}

	r, err := db.Exec(qry2)
	if err != nil {
		return nil, err
	}
	return r, nil
}

//SelectStandardCodes allows a user to select a set of standard codes.
//If string t is empty, the function will return ALL standard codes.
//If t is not empty, it will evaluate the code type table and only return
//values for the associated code type.
func SelectStandardCodes(t string) ([]model.StandardCodeList, error) {
	var scList []model.StandardCodeList
	var scl model.StandardCodeList
	var sclr model.StandardCodeList
	var sc model.StandardCodes
	var scr model.StandardCodes

	var qry string
	if t != "" {
		qry = "SELECT * FROM code_type WHERE description = $1;"

	} else {
		t = "1"
		qry = "SELECT * FROM code_type WHERE 1 = $1;"
	}
	r, err := db.Query(qry, t)
	if err != nil {
		return nil, err
	}

	defer r.Close()

	for r.Next() {
		err := r.Scan(&sclr.ID, &sclr.UuID, &sclr.Description, &sclr.Metadata.CreatedDT, &sclr.Metadata.CreatedBy, &sclr.Metadata.ModifiedDT, &sclr.Metadata.ModifiedBy)
		if err != nil {
			return nil, err
		}
		scl.UuID = sclr.UuID
		scl.Description = sclr.Description

		qry2 := `SELECT * FROM standard_code WHERE type_id = $1;`

		r2, err := db.Query(qry2, sclr.ID)
		if err != nil {
			return nil, err
		}
		defer r2.Close()

		for r2.Next() {
			err := r2.Scan(&scr.ID, &scr.UuID, &scr.TypeID, &scr.Code, &scr.Description, &scr.Metadata.CreatedDT, &scr.Metadata.CreatedBy, &scr.Metadata.ModifiedDT, &scr.Metadata.ModifiedBy)
			if err != nil {
				return nil, err
			}
			sc.UuID = scr.UuID
			sc.Code = scr.Code
			sc.Description = scr.Description

			sc.Metadata.CreatedBy = scr.Metadata.CreatedBy
			sc.Metadata.CreatedDT = scr.Metadata.CreatedDT
			sc.Metadata.ModifiedBy = scr.Metadata.ModifiedBy
			sc.Metadata.ModifiedDT = scr.Metadata.ModifiedDT

			scl.StandardCodes = append(scl.StandardCodes, sc)
		}
		scList = append(scList, scl)
	}
	return scList, nil
}
