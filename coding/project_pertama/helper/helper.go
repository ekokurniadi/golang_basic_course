package helper

type Meta struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ServerSideResponse struct {
	TotalItems   int         `json:"totalItems"`
	Data         interface{} `json:"data"`
	TotalPages   int         `json:"totalPages"`
	CurrentPages int         `json:"currentPages"`
}

func ServerSideResponses(totalItems int, totalPages int, currentPages int, data interface{}) ServerSideResponse {
	serverSideResponse := ServerSideResponse{
		TotalItems:   totalItems,
		TotalPages:   totalPages,
		CurrentPages: currentPages,
		Data:         data,
	}
	jsonResponse := serverSideResponse
	return jsonResponse
}

func ApiResponse(status string, message string, code int, data interface{}) Response {
	meta := Meta{
		Status:  status,
		Message: message,
		Code:    code,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}
