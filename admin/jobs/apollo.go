package jobs

import (
	"log"
)

type Jober interface {
	Execute()
}

type ApolloJob struct {
	ID   int
	Name string
}

// NewJob 创建Job
func NewJob(jobName string) ApolloJob {
	log.Println("begin apollo job!")
	newJob := ApolloJob{}
	newJob.ID = 1
	newJob.Name = jobName

	log.Println("end apollo job!")
	return newJob
}

// TriggerJob 触发job
func TriggerJob(id int) {

}
