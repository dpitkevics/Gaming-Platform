package responses

type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
}
