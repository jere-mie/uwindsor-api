package models

type Building struct {
	Building_Name string  `db:"building_name" json:"building_name"`
	Latitude      float64 `db:"latitude" json:"latitude"`
	Longitude     float64 `db:"longitude" json:"longitude"`
}
