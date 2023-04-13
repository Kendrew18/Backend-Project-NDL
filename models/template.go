package models

import (
	"Backend-Project-NDL/db"
	"net/http"
)

func Update_Template(internal_instruction_number string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlstatement := "UPDATE template SET internal_instruction_number=? WHERE ws_no=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(internal_instruction_number)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}
