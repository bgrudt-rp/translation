package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"translation/model"
)

func InsertCodeType(c string) (sql.Result, error) {
	if len(c) > 50 {
		return nil, fmt.Errorf("code description too long.  do not exceed 50 characters")
	}

	qry := "INSERT INTO code_type (description, created_user, updated_user) VALUES ('" + c + "', '" + appUser + "', '" + appUser + "');"
	r, err := db.Exec(qry)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func InsertSourceSystem(ss model.SourceSystem) (sql.Result, error) {
	if len(ss.Description) > 50 {
		return nil, fmt.Errorf("description too long.  do not exceed 50 characters")
	}

	qry := `SELECT source_system_application_id FROM source_system_application WHERE application_name = $1;`

	row := db.QueryRow(qry, ss.Application.Name)

	var id int
	var qry2 string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		qry2 = "INSERT INTO source_system (description, created_user, updated_user) VALUES ('" + ss.Description + "', '" + appUser + "', '" + appUser + "');"
	case nil:
		qry2 = "INSERT INTO source_system (application_id, description, created_user, updated_user) VALUES ('" + strconv.Itoa(id) + "', '" + ss.Description + "', '" + appUser + "', '" + appUser + "');"
	default:
		return nil, err
	}

	r, err := db.Exec(qry2)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func InsertSourceSystemApplication(ssa model.SourceSystemApplication) (sql.Result, error) {
	if len(ssa.Name) > 50 {
		return nil, fmt.Errorf("application name too long.  do not exceed 50 characters")
	}
	if len(ssa.Vendor) > 50 {
		return nil, fmt.Errorf("vendor name too long.  do not exceed 50 characters")
	}

	qry := "INSERT INTO source_system_application (application_name, vendor_name, created_user, updated_user) VALUES ('" + ssa.Name + "', '" + ssa.Vendor + "', '" + appUser + "', '" + appUser + "');"
	fmt.Println(qry)
	r, err := db.Exec(qry)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func SelectSourceSystems() ([]model.SourceSystem, error) {
	var ssList []model.SourceSystem
	var ss model.SourceSystem

	type saReturn struct {
		id        int
		uuID      string
		name      string
		vendor    string
		createdDT string
		createdBy string
		updatedDT string
		updatedBy string
	}

	var sar saReturn

	type ssReturn struct {
		id        int
		uuID      string
		appID     int
		desc      string
		createdDT string
		createdBy string
		updatedDT string
		updatedBy string
	}

	var ssr ssReturn

	qry := "SELECT * FROM source_system;"
	fmt.Println(qry)
	r, err := db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		err := r.Scan(&ssr.id, &ssr.uuID, &ssr.appID, &ssr.desc, &ssr.createdDT, &ssr.createdBy, &ssr.updatedDT, &ssr.updatedBy)
		if err != nil {
			return nil, err
		}
		ss.UuID = ssr.uuID
		ss.Description = ssr.desc
		qry := `SELECT * FROM source_system_application WHERE source_system_application_id = $1;`

		row := db.QueryRow(qry, ssr.appID)

		switch err := row.Scan(&sar.id, &sar.uuID, &sar.name, &sar.vendor, &sar.createdDT, &sar.createdBy, &sar.updatedDT, &sar.updatedBy); err {
		case nil:
			ss.Application.UuID = sar.uuID
			ss.Application.Name = sar.name
			ss.Application.Vendor = sar.vendor
		}
		ssList = append(ssList, ss)
	}
	return ssList, nil
}
