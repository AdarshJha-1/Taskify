package response

type Response struct {
	Status  int                    `json:"status"`
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
