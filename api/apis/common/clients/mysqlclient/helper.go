package mysqlclient

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Variables for common model
var (
	ErrRecordNotFound = errors.New("RecordNotFound")
)

// Filter is a map for where params in select
type Filter map[string]interface{}

// NewFilter is to get the filter struct
func NewFilter() Filter {
	f := make(Filter)
	return f
}

// Predicate is a map for where params in select
type Predicate map[string]interface{}

// NewPredicate is to get the filter struct
func NewPredicate() Predicate {
	f := make(Predicate)
	return f
}

// GenericTable is an interface for the table
type GenericTable interface {
	TableName() string
}

// GetFirstByFilter is to get the first result using filters
func (db *DBM) GetFirstByFilter(
	table interface{},
	filters Filter,
) (
	err error,
) {
	// return db._getByFilter(table, filters, true, 0, 0)
	return db._getByFilter(table, filters, false, 0, 0)
}

// GetByFilter is to get all the results based on limit and offset
func (db *DBM) GetByFilter(
	table interface{},
	filters Filter,
	limit int,
	offset int,
) (
	err error,
) {
	return db._getByFilter(table, filters, false, limit, offset)
}

// GetByFilter is to get all the results based on limit and offset
func (db *DBM) GetAllByFilter(
	table interface{},
	filters Filter,
	limit int,
	offset int,
) (
	err error,
) {
	return db._getByFilter(table, filters, true, limit, offset)
}

func (db *DBM) _getByFilter(
	table interface{},
	filters Filter,
	singleRow bool,
	limit int,
	offset int,
) (
	err error,
) {
	if len(filters) < 1 {
		return errors.New("No Filter is specified")
	}

	f := map[string]interface{}(filters)
	if singleRow {
		err = db.DBase.Where(f).First(table).Error
	} else {
		err = db.DBase.Where(f).Limit(limit).Offset(offset).Find(table).Error
	}
	return db.handleError(err)
}

// GetByFilter is to get all the results based on limit and offset
func (db *DBM) GetBulk(
	table interface{},
	limit int,
	offset int,
) (
	err error,
) {
	return db._getBulk(table, limit, offset)
}

func (db *DBM) _getBulk(
	table interface{},
	limit int,
	offset int,
) (
	err error,
) {
	err = db.DBase.Find(table).Error
	return db.handleError(err)
}

func (db *DBM) GetByPK(
	table interface{},
	filter map[string]interface{},
) (
	interface{},
	error,
) {
	err := db.DBase.Where(filter).Find(table).Error
	//err := db.DBase.Where(filter).Debug().Find(table).Debug().Error
	return table, err
}

func (db *DBM) handleError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
