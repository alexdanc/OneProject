package TaskService

type RequestBody struct {
	ID     string `gorm:"primaryKey;type:serial" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
