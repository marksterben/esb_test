package handler

type Usecase interface {
}

type ResponseError struct {
	Message string `json:"message"`
}

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{usecase}
}
