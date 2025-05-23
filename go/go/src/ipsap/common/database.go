package common

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)


var dbconn *sql.DB = nil

func DB_pool_connect() {
	db, err := sql.Open(Config.Database.DriverName, Config.Database.Username+":"+Config.Database.Password+"@tcp("+Config.Database.Url+")/"+Config.Database.Database)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(8) //	갯수 늘림
	if err != nil {
		log.Fatal(err)
	}

	dbconn = db
	//  defer dbconn.Close()  //  닫지 않는다.
}

func DBconn() *sql.DB {
	return dbconn
}

func DB_fetch_one(sql string, filter func(map[string]interface{}), args ...interface{}) map[string]interface{} {
	return db_fetch_one(nil, sql, filter, args...)
}

func DB_Tx_fetch_one(tx *sql.Tx, sql string, filter func(map[string]interface{}), args ...interface{}) map[string]interface{} {
	return db_fetch_one(tx, sql, filter, args...)
}

func DB_fetch_all(sql string, filter func(map[string]interface{}), args ...interface{}) []map[string]interface{} {
	return db_fetch_all(nil, sql, filter, args...)
}

func DB_Tx_fetch_all(tx *sql.Tx, sql string, filter func(map[string]interface{}), args ...interface{}) []map[string]interface{} {
	return db_fetch_all(tx, sql, filter, args...)
}

func db_fetch_one(tx *sql.Tx, sqlStr string, filter func(map[string]interface{}), args ...interface{}) map[string]interface{} {

	// todo test 로컬 개발용으로, 로컬시에는 호출안할수 있도록 처리
	if Config.Test.IsQueryDebug {
		log.Println("[SQL] : ", sqlStr)
		log.Print("[Params] : ")
		log.Println(args...)
	}

	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = DBconn().Query(sqlStr, args...)
	} else {
		rows, err = tx.Query(sqlStr, args...)
	}

	if err != nil {
		log.Println(sqlStr)
		log.Println(args...)
		log.Println(err)
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	if rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val		//	wowdolf : 문자 변환 안함!!! 2021.01.06
//				v = ToStr(val)
			}
			entry[col] = v
		}
		if filter != nil {
			filter(entry)
		}
		return entry
	}
	return nil
}

//func DB_fetch_all(tx *sql.Tx, sqlStr string, filter func(map[string]interface{}), args ...interface{}) []map[string]interface{} {
func db_fetch_all(tx *sql.Tx, sqlStr string, filter func(map[string]interface{}), args ...interface{}) []map[string]interface{} {

	// todo test 로컬 개발용으로, 로컬시에는 호출안할수 있도록 처리
	if Config.Test.IsQueryDebug {
		log.Println("[SQL] : ", sqlStr)
		log.Print("[Params] : ")
		log.Println(args...)
	}

	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = DBconn().Query(sqlStr, args...)
	} else {
		rows, err = tx.Query(sqlStr, args...)
	}

	//rows, err := DBconn().Query(sqlStr, args...)

	if err != nil {
		log.Println(sqlStr)
		log.Println(err)
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val		//	wowdolf : 문자 변환 안함!!! 2021.01.06
//				v = ToStr(val)
			}
			entry[col] = v
		}
		if filter != nil {
			filter(entry)
		}
		tableData = append(tableData, entry)
	}
	return tableData
}

func db_exec(sql string) bool {
	result, err := DBconn().Exec(sql)
	if err != nil {
		log.Println(sql)
		log.Println(err)
		return false
	}

	//  DBconn().Exec("INSERT INTO test1 VALUES($1, $2)", 11, "Jack")
	n, err := result.RowsAffected()
	if n == 1 {

	}
	return true
}

func DB_insert_and_get_lastId(sql string, args ...interface{}) (bool, int64) {
	result, err := DBconn().Exec(sql, args...)
	if err != nil {
		log.Println(sql)
		log.Println(err)
		return false, 0
	}

	n, err := result.LastInsertId()
	return true, n
}

func DBmakeSetString(data map[string]interface{}) string {
	var ret string
	for key, val := range data {
		str := ToStr(key) + `="` + ToStr(val) + `"`
		if ret != "" {
			ret = ret + ","
		}
		ret = ret + str
	}

	return ret
}
