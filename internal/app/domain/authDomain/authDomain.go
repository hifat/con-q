package authDomain

type ReqRegister struct {
	Username string `binding:"required,max=100" json:"username" example:"conq"`   // Your username
	Password string `binding:"required,max=75" json:"password" example:"Cq1234_"` // Your password
	Name     string `binding:"required,max=100" json:"name" example:"Corn Dog"`   // Your full name
}
