package users

import (
	"restfulAPI/data"
	"testing"

	
	"github.com/stretchr/testify/assert"
)

func TestGetUserPosts (t *testing.T){
	asserter := assert.New(t)

	user := &data.User{
		ID:       1,
		UserInfo: data.UserInfo{
			Name:     "Jon Snow",
			Username: "direpup45",
			Email:    "jon.snow@winterfell.com",
		},
	}

	posts := []*data.Post{
		{
			UserId: 1,
			Id:     1,
			Title:  "the wall",
			Body:   "it is very cold up here",
		},
		{
			UserId: 1,
			Id:     2,
			Title:  "What do I know",
			Body:   "nothing",
		},
	}


	expectedUP := &data.UserPosts{
		UserId:   1,
		UserInfo: user.UserInfo,
		Posts:    posts,
	}

	up := getUserPosts(user, posts)
	asserter.Equal(expectedUP, up)
}