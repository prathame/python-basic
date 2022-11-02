package main

import (
	"database/sql"
	"db_conn/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	var db sql.DB
	valueType := "Statement"
	updatedDoc := []models.Document{}
	updatedState := []models.Statement{}
	_ = updatedState
	db = *dbConnection()
	updatedDoc, updatedState = readCsv(valueType)
	resultDb := checkDb(&db)
	if valueType == "Document" {
		resulttb, db := checkTb(os.Getenv("TBNAME_1"), valueType)
		if resultDb {
			if resulttb {
				insertData(&updatedDoc, &updatedState, db, valueType, os.Getenv("TBNAME_1"))
			} else {
				log.Fatal("Table not found")
			}
		} else {
			log.Fatal("Database not found")
		}

	} else if valueType == "Statement" {
		resulttb, db := checkTb(os.Getenv("TBNAME_2"), valueType)
		if resultDb {
			if resulttb {
				insertData(&updatedDoc, &updatedState, db, valueType, os.Getenv("TBNAME_2"))
			} else {
				log.Fatal("Table not found")
			}
		} else {
			log.Fatal("Database not found")
		}

	}

}
func dbConnection() *sql.DB {
	//fmt.Println(os.Getenv("HOST"))
	// Database Connection string
	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("PORT"))
	// opening connection
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to Postgres")
	}
	return db
}
func dbUpdate() *sql.DB {
	// Database Connection string
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"))
	// opening connection
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Connected to %s\n", os.Getenv("DBNAME"))
	}
	return db
}

func readCsv(valueType string) ([]models.Document, []models.Statement) {
	doc := []models.Document{}
	statement := []models.Statement{}
	if valueType == "Document" {
		file, err := os.Open(os.Getenv("DOCUMENT"))
		if err != nil {
			log.Fatal(err)
		} else {
			df := csv.NewReader(file)
			data, _ := df.ReadAll()
			for _, value := range data {
				//age, _ := strconv.Atoi(value[1])
				//models.Info = append(models.Info, models.Svalue{Name: value[0], Age: age})
				doc = append(doc, models.Document{Document_id: value[0], Document_path: value[1], Document_type: value[2],
					Name: value[3], Version: value[4], Publish_dt: value[5], Effective_dt: value[6], Review_dt: value[7],
					Publisher: value[8], Reviewe_name: value[9], Domination_of_control_geography: value[10], Domination_of_control_bussiness: value[11],
					Domination_of_control_technology_func: value[12], Scope_deploy_id: value[13], Scope_technique_id: value[14],
					Appli_deploy_id: value[15], Appli_technique_id: value[16]})
				//fmt.Println(info)
			}

		}
		return doc, nil

	} else if valueType == "Statement" {
		file, err := os.Open(os.Getenv("STATEMENT"))
		if err != nil {
			log.Fatal(err)
		} else {
			df := csv.NewReader(file)
			data, _ := df.ReadAll()
			for _, value := range data {
				//age, _ := strconv.Atoi(value[1])
				//models.Info = append(models.Info, models.Svalue{Name: value[0], Age: age})
				statement = append(statement, models.Statement{Statement_type: value[0], Document_type: value[1], Document_section_level_1: value[2],
					Document_section_level_2: value[3], Statement_parent_id: value[4], Statement_id: value[5], Statement_words: value[6], Fulfilled_practice: value[7],
					Prac_id: value[8], Act_id: value[9], Scope_technique_id: value[10], Scope_deploy_id: value[11],
					Appli_technique_id: value[12], Appli_deploy_id: value[13]})
				//fmt.Println(info)
			}

		}
		return nil, statement
	} else {
		return nil, nil
	}

}

