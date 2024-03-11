package internal

type TodoService struct {
}

type Todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var (
	todos = []Todo{
		{
			Id:        1,
			Task:      "Learning Golang Basic",
			Completed: true,
		},
		{
			Id:        2,
			Task:      "Learning Chi",
			Completed: true,
		},
		{
			Id:        3,
			Task:      "Learning Database",
			Completed: false,
		},
	}
)

func (receiver TodoService) deleteTodo(id int) []Todo {
	for i := 0; i < len(todos); i++ {
		if id == todos[i].Id {
			if len(todos) == 1 {
				todos = []Todo{}
			} else {
				todos = append(todos[:i], (todos)[i+1:]...)
			}

			return todos
		}
	}

	return todos
}

func (receiver TodoService) updateTodo(todo Todo) []Todo {
	for i := 0; i < len(todos); i++ {
		if todo.Id == todos[i].Id {
			todos[i] = todo
		}
	}

	return todos
}

func (receiver TodoService) createTodo(todo Todo) []Todo {
	todo.Id = getLastId() + 1
	todos = append(todos, todo)

	return todos
}

func getLastId() int {
	if len(todos) == 0 {
		return 0
	}

	lastId := todos[0].Id
	for i := 0; i < len(todos); i++ {
		if todos[i].Id > lastId {
			lastId = todos[i].Id
		}
	}

	return lastId
}

func (receiver TodoService) FindById(id int) *Todo {
	for _, todo := range todos {
		if id == todo.Id {
			return &todo
		}
	}

	return nil
}

func (receiver TodoService) ListTodos() []Todo {
	return todos
}
