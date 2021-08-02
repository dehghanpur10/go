package entity

//User is struct for receive user data
type User struct {
	Firstname string `json:"firstname,omitempty" binding:"required,min=3"`
	Lastname  string `json:"lastname,omitempty" binding:"required,min=3" `
	Age       int8   `json:"age,omitempty" binding:"gte=1,lte=130"`
}
type UserLogIn struct {
	Firstname string `json:"firstname,omitempty" binding:"required"`
	Lastname  string `json:"lastname,omitempty" binding:"required"`
}
