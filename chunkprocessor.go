package main

import "sync"

const (
	chunkSize = 1000
)

type Output struct {
	out string
	err error
}

type ChunkProcessor interface {
	Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output
}

type ChunkProcessorFactory struct {
}

func (cf ChunkProcessorFactory) CreateChunkProcessor() (ChunkProcessor, error) {

	return ChunkProcessorSerial{}, nil

	// return ChunkProcessorParallelUnordered{}, nil
}

func buildChunks(count int) []Context {
	chunks := []Context{}

	chunksCount := count / chunkSize
	for i := 0; i < chunksCount; i++ {
		chunks = append(chunks, Context{
			FromIndex:    i * chunkSize,
			ToIndex:      (i + 1) * chunkSize,
			CurrentIndex: []int{i * chunkSize},
		})

	}

	oneMoreChunkNeeded := count%chunkSize > 0
	if oneMoreChunkNeeded {
		chunks = append(chunks, Context{
			FromIndex:    chunksCount * chunkSize,
			ToIndex:      count,
			CurrentIndex: []int{chunksCount * chunkSize},
		})
	}
	return chunks
}

// from https://blog.golang.org/pipelines
func merge(done <-chan struct{}, cs [](<-chan Output)) <-chan Output {
	var wg sync.WaitGroup
	out := make(chan Output)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan Output) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
