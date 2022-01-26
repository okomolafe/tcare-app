package data


type User struct{
	ID int `json:"id"`
	UserInfo
	Address Address   `json:"address"`
	Phone   string `json:"phone"`
	Website  string `json:"website"`
	Company Company `json:"company"`
}

type UserInfo struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email  string `json:"email"`
}

type Address struct {
	Street string `json:"street"`
	Suite string `json:"suite"`
	City string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo Geo  `json:"geo"`
}

type Geo struct {
	Lat string   `json:"lat"`
	Long string   `json:"lng"`
	
}

type Company struct {
	Name  string `json:"name"`
	CatchPhrase  string `json:"catchPhrase"`
	BS  string `json:"bs"`
}

type UserPosts struct {
	UserId int `json:"id"`
	UserInfo UserInfo `json:"userInfo"`
	Posts []*Post `json:"posts"`
}