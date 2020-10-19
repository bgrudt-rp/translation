package postgres

import "database/sql"

//Dbbuild will be refined at some point
//We are just trying to form the DB structure for a translation service
//to repattern later
const Dbbuild = `

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS source_system_application (
	source_system_application_id serial PRIMARY KEY,
	source_system_application_uuid uuid DEFAULT uuid_generate_v4 (),
	application_name varchar(50) NOT NULL,
	vendor_name varchar(50),
	created_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	created_user varchar(50) NOT NULL,
	modified_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	modified_user varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS source_system ( 
	source_system_id serial PRIMARY KEY, 
	source_system_uuid uuid DEFAULT uuid_generate_v4 (),
	application_id integer REFERENCES source_system_application(source_system_application_id),
	description varchar(50) NOT NULL,
	created_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	created_user varchar(50) NOT NULL,
	modified_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	modified_user varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS code_type (
	code_type_id serial PRIMARY KEY,
	code_type_uuid uuid DEFAULT uuid_generate_v4 (),
	description varchar(50) NOT NULL,
	created_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	created_user varchar(50) NOT NULL,
	modified_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	modified_user varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS standard_code (
	standard_code_id serial PRIMARY KEY,
	standard_code_uuid uuid DEFAULT uuid_generate_v4 (),
	type_id integer REFERENCES code_type(code_type_id),
	code varchar(50) NOT NULL,
	description varchar(50) NOT NULL,
	created_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	created_user varchar(50) NOT NULL,
	modified_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	modified_user varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS client_code (
	client_code_id serial PRIMARY KEY,
	client_code_uuid uuid DEFAULT uuid_generate_v4 (),
	type_id integer REFERENCES code_type(code_type_id) NOT NULL,
	source_system_id integer REFERENCES source_system(source_system_id) NOT NULL,
	standard_code_id integer REFERENCES standard_code(standard_code_id),
	code varchar(255) NOT NULL,
	description varchar(255),
	automap_int int DEFAULT 0,
	validated_flag boolean DEFAULT FALSE, 
	primary_mapping_flag boolean DEFAULT FALSE,
	created_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	created_user varchar(50) NOT NULL,
	modified_datetime timestamp with time zone DEFAULT current_timestamp NOT NULL,
	modified_user varchar(50) NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS ind_uq_appname ON source_system_application (
	application_name
);

CREATE UNIQUE INDEX IF NOT EXISTS ind_uq_source ON source_system (
	description
);

CREATE UNIQUE INDEX IF NOT EXISTS ind_uq_client_code ON client_code (
 	type_id,
 	source_system_id,
 	code
);

CREATE UNIQUE INDEX IF NOT EXISTS ind_uq_standard_code ON standard_code (
	type_id,
	code
);
`

//BuildDB will call a long string of SQL statements to
//populate a blank PG database.
func BuildDB() error {
	db, err := sql.Open("postgres", Dbconn)
	if err != nil {
		return err
	}

	_, err = db.Exec(Dbbuild)
	if err != nil {
		return err
	}
	return nil
}
