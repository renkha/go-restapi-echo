package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type M map[string]interface{}

func ResponseFormatter(code int, status string, message interface{}, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}

// func ErrorFormatter(err error) []string {
// 	var errors []string
// 	if ev, ok := err.(validator.ValidationErrors); ok {
// 		for _, e := range ev {
// 			errors = append(errors, e.Error())
// 		}
// 	}

// 	return errors
// }
