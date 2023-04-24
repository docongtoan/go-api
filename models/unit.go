package models

import (
	"goserverapi/config/db"

	"github.com/go-playground/validator/v10"
)

type Unit struct {
	Id      int64  `json:"id"`
	Code    string `json:"code" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Status  int    `json:"status" validate:"gte=0,lte=1"`
	MakerId int64  `json:"maker_id" validate:"required"`
}

func ValidateUnit(unit Unit) error {
	validate := validator.New()
	return validate.Struct(unit)
}

func GetListUnit() ([]Unit, error) {

	var result []Unit

	sql := ` SELECT * FROM tb_unit WHERE 1=1 `

	result_rows, err := db.DB.Query(sql)

	for result_rows.Next() {

		var _unit Unit

		if err := result_rows.Scan(&_unit.Id, &_unit.Code, &_unit.Name, &_unit.Status, &_unit.MakerId); err != nil {
			return nil, err
		}

		result = append(result, _unit)

	}

	if err != nil {
		return result, err
	}

	return result, nil

}

func GetRowUnit(id int64) (*Unit, error) {

	sql := ` SELECT * FROM tb_unit WHERE id = $1 `

	result := new(Unit)

	if err := db.DB.QueryRow(sql, id).Scan(&result.Id, &result.Code, &result.Name, &result.Status, &result.MakerId); err != nil {
		return result, err
	}

	return result, nil
}

func CheckExistUnit(id int64, code string) (bool, error) {
	var sql string
	numRows := 0
	if id > 0 {
		sql = ` SELECT count(id) AS count FROM tb_unit WHERE id != $1 AND code = $2`
		if err := db.DB.QueryRow(sql, id, code).Scan(&numRows); err != nil {
			return false, err
		}
	} else {
		sql = ` SELECT count(id) FROM tb_unit WHERE code = $1 `
		if err := db.DB.QueryRow(sql, code).Scan(&numRows); err != nil {
			return false, err
		}
	}

	if numRows > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func InsertUnit(unit Unit) (int, error) {

	sql := ` INSERT INTO tb_unit (code,name,status,maker_id) VALUES ($1,$2,$3,$4) RETURNING id`

	id := 0

	if err := db.DB.QueryRow(sql, unit.Code, unit.Name, unit.Status, unit.MakerId).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func UpdateUnit(unit Unit, id int64) (int, error) {

	sql := ` UPDATE tb_unit SET code = $1 , name = $2, status =$3, maker_id= $4 WHERE id = $5 RETURNING id`

	unit_id := 0

	if err := db.DB.QueryRow(sql, unit.Code, unit.Name, unit.Status, unit.MakerId, id).Scan(&unit_id); err != nil {
		return unit_id, err
	}

	return unit_id, nil
}

func DeleteUnit(id int64) (int, error) {

	sql := ` DELETE FROM tb_unit WHERE id = $1 RETURNING id`

	unit_id := 0

	if err := db.DB.QueryRow(sql, id).Scan(&unit_id); err != nil {
		return unit_id, err
	}

	return unit_id, nil
}
