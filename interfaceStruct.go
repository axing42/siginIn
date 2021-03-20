package main

type LoginS struct {
	Msg  string `json:"msg"`
	Key  string `json:"_key"`
	User struct {
		UserID          int         `json:"userID"`
		Role            int         `json:"role"`
		Nick            string      `json:"nick"`
		Avatar          string      `json:"avatar"`
		Birthday        int64       `json:"birthday"`
		Age             int         `json:"age"`
		Gender          int         `json:"gender"`
		Level           int         `json:"level"`
		Isgold          int         `json:"isgold"`
		IdentityTitle   interface{} `json:"identityTitle"`
		IdentityColor   int         `json:"identityColor"`
		NeedSetPassword int         `json:"needSetPassword"`
		NeedSetUserInfo int         `json:"needSetUserInfo"`
	} `json:"user"`
	SessionKey string `json:"session_key"`
	Status     int    `json:"status"`
}

type CategoryId struct {
	Categories []struct {
		CategoryID int64  `json:"categoryID"`
		Title      string `json:"title"`
	} `json:"categories"`
}
type Account struct {
	Name []string `json:"name"`
	Pwd  string   `json:"pwd"`
}
