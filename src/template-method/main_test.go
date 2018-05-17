package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	// to make sure the job excuting sequentially
	// only allow 1 worker to handler the jobs
	workerNum := 1
	jobNum := 5

	jobs := make(chan int, jobNum)
	results := make(chan int, jobNum)
	expectedResults := make(chan int, jobNum)

	for i := 0; i < jobNum; i++ {
		jobs <- i
		expectedResults <- i * 2
	}

	for i := 0; i < workerNum; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < jobNum; i++ {
		actual := <-results
		expected := <-expectedResults
		assert.Equal(t, expected, actual)
	}

	close(jobs)
	close(results)
	close(expectedResults)
}

func BenchmarkWorker(b *testing.B) {
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		jobs := make(chan int, 1)
		results := make(chan int, 1)
		jobs <- 1
		// is really hard to test the performance of the worker function
		go worker(1, jobs, results)
		<-results
		close(jobs)
		close(results)
	}
}
