package entity

//User is struct for receive user data
type User struct {
	Firstname string `json:"firstname,omitempty" binding:"required"`
	Lastname  string `json:"lastname,omitempty" binding:"required" `
	Age       int8   `json:"age,omitempty" binding:"gte=1,lte=130,required"`
}
