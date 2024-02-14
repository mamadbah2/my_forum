package models

type Category struct {
	Category_id int
	Name        string
}

func (m *ConnDB) GetCategory(id int) (*Category, error) {
	statement := `SELECT * FROM Category WHERE category_id = ?`
	row := m.DB.QueryRow(statement, id)
	c := &Category{}
	err := row.Scan(&c.Category_id, &c.Name)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (m *ConnDB) SetCategory(name string) (int, error) {
	statement := `INSERT INTO Category(name) VALUES (?)`
	result, err := m.DB.Exec(statement, name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *ConnDB) GetAllCategory() ([]*Category, error) {
	statement := `SELECT * FROM Category`
	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []*Category
	for rows.Next() {
		c := &Category{}
		rows.Scan(&c.Category_id, &c.Name)
		categories = append(categories, c)
	}
	return categories, nil
}
