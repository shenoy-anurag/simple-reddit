package configs

const (
	API_SUCCESS string = "success"
	API_FAILURE string = "failure"
	API_ERROR   string = "error"
)

type APIResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
