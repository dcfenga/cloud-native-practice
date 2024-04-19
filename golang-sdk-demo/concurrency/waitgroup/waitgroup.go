package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。

WaitGroup是一个结构体类型，创建 WaitGroup 类型的变量后，其初始值为零值。
WaitGroup 使用计数器来工作。当调用 WaitGroup 的 Add 并传递一个 int 时，WaitGroup 的计数器会加上 Add 的传参。
要减少计数器，可以调用 WaitGroup 的 Done() 方法。Wait() 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。
*/

type Job struct {
	id       int
	randomno int // 用于计算每位数之和
}

type Result struct {
	job         Job // 用于表示所对应的作业
	sumofdigits int // 用于表示计算的结果（每位数字之和）
}

var jobs = make(chan Job, 10)       //接收作业的缓冲信道
var results = make(chan Result, 10) //写入结果的缓冲信道

func main() {
	num := 3
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)

		// 需传递WaitGroup 的地址，如果没有传递 wg 的地址，
		// 那么每个 Go 协程将会得到一个 WaitGroup 值的拷贝，因而当它们执行结束时，main 函数并不会知道。
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
	fmt.Println("==================================")

	// 工作池：一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
	// 示例：计算所输入数字的每一位的和
	// 思路：
	// 1.创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
	// 2.将作业添加到该输入型缓冲信道中。
	// 3.作业完成后，再将结果写入一个输出型缓冲信道。
	// 4.从输出型缓冲信道读取并打印结果。
	startTime := time.Now()
	numOfJobs := 100

	// 向 jobs 信道添加作业
	go allocate(numOfJobs)

	done := make(chan bool)
	// 该协程打印结果，并在完成打印时发出通知
	go result(done)

	numOfWorkers := 20

	// 创建有 10 个协程的工作池
	createWorkerPool(numOfWorkers)

	// 监听 done 信道的通知，等待所有结果打印结束
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

	/*输出：
	Job id 1, input random no 636, sum of digits 15
	Job id 0, input random no 878, sum of digits 23
	Job id 9, input random no 150, sum of digits 6
	...
	total time taken  20.01081009 seconds
	*/
}

// 创建了Go 协程的工作池
func createWorkerPool(numOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// 把作业分配给工作者
func allocate(numOfJobs int) {
	for i := 0; i < numOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 读取 results 信道和打印输出
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// 计算整数的每一位之和，并返回该结果
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

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
