package taskstatus

// // github.com/zzzhr1990/go-common-entity/taskstatus
const (
	// Prepare task
	Prepare int32 = 0
	//Add add
	Add int32 = 100
	//Queue the process fetch but not start
	Queue int32 = 200
	//Process the
	Process int32 = 300
	//Pause paused
	Pause int32 = 400
	//Stop tje
	Stop int32 = 500
	//Fail the task failed.
	Fail int32 = 900
	//Success the task success
	Success int32 = 1000
)
