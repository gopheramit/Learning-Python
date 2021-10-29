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

func (m *UserModel) GetTaskByID(id int) (*models.Tasks, error) {

	stmt := `SELECT TaskID,TaskName,TaskDescription,Difficulty from TASKS where TaskID =?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Tasks{}
	err := row.Scan(&s.TaskID, &s.TaskName, &s.TaskDescription, &s.Difficulty)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *UserModel) IncrementTaskId(taskId int, userId string) (int, error) {
	taskId = taskId + 1

	upd, err := m.DB.Prepare(`update PYTHONUSER set TaskID=? where userID =?`)

	if err != nil {
		panic(err.Error())
	}
	upd.Exec(taskId, userId)
	// log.Println("UPDATE: Name: " + name + " | City: " + city)
	// _, err := m.DB.Exec(stmt, taskId, userId)
	// if err != nil {
	// 	fmt.Println("error in executing db.exec" + err.Error())
	// 	return 0, err
	// }
	return taskId, nil

}
