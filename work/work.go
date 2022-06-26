package work

import (
	"github.com/GoLangDream/iceberg/log"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/robfig/cron/v3"
	"os"
)

type workInfo struct {
	name   string
	spec   string
	taskID cron.EntryID
	cmd    func()
}

var cronTask *cron.Cron = cron.New()
var works = make(map[string]*workInfo)

func Start() {
	cronTask.Start()
	printAllTasks()
}

func Register(name string, spec string, cmd func()) {
	fullCmd := func() {
		log.Infof("定时任务 [%s] 开始执行", name)
		cmd()
		info, ok := works[name]
		if ok {
			task := cronTask.Entry(info.taskID)
			log.Infof("定时任务 [%s] 执行完成，下次执行时间 %s", name, task.Next.Local())
		}
	}

	jobID, _ := cronTask.AddFunc(spec, fullCmd)
	works[name] = &workInfo{
		name:   name,
		spec:   spec,
		taskID: jobID,
		cmd:    fullCmd,
	}
}

func RunWorkNow(name string) {
	info, ok := works[name]
	if ok {
		info.cmd()
	}
}

func printAllTasks() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.Style().Format.Header = text.FormatTitle

	t.SetTitle("定时任务")
	t.Style().Title.Align = text.AlignCenter

	t.AppendHeader(table.Row{"#", "Task Name", "执行频率"})

	index := 1

	for _, info := range works {
		t.AppendRow([]interface{}{
			index + 1,
			info.name,
			info.spec,
		})
		index += 1
	}

	t.Render()
}
