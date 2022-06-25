package serializers

type CreateTodoSerializer struct {
	Content string `json:"content" binding:"required"`
}

type UpdateTodoSerializer struct {
	Content string `json:"content"`
	IsDone  bool   `json:"is_done"`
}
