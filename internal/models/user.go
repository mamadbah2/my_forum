package models

type User struct {
	User_id  int
	Username string
	Email    string
	Password string
}

func (m *ConnDB) getUser(id int) (*User, error) {
	statement := `SELECT * FROM User WHERE user_id = ?`
	row := m.DB.QueryRow(statement, id)
	user := &User{}
	row.Scan(&user.User_id, &user.Username, &user.Email, &user.Password)
	return user, nil
}

func (m *ConnDB) setUser(username string, email string, password string) (int, error) {

	return 0, nil
}

func (m *ConnDB) getAllUser() ([]*User, error) {

	return nil, nil
}
