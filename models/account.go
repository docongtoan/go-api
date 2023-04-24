package models

import (
	"goserverapi/config/db"

	"github.com/go-playground/validator/v10"
)

type Account struct {
	Id      int64  `json:"id"`
	Code    string `json:"code" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Status  int    `json:"status" validate:"gte=0,lte=1"`
	MakerId int64  `json:"maker_id" validate:"required"`
}

func ValidateAccount(account Account) error {
	validate := validator.New()
	return validate.Struct(account)
}

func GetListAccount() ([]Account, error) {

	var result []Account

	sql := ` SELECT * FROM tb_accounts WHERE 1=1 `

	result_rows, err := db.DB.Query(sql)

	for result_rows.Next() {

		var _account Account

		if err := result_rows.Scan(&_account.Id, &_account.Code, &_account.Name, &_account.Status, &_account.MakerId); err != nil {
			return nil, err
		}

		result = append(result, _account)

	}

	if err != nil {
		return result, err
	}

	return result, nil

}

func GetRowAccount(id int64) (*Account, error) {

	sql := ` SELECT * FROM tb_accounts WHERE id = $1 `

	result := new(Account)

	if err := db.DB.QueryRow(sql, id).Scan(&result.Id, &result.Code, &result.Name, &result.Status, &result.MakerId); err != nil {
		return result, err
	}

	return result, nil
}

func CheckExistAccount(id int64, code string) (bool, error) {
	var sql string
	numRows := 0
	if id > 0 {
		sql = ` SELECT count(id) AS count FROM tb_accounts WHERE id != $1 AND code = $2`
		if err := db.DB.QueryRow(sql, id, code).Scan(&numRows); err != nil {
			return false, err
		}
	} else {
		sql = ` SELECT count(id) FROM tb_accounts WHERE code = $1 `
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

func InsertAccount(account Account) (int, error) {

	sql := ` INSERT INTO tb_accounts (code,name,status,maker_id) VALUES ($1,$2,$3,$4) RETURNING id`

	id := 0

	if err := db.DB.QueryRow(sql, account.Code, account.Name, account.Status, account.MakerId).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func UpdateAccount(account Account, id int64) (int, error) {

	sql := ` UPDATE tb_accounts SET code = $1 , name = $2, status =$3, maker_id= $4 WHERE id = $5 RETURNING id`

	account_id := 0

	if err := db.DB.QueryRow(sql, account.Code, account.Name, account.Status, account.MakerId, id).Scan(&account_id); err != nil {
		return account_id, err
	}

	return account_id, nil
}

func DeleteAccount(id int64) (int, error) {

	sql := ` DELETE FROM tb_accounts WHERE id = $1 RETURNING id`

	account_id := 0

	if err := db.DB.QueryRow(sql, id).Scan(&account_id); err != nil {
		return account_id, err
	}

	return account_id, nil
}
