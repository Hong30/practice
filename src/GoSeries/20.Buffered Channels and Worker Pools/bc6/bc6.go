package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

// 所有 Job 结构体变量都会有 id 和 randomno 两个字段，randomno 用于计算其每位数之和。
// 而 Result 结构体有一个 job 字段，表示所对应的作业，还有一个 sumofdigits 字段，表示计算的结果（每位数字之和）。
// 工作协程（Worker Goroutine）会监听缓冲信道 jobs 里更新的作业。一旦工作协程完成了作业，其结果会写入缓冲信道 results。
// digits 函数的任务实际上就是计算整数的每一位之和，最后返回该结果。
// 为了模拟出 digits 在计算过程中花费了一段时间，我们在函数内添加了两秒的休眠时间。
// 读取 jobs 信道的数据，根据当前的 job 和 digits 函数的返回值，创建了一个 Result 结构体变量，然后将结果写入 results 缓冲信道。
// worker 函数接收了一个 WaitGroup 类型的 wg 作为参数，当所有的 jobs 完成的时候，调用了 Done() 方法。
// createWorkerPool 函数创建了一个 Go 协程的工作池。
// 上面函数的参数是需要创建的工作协程的数量。在创建 Go 协程之前，它调用了 wg.Add(1) 方法，于是 WaitGroup 计数器递增。
// 接下来，我们创建工作协程，并向 worker 函数传递 wg 的地址。创建了需要的工作协程后，函数调用 wg.Wait()，等待所有的 Go 协程执行完毕。
// 所有协程完成执行之后，函数会关闭 results 信道。
// 因为所有协程都已经执行完毕，于是不再需要向 results 信道写入数据了。
// 现在我们已经有了工作池，我们继续编写一个函数，把作业分配给工作者。
// 上面的 allocate 函数接收所需创建的作业数量作为输入参数，生成了最大值为 998 的伪随机数，并使用该随机数创建了 Job 结构体变量。
// 这个函数把 for 循环的计数器 i 作为 id，最后把创建的结构体变量写入 jobs 信道。当写入所有的 job 时，它关闭了 jobs 信道。
// 下一步是创建一个读取 results 信道和打印输出的函数。
// result 函数读取 results 信道，并打印出 job 的 id、输入的随机数、该随机数的每位数之和。
// result 函数也接受 done 信道作为参数，当打印所有结果时，done 会被写入 true。
// 现在一切准备充分了。我们继续完成最后一步，在 main() 函数中调用上面所有的函数。
// 我们首先在 main 函数的第 2 行，保存了程序的起始时间，并在最后一行（第 12 行）计算了 endTime 和 startTime 的差值，显示出程序运行的总时间。
// 由于我们想要通过改变协程数量，来做一点基准指标（Benchmark），所以需要这么做。
// 我们把 noOfJobs 设置为 100，接下来调用了 allocate，向 jobs 信道添加作业。
// 我们创建了 done 信道，并将其传递给 result 协程。于是该协程会开始打印结果，并在完成打印时发出通知。
// 通过调用 createWorkerPool 函数，我们最终创建了一个有 10 个协程的工作池。main 函数会监听 done 信道的通知，等待所有结果打印结束。
// 程序总共会打印 100 行，对应着 100 项作业，然后最后会打印一行程序消耗的总时间。你的输出会和我的不同，
// 因为 Go 协程的运行顺序不一定，同样总时间也会因为硬件而不同。
// 在我的例子中，运行程序大约花费了 20 秒。
// 现在我们把 main 函数里的 noOfWorkers 增加到 20。我们把工作者的数量加倍了。
// 由于工作协程增加了（准确说来是两倍），因此程序花费的总时间会减少（准确说来是一半）。
