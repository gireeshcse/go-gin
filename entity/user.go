package entity

// Address : stores address of user
type Address struct {
	Street  string `json:"street"`
	HouseNo string `json:"house_no"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	Country string `json:"country" binding:"required"`
}

// User : Defines user entity
type User struct {
	Name       string `json:"name" xml:"title" form:"title"`
	Username   string `json:"username" binding:"min=2,max=10" validate:"containsUser"`
	Email      string `json:"email" xml:"email" form:"email" binding:"required,email" `
	ProfileURL string `json:"profile_url" binding:"required,url"`
	Age        int8   `json:"age" binding:"gte=1,lte=130"`
	Address    string `json:"address" binding:"required"`
}
