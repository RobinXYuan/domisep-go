package constants

//ErrorRrturn 错误信息返回格式
type ErrorReturn struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}
