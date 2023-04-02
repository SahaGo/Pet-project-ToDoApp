package Pet_project_ToDoApp //файл с описанием структуры и сущностей

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
