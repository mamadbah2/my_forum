package models

type PostCategory struct {
	Post_category_id int
	Post_id          int
	Category_id      int
}

func (m *ConnDB) SetPostCategory(postId, categoryId int) (int, error) {
	statement := `INSERT INTO PostCategory (post_id, category_id) VALUES (?, ?)`
	result, err := m.DB.Exec(statement, postId, categoryId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *ConnDB) GetCategoriesByPost(post_id int) ([]*Category, error) {
	statement := `SELECT category_id FROM PostCategory WHERE post_id = ?`
	rows, err := m.DB.Query(statement, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []*Category{}
	for rows.Next() {
		pc := &PostCategory{}
		err = rows.Scan(&pc.Category_id)
		if err != nil {
			return nil, err
		}

		c, err := m.GetCategory(pc.Category_id)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}
