package main

import (
	"os"
	"fmt"
	"time"
	"os/signal"
	"strconv"
	"syscall"
	"timeloop/timer"
)

var (
	timerCtl *timer.TimerHeapHandler
)

func init() {
	timerCtl =  timer.New(10, 1000)
}

type timerHandler struct {
}

func AddTimerTask(dueInterval int, taskId string) {
	timerCtl.AddFuncWithId(time.Duration(dueInterval)*time.Second, taskId, func() {
		fmt.Printf("taskid is %v, time Duration is %v", taskId, dueInterval )	
	})
}

func (t *timerHandler) StartLoop() {
	timerCtl.StartTimerLoop(timer.MIN_TIMER) // 扫描的间隔时间 eq cpu hz/tick
}


func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	timerEntry := timerHandler{}
	timerEntry.StartLoop()

	num := 5000
	interval := 1000 * time.Millisecond
	go func (){
		for i := 0; i < num; i++ {
			taskId := strconv.Itoa(i)
			timerCtl.AddFuncWithId(10 *interval, taskId, func() {
				fmt.Printf("taskid is %v, time Duration is %v \n", taskId, interval )	
			})
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	<- sigs
}
