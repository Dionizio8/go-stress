package entity

import (
	"fmt"
	"strconv"
	"time"
)

type Stress struct {
	URL            string
	Requests       int
	Concurrency    int
	StartTime      time.Time
	EndTime        time.Time
	Duration       time.Duration
	ListStatusOK   int
	ListStatusFail map[int]int
}

func NewStress(url string, requests, concurrency int) *Stress {
	return &Stress{
		URL:            url,
		Requests:       requests,
		Concurrency:    concurrency,
		StartTime:      time.Now(),
		ListStatusFail: make(map[int]int, 0),
	}
}

func (s *Stress) Start() {
	s.StartTime = time.Now()
}

func (s *Stress) End() {
	s.EndTime = time.Now()
	s.Duration = s.EndTime.Sub(s.StartTime)
}

func (s *Stress) AddListStatusOK(int) {
	s.ListStatusOK++
}

func (s *Stress) AddListStatusFail(statusCode int) {
	s.ListStatusFail[statusCode]++
}

func (s *Stress) ReportCLI() {
	fmt.Print("Stress Report 🚀...\n\n")
	fmt.Println("URL:                   ", s.URL)
	fmt.Println("Requests:              ", s.Requests)
	fmt.Println("Concurrency:           ", s.Concurrency)
	fmt.Println("Duration:              ", s.Duration.Seconds(), "s")
	fmt.Println("Total Status Code 200: ", s.ListStatusOK)
	for statusCode, count := range s.ListStatusFail {
		fmt.Println("Total Status Code", strconv.Itoa(statusCode)+": ", count)
	}
}
