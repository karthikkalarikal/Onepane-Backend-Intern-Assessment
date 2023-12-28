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

func FindPostByID(posts []models.Post, postId int) models.Post {
	for _, post := range posts {
		if post.ID == postId {
			return post
		}
	}
	return models.Post{}
}

func FindUserByID(users []models.Users, userId int) models.Users {
	for _, user := range users {
		if user.Id == userId {
			return user
		}
	}
	return models.Users{}
}

func GetCommentsCount(comments []models.Comments, postId int) int {
	count := 0
	for _, cmt := range comments {
		if cmt.PostId == postId {
			count++
		}
	}
	return count
}
