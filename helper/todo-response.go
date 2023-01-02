package helper

import "github.com/KuraoHikari/golang-todo-api/entity"

type TodoResponse struct {
	ID          int64        `json:"id"`
	ProductName string       `json:"product_name"`
	Image       string       `json:"image"`
	Bounty      uint64       `json:"bounty"`
	User        UserResponse `json:"user,omitempty"`
}

func NewTodoResponse(todo entity.Todo) TodoResponse {
	return TodoResponse{
		ID:          todo.ID,
		ProductName: todo.Name,
		Image:       todo.Image,
		Bounty:      todo.Bounty,
		User:        NewUserResponse(todo.User),
	}
}

func NewTodoArrayResponse(todos []entity.Todo) []TodoResponse {
	todoRes := []TodoResponse{}
	for _, v := range todos {
		p := TodoResponse{
			ID:          v.ID,
			ProductName: v.Name,
			Image:      v.Image,
			Bounty:       v.Bounty,
			User:        NewUserResponse(v.User),
		}
		todoRes = append(todoRes, p)
	}
	return todoRes
}