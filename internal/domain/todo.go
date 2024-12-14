package domain

type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	FileID      string `json:"file_id"`
}

