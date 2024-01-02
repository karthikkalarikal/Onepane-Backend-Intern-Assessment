package usecase

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/karthikkalarikal/assignmentOnepane/pkg/models"
	"github.com/karthikkalarikal/assignmentOnepane/pkg/utils"
)

func fetchComments() []models.Comments {
	commetsURL := "https://jsonplaceholder.typicode.com/comments"
	commentsData, err := http.Get(commetsURL)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	commentsByte, err := io.ReadAll(commentsData.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	comments := utils.MarchelCommentsData(commentsByte)

	return comments
}

func fetchPosts() []models.Post {
	postsURL := "https://jsonplaceholder.typicode.com/posts"
	postsData, err := http.Get(postsURL)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	postsByte, err := io.ReadAll(postsData.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	posts := utils.MarchelPostData(postsByte)

	return posts
}

func fetchUsers() []models.Users {
	usersURL := "https://jsonplaceholder.typicode.com/users"
	usersData, err := http.Get(usersURL)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	usersByte, err := io.ReadAll(usersData.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	users := utils.MarchelUserData(usersByte)

	return users
}

func FetchDataConcurrently() ([]models.Comments, []models.Post, []models.Users) {
	var comments []models.Comments
	var posts []models.Post
	var users []models.Users

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		comments = fetchComments()
	}()

	go func() {
		defer wg.Done()
		posts = fetchPosts()
	}()

	go func() {
		defer wg.Done()
		users = fetchUsers()
	}()

	wg.Wait()

	return comments, posts, users

}

func CombineResults(commentsData []models.Comments, postsData []models.Post, usersData []models.Users) []models.Response {

	combinedResults := []models.Response{}

	for _, post := range postsData {

		commentsOfPost := utils.GetCommentsById(commentsData, post.ID)
		userWhoPosted := utils.GetUserByID(usersData, post.UserID)

		result := models.Response{
			PostId:        post.ID,
			PostName:      post.Title,
			CommentsCount: len(commentsOfPost),
			UserName:      userWhoPosted.UserName,
			Body:          post.Body,
		}

		combinedResults = append(combinedResults, result)
	}

	return combinedResults
}
