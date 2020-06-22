package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var (
	invalidMonitor  = 0
	invalidPrisoner = 0
	wg              sync.WaitGroup
	cpuprofile      = flag.String("cpuprofile", "", "write cpu profile to file")
)

func simulateNormal(person int, randGenerator *rand.Rand) int {
	counter := 0
	personSet := make(map[int]int, person)
	for len(personSet) < person {
		choosePerson := randGenerator.Intn(person)
		_, ok := personSet[choosePerson]
		if !ok {
			personSet[choosePerson] = choosePerson
		}
		counter++
	}
	return counter
}

func simulateMonitor(person int, randGenerator *rand.Rand) int {
	counter := 0
	personSet := make(map[int]int, person)
	processSet := make([]int, 0)
	switchOn := false
	for {
		choosePerson := randGenerator.Intn(person)
		if choosePerson == 0 {
			if switchOn {
				switchOn = false
				if len(personSet) == person-1 {
					fmt.Printf("%v\n", processSet)
					break
				}
			} else {
				invalidMonitor++
			}
		} else {
			_, ok := personSet[choosePerson]
			if !ok {
				if switchOn {
					invalidPrisoner++
				} else {
					switchOn = true
					personSet[choosePerson] = choosePerson
				}
			}
		}
		processSet = append(processSet, choosePerson)
		counter++
	}
	return counter
}

func calculate(person int, loopTimes int, simulateFunc func(int, *rand.Rand) int) []int {
	index := 0
	go displayPercentage(loopTimes, &index)

	counterSet := make([]int, loopTimes)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index = range counterSet {
		counter := simulateFunc(person, r)
		counterSet[index] = counter
	}

	sort.Ints(counterSet)
	return counterSet
}

func displayPercentage(loopTimes int, index *int) {
	for *index < loopTimes-1 {
		fmt.Printf("%0.2f %%\n", float64(*index*100)/float64(loopTimes))
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func eachCalculate(person int, eachLoopTimes int, simulateFunc func(int, *rand.Rand) int, seed int64, channel chan int) {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano() * seed))
	for i := 0; i < eachLoopTimes; i++ {
		channel <- simulateFunc(person, randGenerator)
		if i%1000 == 0 {
			fmt.Printf("%0.2f %%\n", float64(i*100)/float64(eachLoopTimes))
		}
	}
}

func calculateSimultaneous(person int, loopTimes int, simulateFunc func(int, *rand.Rand) int) []int {
	counterSet := make([]int, loopTimes)
	eachLoopTimes := loopTimes / 4
	channel := make(chan int, loopTimes)
	for i := 1; i <= 4; i++ {
		go eachCalculate(person, eachLoopTimes, simulateFunc, int64(i), channel)
	}
	for i := 0; i < loopTimes; i++ {
		counterSet[i] = <-channel
	}

	wg.Done()
	sort.Ints(counterSet)
	return counterSet
}

func main() {
	//// pprof monitor
	//flag.Parse()
	//if *cpuprofile != "" {
	//	f, err := os.Create(*cpuprofile)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	_ = pprof.StartCPUProfile(f)
	//	defer pprof.StopCPUProfile()
	//}

	person := 50
	probability := 0.5
	loopTimes := 16
	wg.Add(1)

	fmt.Println("Simulating...")

	startTime := time.Now()
	//counterSet := calculate(person, loopTimes, simulateMonitor)
	counterSet := calculateSimultaneous(person, loopTimes, simulateMonitor)
	endTime := time.Now()

	wg.Wait()
	fmt.Printf("\naverage: %d\n", counterSet[int(probability*float64(loopTimes))])
	fmt.Printf("invalidMonitor: %d\n", invalidMonitor)
	fmt.Printf("invalidPrisoner: %d\n", invalidPrisoner)
	fmt.Printf("Spent %d ms\n", endTime.Sub(startTime).Milliseconds())
}
