package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"translation/model"
)

//InsertSourceSystem inserts a new source system for use in
//understanding source system data
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
		qry2 = "INSERT INTO source_system (description, created_user, modified_user) VALUES ('" + ss.Description + "', '" + appUser + "', '" + appUser + "');"
	case nil:
		qry2 = "INSERT INTO source_system (application_id, description, created_user, modified_user) VALUES ('" + strconv.Itoa(id) + "', '" + ss.Description + "', '" + appUser + "', '" + appUser + "');"
	default:
		return nil, err
	}

	r, err := db.Exec(qry2)
	if err != nil {
		return nil, err
	}
	return r, nil
}

//SelectSourceSystems returns a list of SourceSystem structs to be used
//to define all current source systems in the environment.
func SelectSourceSystems() ([]model.SourceSystem, error) {
	var ssList []model.SourceSystem
	var ss model.SourceSystem
	var sar model.SourceSystemApplication
	var ssr model.SourceSystem

	qry := "SELECT * FROM source_system;"
	fmt.Println(qry)
	r, err := db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		err := r.Scan(&ssr.ID, &ssr.UuID, &ssr.Application.ID, &ssr.Description, &ssr.Metadata.CreatedDT, &ssr.Metadata.CreatedBy, &ssr.Metadata.ModifiedDT, &ssr.Metadata.ModifiedBy)
		if err != nil {
			return nil, err
		}
		ss.UuID = ssr.UuID
		ss.Description = ssr.Description
		ss.Metadata.CreatedBy = ssr.Metadata.CreatedBy
		ss.Metadata.CreatedDT = ssr.Metadata.CreatedDT
		ss.Metadata.ModifiedBy = ssr.Metadata.ModifiedBy
		ss.Metadata.ModifiedDT = ssr.Metadata.ModifiedDT

		qry := `SELECT * FROM source_system_application WHERE source_system_application_id = $1;`

		row := db.QueryRow(qry, ssr.Application.ID)

		switch err := row.Scan(&sar.ID, &sar.UuID, &sar.Name, &sar.Vendor, &sar.Metadata.CreatedDT, &sar.Metadata.CreatedBy, &sar.Metadata.ModifiedDT, &sar.Metadata.ModifiedBy); err {
		case nil:
			ss.Application.UuID = sar.UuID
			ss.Application.Name = sar.Name
			ss.Application.Vendor = sar.Vendor

			ss.Application.Metadata.CreatedBy = sar.Metadata.CreatedBy
			ss.Application.Metadata.CreatedDT = sar.Metadata.CreatedDT
			ss.Application.Metadata.ModifiedBy = sar.Metadata.ModifiedBy
			ss.Application.Metadata.ModifiedDT = sar.Metadata.ModifiedDT

		}
		ssList = append(ssList, ss)
	}
	return ssList, nil
}
