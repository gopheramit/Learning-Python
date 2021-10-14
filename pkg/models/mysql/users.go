package mysql

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

	stmt := `INSERT INTO PYTHONUSER values(?,?,?);`
	result, err := m.DB.Exec(stmt, userid, email, taskid)
	if err != nil {
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

	stmt := `SELECT userid,email,taskid from  PYTHONUSER where userid =?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.PythonUser{}
	err := row.Scan(&s.UserID, &s.Email, &s.TaskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}
