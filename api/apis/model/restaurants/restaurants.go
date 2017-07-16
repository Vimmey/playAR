package restaurants

import (
	"crypto/rand"
	"math/big"
	"playAR/apis/common/clients/mysqlclient"
)

const (
	Id         = "id"
	Name       = "name"
	Cuisine    = "cuisine"
	Area_id    = "area_id"
	City_id    = "city_id"
	Lat        = "lat"
	Lng        = "lng"
	Avg_rating = "avg_rating"
)

type restaurants struct {
	Id         *int     `gorm:"column:id;primary_key"`
	Name       *string  `gorm:"column:name"`
	Cuisine    *string  `gorm:"column:cuisine"`
	Area_id    *int     `gorm:"column:area_id"`
	City_id    *int     `gorm:"column:city_id"`
	Lat        *string  `gorm:"column:lat"`
	Lng        *string  `gorm:"column:lng"`
	Avg_rating *float64 `gorm:"column:avg_rating"`
}

func NewRestaurant() *restaurants {
	return &restaurants{}
}

//Implement mysql.GenericTable interface
func (*restaurants) TableName() string {
	return "restaurants"
}

func GetById(
	conn *mysqlclient.DBM,
	id int,
) (
	*restaurants,
	error,
) {
	obj := NewRestaurant()
	filters := make(map[string]interface{})
	filters["id"] = id
	err := conn.GetFirstByFilter(obj, filters)
	if err == mysqlclient.ErrRecordNotFound {
		return nil, nil
	}
	return obj, err
}

func GetByLatLong(
	conn *mysqlclient.DBM,
	lat string,
	long string,
) (
	[]*restaurants,
	error,
) {
	var obj []*restaurants
	limit := 20

	// create random offset
	nBig, err := rand.Int(rand.Reader, big.NewInt(800))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	offset := int(n)
	err = conn.GetBulk(obj, limit, offset)
	if err == mysqlclient.ErrRecordNotFound {
		return nil, nil
	}
	return obj, err
}
