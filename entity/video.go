package entity

type Person struct {
	FirstName string `json:"firstname,omitempty" binding:"required"`
	LastName  string `json:"lastname,omitempty" binding:"required"`
	Age       int8   `json:"age,omitempty" binding:"gte=1,lte=100"`
	Email     string `json:"email,omitempty" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title,omitempty" binding:"min=2,max=100" validate:"is-cool"`
	Description string `json:"description,omitempty" binding:"required,max=500"`
	URL         string `json:"url,omitempty" binding:"required,url"`
	Author      Person `json:"author,omitempty" binding:"required"`
}