func checkDb(db *sql.DB) bool {
	//sqlStatement := `select exists (SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower(''));`
	result := ""
	e := db.QueryRow("SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower($1)", os.Getenv("DBNAME")).Scan(&result)
	if e != nil {
		//fmt.Println(e)
		fmt.Println("Creating database")
		_, e1 := db.Exec("create database " + os.Getenv("DBNAME"))
		if e1 != nil {
			log.Fatal(e1)
			return false
		} else {
			result = os.Getenv("DBNAME")
			return true
		}
	}
	if result == os.Getenv("DBNAME") {
		return true
	} else {
		return false
	}

}
func checkTb(tbname string, value_type string) (bool, sql.DB) {
	//var db sql.DB
	db := *dbUpdate()
	if value_type == "Document" {
		_, table_check := db.Query("select * from " + tbname + ";")
		if table_check == nil {
			fmt.Println("Table is Present ")
			return true, db
		} else {
			fmt.Println("Table not found creating table ")
			_, err := db.Exec("CREATE TABLE " + tbname + "( document_id varchar(255), document_path varchar(255), document_type varchar(255), name varchar(255), version varchar(255),publish_dt varchar(255), effective_dt varchar(255), review_dt varchar(255), publisher varchar(255),  reviewe_name varchar(255),domination_of_control_geography varchar(255), domination_of_control_bussiness varchar(255), domination_of_control_technology_func varchar(255), scope_deploy_id varchar(255),scope_technique_id varchar(255), appli_deploy_id varchar(255), appli_technique_id varchar(255))")
			if err != nil {
				log.Fatal(err)
				return false, db
			} else {
				return true, db
			}

		}
	} else if value_type == "Statement" {
		_, table_check := db.Query("select * from " + tbname + ";")
		if table_check == nil {
			fmt.Println("Table is Present ")
			return true, db
		} else {
			fmt.Println("Table not found creating table ")
			_, err := db.Exec("CREATE TABLE " + tbname + "(statement_type  varchar(255), document_type varchar(255), document_section_level_1 varchar(255), document_section_level_2 varchar(255), statement_parent_id varchar(255),statement_id varchar(255), statement_words varchar(255), fulfilled_practice varchar(255), prac_id varchar(255),  act_id varchar(255),scope_technique_id varchar(255), scope_deploy_id varchar(255),appli_technique_id varchar(255), appli_deploy_id varchar(255))")
			if err != nil {
				log.Fatal(err)
				return false, db
			} else {
				return true, db
			}

		}
	} else {
		return false, db
	}

}

func insertData(recievedDoc *[]models.Document, recievedStatement *[]models.Statement, db sql.DB, value_type string, tbname string) {
	var err error
	doc := *recievedDoc
	state := *recievedStatement
	if value_type == "Document" {
		for i := 1; i < len(doc); i++ {
			sqlStatement := `insert into ` + tbname + `(document_id, document_path, document_type, name, version, publish_dt, effective_dt, review_dt, publisher, reviewe_name, domination_of_control_geography, domination_of_control_bussiness,domination_of_control_technology_func, scope_deploy_id, scope_technique_id, appli_deploy_id, appli_technique_id) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`
			_, err := db.Exec(sqlStatement, doc[i].Document_id, doc[i].Document_path, doc[i].Document_type, doc[i].Name, doc[i].Version, doc[i].Publish_dt, doc[i].Effective_dt, doc[i].Review_dt, doc[i].Publisher, doc[i].Reviewe_name, doc[i].Domination_of_control_geography, doc[i].Domination_of_control_bussiness, doc[i].Domination_of_control_technology_func, doc[i].Scope_deploy_id, doc[i].Scope_technique_id, doc[i].Appli_deploy_id, doc[i].Appli_technique_id)
			if err != nil {
				log.Fatal(err)
			}
		}
		if err == nil {
			fmt.Println("Data Inserted")
			db.Close()
		}

	} else if value_type == "Statement" {
		fmt.Println(len(state))
		for i := 1; i < len(state); i++ {
			sqlStatement := `insert into ` + tbname + `(statement_type, document_type, document_section_level_1,Document_section_level_2, statement_parent_id, statement_id , statement_words,fulfilled_practice,prac_id,act_id,scope_technique_id,scope_deploy_id,appli_technique_id,appli_deploy_id) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`
			_, err := db.Exec(sqlStatement, state[i].Statement_type, state[i].Document_type, state[i].Document_section_level_1, state[i].Document_section_level_2, state[i].Statement_parent_id, state[i].Statement_id, state[i].Statement_words, state[i].Fulfilled_practice, state[i].Prac_id, state[i].Act_id, state[i].Scope_technique_id, state[i].Scope_deploy_id, state[i].Appli_technique_id, state[i].Appli_deploy_id)
			if err != nil {
				log.Fatal(err)
			}
		}
		if err == nil {
			fmt.Println("Data Inserted")
			db.Close()
		}
	} else {
		fmt.Println("Sonething went wrong")
		db.Close()
	}

}
