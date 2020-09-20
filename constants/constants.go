package constants

//KEY gin session key
const KEY = "SESSKEYV1"

//ParameterType 参数类型
type ParameterType string

//JSONTypeParam Json 类型入参
//FormTypeParam 表单类型入参
const (
	JSONTypeParam ParameterType = "JSON"
	FormTypeParam ParameterType = "FORM"
)
