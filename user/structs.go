package user

type Data struct {
	Email   string
	Name    string
	Picture string
}

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type CreateRequest struct {
	Email string `json:"email"`
}

type CreateResponse struct {
	Rows   int64  `json:"rows"`
	Status string `json:"status"`
}
