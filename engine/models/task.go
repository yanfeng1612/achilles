package models

import (
	"achilles/engine/db"
	"fmt"
	"time"
)

type Task struct {
	Id           int64  `db:"id"`
	GraphId      int64  `db:"graphId"`
	TraceId      string `db:"traceId"`
	NodeId       int64  `db:"nodeId"`
	State        int    `db:"state"`
	ExecuteCount int    `db:"executeCount"`
	Priority     int    `db:"priority"`
	InitDelay    int    `db:"initDelay"`
	RetryNumber  int    `db:"retryNumber"`
	StartTime    string `db:"startTime"`
	EndTime      string `db:"endTime"`
	CreateTime   string `db:"createTime"`
	RefreshTime  string `db:"refreshTime"`
	TaskNode     *Node
	Graph        *Graph
}

var (
	isTaskRunning = false
)

func Start() error {
	if isTaskRunning {
		return nil
	}
	isTaskRunning = true
	go loop()
	return nil
}

func Shutdown() error {
	if !isTaskRunning {
		return nil
	}
	isTaskRunning = false
	return nil
}

func loop() {
	fmt.Println("begin achilles goroutine")
	for isTaskRunning {
		// todo config delay
		time.Sleep(time.Second * 3)
		for _, task := range getPendingTasks() {
			fmt.Println("begin to process task : ", task.Id)
			processTask(task)
		}
	}
}

func getPendingTasks() []Task {
	list := []Task{}
	err := db.DefaultDB.Select(&list, "SELECT * FROM task WHERE state IN (0,30) AND executeCount <= retryNumber")
	if err != nil {
		fmt.Println(err.Error())
	}
	return list
}

func processTask(task Task) {
	if !getLock(task.Id) {
		fmt.Println("未获取到任务处理锁!")
		return
	}
	graph, _ := getGraphBy(task.Id)
	task.Graph = graph
	var node Node
	for _, n := range graph.NodeList {
		if task.NodeId == n.Id {
			node = n
			break
		}
	}

	task.TaskNode = &node

	switch node.Type {
	case 0:
		// start node
		fmt.Println("begin start!")
		break
	case 1:
		// commmon node
		processCommonNode(node)
		break
	case 2:
		// judge node

		break

	case 3:
		// wait node

		break

	case 4:
		fmt.Println("begin end node!")
		break
	default:
		fmt.Println("node type is't right")
		break
	}

	handleTaskSucc(task)
}

func processCommonNode(node Node) {
	if node.ExecType == 0 {
		// http
		// io.Reader
		// http.Post(node.Url, node.HttpContentType, node.ReqPara)
		fmt.Println("http task impl")
	} else if node.ExecType == 1 {

	} else {

	}
}

func getLock(id int64) bool {
	result, err := db.DefaultDB.Exec("UPDATE task SET state = 1 WHERE id = ? AND state = 0", id)
	if err != nil {
		return false
	}
	count, err := result.RowsAffected()
	if err != nil || count <= 0 {
		return false
	}
	return true
}

func handleTaskSucc(task Task) error {
	nextNodes := getNextNode(task.Graph, task.TaskNode)
	if len(nextNodes) > 1 {
		fmt.Println("not support parallel run")
	}
	if len(nextNodes) < 1 {
		fmt.Println("graph has already go last node!")
		return nil
	}
	tx, err := db.DefaultDB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM task WHERE id = ?", task.Id)
	if err != nil {
		return err
	}
	nextNode := nextNodes[0]
	tx.Exec("INSERT INTO task(graphId,traceId,priority,initDelay,retryNumber,createTime) VALUES (?,?,?,?,?,NOW())", task.GraphId, "uuid-0", nextNode.Priority, nextNode.InitDelay, nextNode.RetryNumber)
	tx.Commit()
	return nil
}

func handleTaskFail(id int64) {
	db.DefaultDB.Exec("UPDATE task SET state = 3 WHERE id = ? AND state = 1", id)
}
