package mysqlclient

import (
	"errors"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Struct to store mysql connection & config(for reconnecting)
type DBM struct {
	mysqlName          string
	DBase              *gorm.DB
	maxOpenConnections int
	maxIdleConnections int
	connectionString   string
}

func GetMysqlClient(
	mysqlMap map[string]string,
) (dbm *DBM, err error) {
	dbm = getNewDBM()
	err = dbm.init(mysqlMap)
	return
}

func (dbm *DBM) ReConnect() error {
	return dbm.initMysqlConnection()
}

func (dbm *DBM) IsConnected() bool {
	if dbm.DBase.DB().Ping() == nil {
		return true
	}
	return false
}

// Close is to close db connection
func (db *DBM) Close() {
	db.DBase.Close()
}

// Private functions

// getNewDBM returns new DBM <very helpful comment, ryt?>
func getNewDBM() *DBM {
	return &DBM{}
}

// InitConnection is to initialize DB connections
func (dbm *DBM) init(
	mysqlMap map[string]string,
) (err error) {
	dbm.mysqlName = mysqlMap["mysqlName"]
	dbm.maxOpenConnections, err = strconv.Atoi(mysqlMap["maxOpenConnections"])
	dbm.maxIdleConnections, err = strconv.Atoi(mysqlMap["maxIdleConnections"])
	dbm.connectionString = getMysqlConnectionString(mysqlMap)

	return dbm.initMysqlConnection()
}

func (dbm *DBM) initMysqlConnection() error {
	db, err := gorm.Open("mysql", dbm.connectionString)

	if err != nil {
		return errors.New("Mysql conn err: " + err.Error())
	} else if db.DB().Ping() != nil {
		return errors.New("Unable to ping DB: " + db.DB().Ping().Error())
	}

	db.DB().SetMaxIdleConns(dbm.maxIdleConnections)
	db.DB().SetMaxOpenConns(dbm.maxOpenConnections)

	dbm.DBase = db
	return nil
}

func getMysqlConnectionString(dsn map[string]string) string {
	//  db, err := gorm.Open("mysql", "user:password@tcp(hostname:port)/mysqlDbName?charset=utf8&parseTime=True&loc=Local")
	conString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=GMT",
		dsn["username"],
		dsn["password"],
		dsn["hostname"],
		dsn["port"],
		dsn["dbname"],
	)
	return conString
}
