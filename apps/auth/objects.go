package auth

//CreateRoleParams 创建角色参数
type CreateRoleParams struct {
	Name        string   `json:"name" binding:"required,max=40"`
	Permissions []string `json:"permissions" binding:"required"`
}

//CreateRoleReturn 创建角色返回值
type CreateRoleReturn struct {
	ID          string   `json:"id" binding:"required,uuid"`
	Name        string   `json:"name" binding:"required,max=40"`
	Permissions []string `json:"permissions" binding:"required"`
}

//LoginParams 用户登录参数
type LoginParams struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//LoginResult 登录成功返回参数
type LoginResult struct {
	Token string `json:"token"`
}

//RegisterParams 用户注册参数
type RegisterParams struct {
	Name            string `json:"name" binding:"required,min=4,max=80"`
	Password        string `json:"password" binding:"required,password"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
	Email           string `json:"email" binding:"required,email"`
	Gender          string `json:"gender"`
}
