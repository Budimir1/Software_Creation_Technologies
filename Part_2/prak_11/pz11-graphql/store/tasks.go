package store

// Task представляет сущность задачи в хранилище
type Task struct {
	ID          string
	Title       string
	Description *string
	Done        bool
}

// Tasks — глобальное хранилище задач (в памяти)
var Tasks = []*Task{
	{
		ID:          "t_001",
		Title:       "Первая задача",
		Description: strPtr("Учебный пример"),
		Done:        false,
	},
	{
		ID:          "t_002",
		Title:       "Вторая задача",
		Description: strPtr("GraphQL API"),
		Done:        true,
	},
}

// strPtr вспомогательная функция для получения указателя на строку
func strPtr(s string) *string {
	return &s
}
