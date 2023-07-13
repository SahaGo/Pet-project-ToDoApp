package Pet_project_ToDoApp //файл с описанием структуры и сущностей

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
