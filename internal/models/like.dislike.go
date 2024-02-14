package models

import (
	"database/sql"
	"errors"
)

type LikeDislike struct {
	LikeDislike_id int
	User_id        int
	Post_id        int
	Liked          bool
	Disliked       bool
}

func (m *ConnDB) getLikeDislike(likeDislikeId int) (*LikeDislike, error) {
	statement := `SELECT * FROM LikeDislike WHERE like_dislike_id=?`
	row := m.DB.QueryRow(statement, likeDislikeId)
	ld := &LikeDislike{}
	err := row.Scan(&ld.LikeDislike_id, &ld.User_id, &ld.Post_id, &ld.Liked, &ld.Disliked)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return ld, nil
}

func (m *ConnDB) getLikeDislikePU(user_id int, post_id int) (*LikeDislike, error) {
	statement := `SELECT like_dislike_id, liked, disliked FROM LikeDislike WHERE user_id=? AND post_id=?`
	row := m.DB.QueryRow(statement, user_id, post_id)
	ld := &LikeDislike{}
	err := row.Scan(&ld.LikeDislike_id, &ld.Liked, &ld.Disliked)
	if err != nil {
		return nil, err
	}
	return ld, nil
}

func (m *ConnDB) SetLike(userId int, postId int, liked bool) (int, error) {
	_, err := m.getLikeDislikePU(userId, postId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		statement := `INSERT INTO LikeDislike (user_id, post_id, liked, disliked) VALUES (?, ?, TRUE, FALSE)`

		result, err := m.DB.Exec(statement, userId, postId)
		if err != nil {
			return 0, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return int(id), nil
	}
	var statement string
	if !liked {
		statement = `UPDATE LikeDislike SET liked = TRUE, disliked = FALSE WHERE user_id=? AND post_id=?`
	} else {
		statement = `UPDATE LikeDislike SET liked = FALSE WHERE user_id=? AND post_id=?`
	}
	result, err := m.DB.Exec(statement, userId, postId)
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *ConnDB) SetDislike(userId int, postId int, disliked bool) (int, error) {
	_, err := m.getLikeDislikePU(userId, postId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		statement := `INSERT INTO LikeDislike (user_id, post_id, liked, disliked) VALUES (?, ?, FALSE, TRUE)`

		result, err := m.DB.Exec(statement, userId, postId)
		if err != nil {
			return 0, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return int(id), nil
	}
	var statement string
	if !disliked {
		statement = `UPDATE LikeDislike SET liked = FALSE, disliked = TRUE WHERE user_id=? AND post_id=?`
	} else {
		statement = `UPDATE LikeDislike SET disliked = FALSE WHERE user_id=? AND post_id=?`
	}
	result, err := m.DB.Exec(statement, userId, postId)
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *ConnDB) getLikeNumberByPost(post_id int) (int, error) {
	statement := `SELECT COUNT(*) FROM LikeDislike WHERE liked=TRUE AND post_id= ?`
	row := m.DB.QueryRow(statement, post_id)
	var counter int
	err := row.Scan(&counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (m *ConnDB) getDislikeNumberByPost(post_id int) (int, error) {
	statement := `SELECT COUNT(*) FROM LikeDislike WHERE disliked=TRUE AND post_id= ?`
	row := m.DB.QueryRow(statement, post_id)
	var counter int
	err := row.Scan(&counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}
