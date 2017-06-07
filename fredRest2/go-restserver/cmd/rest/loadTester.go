package main

import (
	"fmt"
	"net/http"
	"strconv"
	//"reflect"
	"runtime"
	"strings"
	"time"
	"math/rand"
)

const (
LIMIT int = 1
HTTPCALL_FLAG bool = false

//LIMIT int = 100
//HTTPCALL_FLAG bool = true
)
var fin = make(chan string)
var limits = make(chan int, 1000)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func doit(url string) {
	limits <- 1
	defer func() {
		<-limits
	}()
	httpget(url)
}

func httpget(url string) {
	resp, err := http.Get(url)
	defer func() {
		//fmt.Printf("url is %s, goroutine id is %d\n", url, GoID())
		//v := reflect.ValueOf(resp)
		//fmt.Printf("pv, %p,r, %v\n",v,resp)
		//count := v.NumField()
		//for i := 0; i < count; i++ {
		//    f := v.Field(i)
		//    switch f.Kind() {
		//        case reflect.String:
		//            fmt.Println(f.String())
		//        case reflect.Int:
		//            fmt.Println(f.Int())
		//    }
		//    //fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//}
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		fmt.Println(url, err)
	} else {
		fmt.Println(url, resp.StatusCode)
	}
	fin <- url
}
	func numRand(min, max int) int {
		rand.Seed(time.Now().UTC().UnixNano())
		return rand.Intn(max-min) + min
	}
	func allocMemory(sizeIn string) bool {
		var size int

		if  sizeIn == "large"{
			size = 1024 * 1024 * 256
		}else if  sizeIn == "medium" {
			size = 1024 * 1024 * 124
		}else if  sizeIn == "small"{
			size = 1024 * 1024 * 52
		}else{
			size = 1024 * 1024 * 256
		}

		t := time.Now()
		for j := 0; j < LIMIT; j += 1 {
			a := make([]int, size)
			for i := 0; i < size; i += 1 {
				a[i] = i
			}
			a = nil
		}
		t1 := time.Now()
		fmt.Printf("Duration: %1d", t1.Sub(t).Seconds())
		return true
	}
	func allocCpu(sizeIn string) bool {
	var size int

	if  sizeIn == "large"{
		size = 1024 * 1024 * 1024
	}else if  sizeIn == "medium" {
		size = 1024 * 1024 * 512
	}else if  sizeIn == "small"{
		size = 1024 * 1024 * 1024
	}else{
		size = 1024 * 1024 * 1024
	}

	t := time.Now()
	for j := 0; j < LIMIT; j += 1 {
		a := make([]int, size)
		for i := 0; i < size; i += 1 {
			a[i] = i
		}
		a = nil
	}
	t1 := time.Now()
	fmt.Printf("Duration: %1d", t1.Sub(t).Seconds())
	return true
	}

func main() {
	aryUrl := []string{"http://127.0.0.1:3333", "http://127.0.0.1:3333/articles", "http://127.0.0.1:3333/articles/1"}
	aryCpuOperation := []string{"small", "medium", "large"}
	aryMemoryOperation := []string{"small", "medium", "large"}
        var MULTICORE int = runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(MULTICORE) //running in multicore
	fmt.Printf("Using Max Number of CPU's available with %d core\n", MULTICORE)
	var numParallelthreadsToRun int
	for {

		randomCpuOperationCount := numRand(1, 3)
		fmt.Println("Random API Number To call: ", randomCpuOperationCount)

		randomMemoryOperationCount := numRand(1, 3)
		fmt.Println("Random API Number To call: ", randomMemoryOperationCount)

                randomCpuOperationIndex := numRand(0, 3)
		fmt.Println("Random API Number To call: ", randomCpuOperationIndex)

		randomMemoryOperationIndex := numRand(0, 3)
		fmt.Println("Random API Number To call: ", randomMemoryOperationIndex)


		slcCpuOperation := make([]string, len(aryCpuOperation))
		for i := 0; i<len(slcCpuOperation); i++{
			slcCpuOperation[i] = aryCpuOperation[i]
		}

		slcMemoryOperation := make([]string,len(aryMemoryOperation))
		for i := 0; i<len(slcMemoryOperation); i++{
			slcMemoryOperation[i] = aryMemoryOperation[i]
		}


		allocMemory(slcMemoryOperation[randomMemoryOperationIndex])
		allocCpu(slcCpuOperation[randomCpuOperationIndex])

		t := time.Now()
		fmt.Println(t.Format(time.RFC3339))
		randomTimeToSleep := numRand(1000, 3000)
		fmt.Println(strings.Join([]string{"About to sleep for ", strconv.Itoa(randomTimeToSleep), " seconds"}, ""))
		time.Sleep(time.Duration(randomTimeToSleep)* time.Millisecond)
		randomParallelThreadsToRun := numRand(1, 50)
		numParallelthreadsToRun = randomParallelThreadsToRun
		fmt.Println("Random Number of Parallel Threads to run:", randomParallelThreadsToRun)
		httplist := make([]string, randomParallelThreadsToRun)
		numParallelthreadsToRun = randomParallelThreadsToRun
		for  (HTTPCALL_FLAG == true) {
			for i := 0; i < len(httplist); i++ {
				randomApiToCall := numRand(0, 3)
				fmt.Println("Random API Number To call: ", randomApiToCall)
				//httplist[i] = "http://www.baidu.com/s?wd=search" + strconv.Itoa(i+1)
				//httplist[i] = "http://192.168.6.151:2003/"
				httplist[i] = aryUrl[randomApiToCall]
				go doit(httplist[i])
			}
		}
	}

	for i := 0; i < numParallelthreadsToRun; i++ {
		<-fin
	}


}
