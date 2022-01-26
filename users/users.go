package users

import (
	"fmt"
	"sync"

	"restfulAPI/data"
	"restfulAPI/httputil"
)

func GetUserPosts(id string) (*data.UserPosts, error) {
wg := sync.WaitGroup{}


	//get user information
	wg.Add(1)
	var user *data.User
	var err1 error
	go func(id string) {
		defer wg.Done()

		user, err1 = getUser(id)
	}(id)


	
	//get users posts
	wg.Add(1)
	var posts []*data.Post
	var err2 error
	go func(id string) {
		defer wg.Done()
		
		posts, err2 = getPosts(id)	
	}(id)

	wg.Wait()

	if err1 != nil {
		return nil, fmt.Errorf("error getting user. Error:%v", err1.Error())
	}

	
	if err2 != nil {
		return nil, fmt.Errorf("error getting  posts. Error:%v", err2.Error())
	}

	up := getUserPosts(user, posts) 
	return up, nil 
} 

func getUser(id string) (*data.User, error){
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s", id)
	
	var user *data.User
	err := httputil.MakeGetRequest(url, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getPosts(id string) ([]*data.Post, error){
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%s", id)
	
	var posts []*data.Post

	err := httputil.MakeGetRequest(url, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func getUserPosts(user *data.User, posts []*data.Post) *data.UserPosts {
	return &data.UserPosts{
		UserId:   user.ID,
		UserInfo: user.UserInfo,
		Posts:    posts,
	}
}

