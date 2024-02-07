package models

type Task struct {
	TaskID      string `json:"task_id"`
	Belongs     string `json:"belongs"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func ValidateTask(task Task, flag bool) (msg string) {
	if task.Title == "" {
		msg = "Title must not be empty"
	} else if len(task.Description) > 100 {
		msg = "Description must not be greater than 100 characters"
	} else if task.Done == true && flag == false {
		msg = "Task must be false by default"
	}
	return
}

var TaskTableQuery = `CREATE TABLE IF NOT EXISTS tasks (
	taskID uuid DEFAULT gen_random_uuid(),
	belongs VARCHAR(100) REFERENCES users(username),
	title VARCHAR(300) NOT NULL,
	description VARCHAR(1000),
	done BOOL NOT NULL,
	PRIMARY KEY(taskID)
);`
