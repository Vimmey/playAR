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
	Id            *int     `gorm:"column:id;primary_key"`
	Item_name     *string  `gorm:"column:item_name"`
	Restaurant_id *int     `gorm:"column:restaurant_id"`
	Is_enabled    *int     `gorm:"column:is_enabled"`
	Item_price    *float64 `gorm:"column:item_price"`
}

type MenuRes struct {
	Item_name  *string `json:"Item_name"`
	Item_price *string `json:"Item_price"`
}

func NewMenu() *Menu {
	return &Menu{}
}

func NewMenuRes(
	s *string,
	f string,
) *MenuRes {
	return &MenuRes{
		Item_name:  s,
		Item_price: &f,
	}
}

//Implement mysql.GenericTable interface
func (*Menu) TableName() string {
	return "restaurant_menu"
}

func GetById(
	conn *mysqlclient.DBM,
	id int,
) (
	objs []Menu,
	err error,
) {
	objs = []Menu{}

	for c := 1; c <= 903; c++ {
		obj := NewMenu()
		filters := make(map[string]interface{})
		filters["restaurant_id"] = id
		// log.Println("line 47; c = ", c)

		err = conn.GetByFilter(&obj, filters, 20, c)
		// spew.Dump(obj)

		if err == mysqlclient.ErrRecordNotFound {
			// return nil, nil
		}
		if obj.Id != nil {
			objs = append(objs, *obj)
		}

	}
	// spew.Dump(objs)
	return objs, err
}
