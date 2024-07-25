package usecase

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Dionizio8/go-stress/internal/entity"
)

type CreateStress struct {
	URL         string
	Requests    int
	Concurrency int
	PrintLog    bool
}

func NewCreateStress(url string, requests, concurrency int, printLog bool) *CreateStress {
	return &CreateStress{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
		PrintLog:    printLog,
	}
}

func (c *CreateStress) Execute() {
	stress := entity.NewStress(c.URL, c.Requests, c.Concurrency)
	stress.Start()
	fmt.Println("Initializing test 🧐")

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, c.Concurrency)
	for i := 0; i < c.Requests; i++ {
		semaphore <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			statusCode := c.getHTTPRequest()
			if statusCode == 200 {
				stress.AddListStatusOK(statusCode)
			} else {
				stress.AddListStatusFail(statusCode)
			}
		}()
	}

	wg.Wait()
	close(semaphore)

	stress.End()
	fmt.Println("Finishing test 🥵")
	stress.ReportCLI()
}

func (c *CreateStress) getHTTPRequest() int {
	req, err := http.NewRequest(http.MethodGet, c.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.printLog(resp.StatusCode)
		return resp.StatusCode
	}
	defer resp.Body.Close()

	c.printLog(resp.StatusCode)
	return resp.StatusCode
}

func (c *CreateStress) printLog(statusCode int) {
	if c.PrintLog {
		fmt.Println("Status Code: ", statusCode)
	}
}
