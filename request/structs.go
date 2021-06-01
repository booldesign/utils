package request

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/5/31 11:47
 * @Desc:
 */

type (
	HttpResponse struct {
		Success bool         `json:"success"`
		Data    interface{}  `json:"data"`
		Error   ErrorMessage `json:"error"`
	}

	ErrorMessage struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)
