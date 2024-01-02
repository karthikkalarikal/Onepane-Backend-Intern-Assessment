package utils

import (
	"encoding/json"

	"github.com/karthikkalarikal/assignmentOnepane/pkg/models"
)

// for marcheling the comments
func MarchelCommentsData(responseData []byte) []models.Comments {
	var commentsResponseObject []models.Comments

	json.Unmarshal(responseData, &commentsResponseObject)
	return commentsResponseObject

}

func MarchelPostData(responseData []byte) []models.Post {
	var commentsResponseObject []models.Post

	json.Unmarshal(responseData, &commentsResponseObject)
	return commentsResponseObject

}

func MarchelUserData(responseData []byte) []models.Users {
	var commentsResponseObject []models.Users

	json.Unmarshal(responseData, &commentsResponseObject)
	// fmt.Println("users", string(responseData))
	return commentsResponseObject

}

// for doing json data queries in go itself

func GetCommentsById(comments []models.Comments, postId int) []models.Comments {
	var commentsOfPost []models.Comments
	for _, comment := range comments {
		if comment.PostId == postId {
			commentsOfPost = append(commentsOfPost, comment)
		}
	}
	return commentsOfPost
}

func GetUserByID(users []models.Users, id int) models.Users {

	for _, user := range users {
		if user.Id == id {
			return user
		}
	}
	return models.Users{}
}
