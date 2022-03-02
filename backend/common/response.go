package common

const (
	API_SUCCESS                string = "success"
	API_FAILURE                string = "failure"
	API_ERROR                  string = "error"
	REQUEST_VALIDATION_ERROR   string = "request_validation_error"
	INVALID_REQUEST_DATA_ERROR string = "request_invalid_error"
)

type APIResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type APIMessage struct {
	Status  int    `json:"status" uri:"status" validate:"required"`
	Message string `json:"message" uri:"message" validate:"required"`
}
