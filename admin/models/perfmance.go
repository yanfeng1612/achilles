package models

import (
	"apollo-auto/cdo"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	// "runtime"
	"sync"
)

// 性能测试请求结构体
type PermanceRequest struct {
	Url           string // 请求url
	Para          string // 请求参数
	ConcurrentNum int    // 并发客户端数目
	RequestNum    int    // 请求次数
}

// 性能测试响应结构体
type PermanceResponse struct {
	state      int      // 状态 0-失败 1-成功
	reponseTxt []string // 响应结果
	tps        float64
}

var wg sync.WaitGroup

func PerfermanceGo(request PermanceRequest) PermanceResponse {
	// log.Println(runtime.GOMAXPROCS)
	wg.Add(request.RequestNum)
	var response PermanceResponse

	tasks := make(chan string, request.RequestNum)

	log.Println("start perfmance test")
	startTime := time.Now()

	// for gr := 1; gr <= request.ConcurrentNum; gr++ {
	// 	go worker(tasks, gr)
	// }

	//
	for i := 0; i < request.RequestNum; i++ {
		// go func() {
		// 	tasks <- fmt.Sprintf("Task : %d", i)
		// }()
		go directWork()
	}

	wg.Wait()

	close(tasks)

	endTime := time.Now()
	log.Println("end perfmance test")

	response.state = 1
	response.tps = float64(request.RequestNum) / endTime.Sub(startTime).Seconds()

	log.Printf("tps : " + strconv.FormatFloat(response.tps, 'E', -1, 32))
	return response
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	task, ok := <-tasks
	if !ok {
		log.Println("Worker: %d : Shutting Down", worker)
		return
	}

	log.Println("Worker : %d  Started  %s", worker, task)

	// do work
	time.Sleep(10 * time.Millisecond)

	log.Println("Worker : %d Completed %s", worker, task)
}

func directWork() {
	defer wg.Done()
	// log.Println("Worker start : ")

	// do work
	httpGet()

	// log.Println("Worker Completed")
}

func httpGet() {
	url := "http://apollocore115.dafy.service/handleTrans.cdo"
	cdoRequest := cdo.NewCDOWithService("ApolloApiService", "triggerTaskNow")
	cdo.HandleTrans(url, cdoRequest)
	// resp, err := http.Get("http://apollocore115.dafy.service/a.cdo")
	// if err != nil {
	// 	// handle error
	// 	return
	// }

	// defer resp.Body.Close()
	// _, err = ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	// handle error
	// }

	// log.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	log.Println(string(body))
}
