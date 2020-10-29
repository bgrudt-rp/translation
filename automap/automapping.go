package automap

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"translation/model"
)

//AutoMapping - main automapping function.
//Passes off to indivudal functions for automapping
//levels.
func AutoMapping(nc *model.ClientCode) error {
	nc.ValidatedFlag = false

	switch {
	case nc.AutoMapInt > 0:
		filePath := "./automap/level_one.json"
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		var mappings model.Mappings

		err = json.Unmarshal(file, &mappings)
		if err != nil {
			return err
		}

		for _, m := range mappings.Mapping {
			if nc.Code == m.ClientCode && nc.CodeType.Description == m.CodeType {
				log.Printf("We found a match")
				nc.StandardCode.Code = m.StandardCode
				break
			}
		}

		qry := `SELECT standard_code_id FROM standard_code WHERE type_id = $1 AND code = $2;`

		row := model.Db.QueryRow(qry, nc.CodeType.ID, nc.StandardCode.Code)
		switch err := row.Scan(&nc.StandardCode.ID); err {
		case nil:
			nc.AutoMapInt = 1
		case sql.ErrNoRows:
			if nc.AutoMapInt == 1 {
				log.Printf("oops you hit\n")
				nc.AutoMapInt = 0
				break
			}
		default:
			return err
		}
		fallthrough

	//Case 1 - compare studies to all existing studies from
	//the same EMR.  If client code maps to standard code  > 95%
	//of the time (verified mappings only), then automap.
	case nc.AutoMapInt > 1:
		qry := `SELECT
					COUNT(*),
					st.standard_code_id,
					st.code,
					st.description
				FROM 
					client_code as cc
				LEFT JOIN source_system as sc
					ON cc.source_system_id  = sc.source_system_id 
				LEFT JOIN source_system_application as sa 
					ON sc.application_id = sa.source_system_application_id 
				LEFT JOIN standard_code as st 
					ON cc.standard_code_id = st.standard_code_id 
				LEFT JOIN code_type as ct 
					ON cc.type_id = ct.code_type_id 
				WHERE ct.description = $1
				AND cc.code = $2
				AND sa.source_system_application_id = $3
				AND cc.validated_flag = true
				GROUP BY st.standard_code_id, st.code, st.description;`

		row, err := model.Db.Query(qry, nc.CodeType.Description, nc.Code, nc.SourceSystem.Application.ID)
		if err != nil {
			return err
		}
		defer row.Close()

		var count int
		var ret []model.StandardCodeByEMR

		for row.Next() {
			var r model.StandardCodeByEMR
			err := row.Scan(&r.Number, &r.ID, &r.Code, &r.Description)
			if err != nil {
				log.Printf("error 2")
				return err
			}
			count = count + r.Number
			ret = append(ret, r)
		}
		hitRate := ret[0].Number / count

		minRate := 0.95
		minNum := 5

		log.Printf("Hitrate: %v\nCount: %v\n", hitRate, count)

		if float64(hitRate) > minRate && count >= minNum {
			nc.StandardCode.ID = ret[0].ID
			nc.StandardCode.Code = ret[0].Code
			nc.StandardCode.Description = ret[0].Description

			nc.ValidatedFlag = true
			nc.AutoMapInt = 2
		}
		fallthrough

	case nc.AutoMapInt > 2:
		fallthrough

	default:
		nc.AutoMapInt = 0
	}
	return nil
}

/*
r2, err := model.Db.Query(qry2, sclr.ID)
		if err != nil {
			return nil, err
		}
		defer r2.Close()

		for r2.Next() {
			var sc model.StandardCodes
			var scr model.StandardCodes

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

*/
