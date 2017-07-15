package connections

import (
	"playAR/apis/common/clients/mysqlclient"
)

type MysqlConn struct {
	Dbm                 *mysqlclient.DBM
	mysql_name          string
	hostname            string
	dbname              string
	port                int
	username            string
	password            string
	max_open_connection int
	max_idle_connection int
}

func GetNewMysqlConn() *MysqlConn {
	return &MysqlConn{}
}

func (mcs *MysqlConn) InitMysqlClient() (err error) {
	// todo : Read from config
	mysqlMap := make(map[string]string)
	mysqlMap["mysql_name"] = "swiggy_db"
	mysqlMap["hostname"] = "localhost"
	mysqlMap["dbname"] = "swiggy"
	mysqlMap["port"] = "3306"
	mysqlMap["username"] = "root"
	mysqlMap["password"] = "Playaradmin@123"
	mysqlMap["max_open_connection"] = "150"
	mysqlMap["max_idle_connection"] = "5"

	mcs.Dbm, err = mysqlclient.GetMysqlClient(mysqlMap)
	if err == nil {
		return
	}
	return
}
