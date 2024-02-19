package models

import (
	"time"
)

type Comment struct {
	Comment_id    int
	Comment       string
	Date_Creation time.Time
	Post_id       int
	User_id       int
}

type CommentInfo struct {
	Comment_id    int
	Comment       string
	Date_Creation string
	Username      string
}

func (m *ConnDB) getCommentByPost(postId int) ([]*Comment, error) {
	statement := `SELECT * FROM Comment WHERE post_id = ? 
	ORDER BY comment_id DESC`
	rows, err := m.DB.Query(statement, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []*Comment
	for rows.Next() {
		c := &Comment{}
		err := rows.Scan(&c.Comment_id, &c.Comment, &c.Date_Creation, &c.Post_id, &c.User_id)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func (m *ConnDB) getCommentNumberByPost(postId int) (int, error) {
	statement := `SELECT COUNT(*) FROM Comment WHERE post_id = ?`
	row := m.DB.QueryRow(statement, postId)
	var nbreComment int
	err := row.Scan(&nbreComment)
	if err != nil {
		return 0, err
	}
	return nbreComment, nil
}

func (m *ConnDB) GetCommentNumberByUser(userId int) (int, error) {
	statement := `SELECT COUNT(*) FROM Comment WHERE user_id = ?`
	row := m.DB.QueryRow(statement, userId)
	var nbreComment int
	err := row.Scan(&nbreComment)
	if err != nil {
		return 0, err
	}
	return nbreComment, nil
}

func (m *ConnDB) SetComment(comment string, postId, userId int) (int, error) {
	statement := `INSERT INTO Comment (comment, date_creation, post_id, user_id)
	VALUES (?, CURRENT_TIMESTAMP, ?, ?)`
	result, err := m.DB.Exec(statement, comment, postId, userId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *ConnDB) GetCommentsInfoByPost(postId int) ([]*CommentInfo, error) {
	comments, err := m.getCommentByPost(postId)
	if err != nil {
		return nil, err
	}
	var commentsInfo []*CommentInfo

	for _, comment := range comments {
		commentInfo := &CommentInfo{}
		commentInfo.Comment_id = comment.Comment_id
		commentInfo.Comment = comment.Comment
		commentInfo.Date_Creation = comment.Date_Creation.String()
		user, err := m.GetUser(comment.User_id)
		if err != nil {
			return nil, err
		}
		commentInfo.Username = user.Username

		commentsInfo = append(commentsInfo, commentInfo)
	}

	return commentsInfo, nil
}
