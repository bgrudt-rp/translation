package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"translation/model"
)

//GetStandardCode returns a standard code with the corresponding
//client code.
func GetStandardCode(cc *model.ClientCode) error {
	//If a standard code was provided, try to return it.
	//If no standard code was sent, or if the standard code
	//does not return a valid standard code, check the automapping
	//level:
	//
	//Automapping Level 0:
	//Do not automap (default)
	//
	//Automapping Level 1:
	//Only automap explicitly (based on json file)
	//
	//Automapping Level 2:
	//String automapping (TBD)
	//
	//Automapping Level 3:
	//Expanded automapping (throw the kitchen sink at the code to map)

	qry := `SELECT standard_code_id FROM standard_code WHERE type_id = $1 AND code = $2;`

	row := db.QueryRow(qry, cc.CodeType.ID, cc.StandardCode.Code)
	switch err := row.Scan(&cc.StandardCode.ID); err {
	case sql.ErrNoRows:
		//Will add a nested switch to cycle through automap
		//levels and stop if/when a value is mapped.
	case nil:
		cc.ValidatedFlag = true
		cc.AutoMapInt = 0
	default:
		return err
	}

	return nil
}

//InsertClientCode adds a new value to the client code
//table, along with any standard mapping values associated
func InsertClientCode(cc model.ClientCode) (sql.Result, error) {
	//Perform easy eval for bad data
	if len(cc.Code) > 255 {
		return nil, fmt.Errorf("client code too long.  do not exceed 255 characters")
	}
	if len(cc.Description) > 255 {
		return nil, fmt.Errorf("client code description too long.  do not exceed 255 characters")
	}

	//Lookup required primary keys for inclusion
	qry := `SELECT source_system_id FROM source_system WHERE description = $1;`
	row := db.QueryRow(qry, cc.SourceSystem.Description)
	switch err := row.Scan(&cc.SourceSystem.ID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("source system not found.  please verify")
	case nil:
	default:
		return nil, fmt.Errorf("error finding source system.  please verify")
	}

	qry = `SELECT code_type_id FROM code_type WHERE description = $1;`
	row = db.QueryRow(qry, cc.CodeType.Description)
	switch err := row.Scan(&cc.CodeType.ID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("code type not found.  please verify")
	case nil:
	default:
		return nil, fmt.Errorf("error finding code type.  please verify")
	}

	//Run code through autocoding
	if len(cc.StandardCode.Code) > 0 || cc.AutoMapInt > 0 {
		log.Printf(cc.StandardCode.Code)
		err := GetStandardCode(&cc)
		log.Printf(cc.StandardCode.ID)
		if err != nil {
			return nil, err
		}
	}

	//Insert new client code
	if len(cc.StandardCode.ID) > 0 {
		qry = "INSERT INTO client_code" +
			" (type_id, source_system_id, standard_code_id, code, description, automap_int, validated_flag, primary_mapping_flag, created_user, modified_user)" +
			" VALUES " +
			"('" + cc.CodeType.ID + "', '" + cc.SourceSystem.ID + "', '" + cc.StandardCode.ID + "', '" + cc.Code + "', '" + cc.Description + "', '" +
			strconv.Itoa(cc.AutoMapInt) + "', " + strconv.FormatBool(cc.ValidatedFlag) + ", " +
			strconv.FormatBool(cc.PrimaryMappingFlag) + ", '" + appUser + "', '" + appUser + "');"
	} else {
		qry = "INSERT INTO client_code" +
			" (type_id, source_system_id, code, description, automap_int, validated_flag, primary_mapping_flag, created_user, modified_user)" +
			" VALUES " +
			"('" + cc.CodeType.ID + "', '" + cc.SourceSystem.ID + "', '" + cc.Code + "', '" + cc.Description + "', '" +
			strconv.Itoa(cc.AutoMapInt) + "', " + strconv.FormatBool(cc.ValidatedFlag) + ", " +
			strconv.FormatBool(cc.PrimaryMappingFlag) + ", '" + appUser + "', '" + appUser + "');"
	}

	r, err := db.Exec(qry)
	if err != nil {
		return nil, err
	}

	return r, nil
}
