package config

import (
	"context"

	"github.com/murilocarbol/stress-test/application/client"
	"github.com/murilocarbol/stress-test/application/usecase"
)

func Initialize(url *string, requests, concurrency *int) {
	// Client
	genericClient := client.NewGenericClient()

	// Usecase
	stressUseCase := usecase.NewStressUseCase(genericClient)

	stressUseCase.StressUseCase(context.Background(), url, requests, concurrency)
}
