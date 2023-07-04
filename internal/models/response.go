package models

type Resp struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Detail interface{} `json:"detail"`
	Result string      `json:"result,omitempty"`
}
type RespToken struct {
	Status    int         `json:"status"`
	Msg       string      `json:"msg"`
	Token     string      `json:"token"`
	Signature string      `json:"signature"`
	Note      interface{} `json:"note"`
}
type OrderTask struct {
	OrderId   string `json:"order_id"`
	TaskId    string `json:"task_id"`
	Major     string `json:"major"`
	Services  string `json:"services"`
	TaskState string `json:"task_state"`
	TCreate   string `json:"t_create"`
}
type TenantService struct {
	TenantId   string   `json:"tenant_id"`
	TenantName string   `json:"tenant_name"`
	Services   []string `json:"services"`
}

type RespLocal struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
type RespWeb struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Detail interface{} `json:"detail,omitempty"`
}

type RespWebReport struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Detail interface{} `json:"detail"`
	Header interface{} `json:"header"`
}

type RespPromotionLocal struct {
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	Data       map[string][]string `json:"data"`
}
type ErrorDetail struct {
	TypeError        string `json:"type_error"`
	ErrorDescription string `json:"error_description"`
}

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}
