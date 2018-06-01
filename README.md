## 一个封装的定时器库

底层通过map 操作元素的更改，逻辑删除，通过最小堆/环形数组 实现定时器功能
```
example:
    import timer
    timerCtl =  timer.New(10, 1000) //worker 10, buffer 1000, if worker block ,and buffer full ,will drop 
    timerCtl.StartTimerLoop(timer.MIN_TIMER) // 扫描的间隔时间 eq cpu hz/tick
    
    timerCtl.AddFuncWithId(time.Duration(dueInterval)*time.Second, taskId, func() {
		fmt.Printf("taskid is %v, time Duration is %v", taskId, dueInterval )	
	})
```
