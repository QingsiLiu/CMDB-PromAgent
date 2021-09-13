package init

import (
	"promagent/task"
	"promagent/task/plugins"
)

func init() {
	task.Register("register", &plugins.Register{})
}
