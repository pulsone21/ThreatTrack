package workbook

import "ContentManagement/api/models/task"

type Workbook struct {
	Name string
	Task []task.Task
}
