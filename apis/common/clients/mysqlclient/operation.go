package mysqlclient

import _ "github.com/go-sql-driver/mysql"

func (db *DBM) FirstOrCreate(
	primaryKeyObj GenericTable,
	attributesObj GenericTable,
	obj interface{},
) error {
	return db.DBase.Where(primaryKeyObj).Attrs(attributesObj).FirstOrCreate(obj).Error
	//return db.DBase.Where(primaryKeyObj).Debug().Attrs(attributesObj).Debug().FirstOrCreate(obj).Debug().Error
}
