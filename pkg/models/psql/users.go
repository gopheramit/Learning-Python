package psql

import (
	"database/sql"
	"errors"

	// "errors"
	"fmt"

	"github.com/gopheramit/Learning-Python/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(userid string, email string, taskid int) (int, error) {

	stmt := `INSERT INTO pythonuser values($1,$2,$3);`
	result, err := m.DB.Exec(stmt, userid, email, taskid)
	if err != nil {
		// var mySQLError *psql.MySQLError
		// if errors.As(err, &mySQLError) {
		// 	if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
		// 		return 0, models.ErrDuplicateEmail
		// 	}
		// }
		fmt.Println("error in executing db.exec" + err.Error())
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

func (m *UserModel) GetID(id string) (*models.PythonUser, error) {

	stmt := `SELECT userid,email,taskid from pythonuser where userid = $1`
	row := m.DB.QueryRow(stmt, id)
	s := &models.PythonUser{}
	err := row.Scan(&s.TaskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *UserModel) Latest() ([]*models.PythonUser, error) { return nil, nil }
