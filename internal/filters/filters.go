package filters

import (
	"forum.01/internal/models"
)

func CategoryFilter(postsInfo []*models.PostInfo, categories ...string) []*models.PostInfo {
	validPost := []*models.PostInfo{}
	for _, c := range categories {
		for _, pi := range postsInfo {
			for _, categoriesPI := range pi.Categories {
				if c == categoriesPI && !doublonPI(validPost, pi) {
					validPost = append(validPost, pi)
				}
			}
		}
	}
	return validPost
}

func LikedPostFilter(postsInfo []*models.PostInfo) []*models.PostInfo {
	validPost := []*models.PostInfo{}
	for _, pi := range postsInfo {
		if pi.LikeActualUser && !doublonPI(validPost, pi) {
			validPost = append(validPost, pi)
		}
	}
	return validPost
}

func CreatedPostFilter(postsInfo []*models.PostInfo, actualUserId int) []*models.PostInfo {
	validPost := []*models.PostInfo{}
	for _, pi := range postsInfo {
		if pi.User_id == actualUserId && !doublonPI(validPost, pi) {
			validPost = append(validPost, pi)
		}
	}
	return validPost
}

func doublonPI(postsInfo []*models.PostInfo, pi *models.PostInfo) bool {
	for _, postInfo := range postsInfo {
		if postInfo.Post_id == pi.Post_id {
			return true
		}
	}
	return false
}
