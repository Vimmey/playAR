package menu

import (
	"playAR/apis/common/clients/mysqlclient"
)

const (
	Id         = "id"
	Item_name  = "item_name"
	Menu_id    = "restaurant_id"
	Is_enabled = "is_enabled"
	Item_price = "item_price"
)

type Menu struct {
	Id         *int     `gorm:"column:id;primary_key"`
	Item_name  *string  `gorm:"column:item_name"`
	Menu_id    *int     `gorm:"column:restaurant_id"`
	Is_enabled *int     `gorm:"column:is_enabled"`
	Item_price *float64 `gorm:"column:item_price"`
}

func NewMenu() *Menu {
	return &Menu{}
}

//Implement mysql.GenericTable interface
func (*Menu) TableName() string {
	return "restaurant_menu"
}

func GetById(
	conn *mysqlclient.DBM,
	id int,
) (
	*Menu,
	error,
) {
	obj := NewMenu()
	filters := make(map[string]interface{})
	filters["id"] = id
	err := conn.GetFirstByFilter(obj, filters)
	if err == mysqlclient.ErrRecordNotFound {
		return nil, nil
	}
	return obj, err
}
