package usecase

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/murilocarbol/stress-test/application/client"
)

type StressUseCaseInterface interface {
	StressUseCase(url string, requests, calls int) (int, error)
}

type StressUseCase struct {
	genericClient client.GenericClient
}

func NewStressUseCase(genericClient *client.GenericClient) *StressUseCase {
	return &StressUseCase{
		genericClient: *genericClient,
	}
}

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

func (s *StressUseCase) StressUseCase(ctx context.Context, url *string, requests, concurrency *int) {
	fmt.Sprintln("Iniciando teste de carga...")

	start := time.Now()
	results := make(chan Result, *requests)
	var wg sync.WaitGroup

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < *requests / *concurrency; j++ {
				s.makeRequest(*url, results)
			}
		}()
	}

	wg.Wait()
	close(results)

	s.report(results, *requests, time.Since(start))
}

func (s *StressUseCase) makeRequest(url string, results chan<- Result) {
	start := time.Now()
	resp, err := s.genericClient.CallClient(url)
	duration := time.Since(start)

	if err != nil {
		results <- Result{Error: err}
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		results <- Result{Error: err}
		return
	}

	results <- Result{StatusCode: resp.StatusCode, Duration: duration}
}

func (s *StressUseCase) report(results <-chan Result, totalRequests int, totalTime time.Duration) {
	statusCounts := make(map[int]int)
	var totalDuration time.Duration
	successfulRequests := 0

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Erro ao fazer request: %v\n", result.Error)
			continue
		}

		if result.StatusCode == http.StatusOK {
			successfulRequests++
		}
		statusCounts[result.StatusCode]++
		totalDuration += result.Duration
	}

	fmt.Println("Relatório de Teste de Carga:")
	fmt.Printf("Tempo total gasto na execução: %v\n", totalTime)
	fmt.Printf("Total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Requests com status HTTP 200: %d\n", successfulRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for status, count := range statusCounts {
		if status != http.StatusOK {
			fmt.Printf("Status %d: %d requests\n", status, count)
		}
	}
	fmt.Printf("Tempo médio por request: %v\n", totalDuration/time.Duration(totalRequests))
}
