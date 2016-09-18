package main

import "sync"

const (
	chunkSize = 1000
)

type Chunk struct {
	from   int
	to     int
	index  int
	values map[string]string
	parent *Chunk
}

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
	if opts.Parallel {
		return ChunkProcessorParallelUnordered{}, nil
	}
	return ChunkProcessorSerial{}, nil
}

func buildChunks(count int) []Chunk {
	chunks := []Chunk{}

	chunksCount := count / chunkSize
	for i := 0; i < chunksCount; i++ {
		chunks = append(chunks, Chunk{
			from:   i * chunkSize,
			to:     (i + 1) * chunkSize,
			index:  i * chunkSize,
			values: make(map[string]string),
		})

	}

	oneMoreChunkNeeded := count%chunkSize > 0
	if oneMoreChunkNeeded {
		chunks = append(chunks, Chunk{
			from:   chunksCount * chunkSize,
			to:     count,
			index:  chunksCount * chunkSize,
			values: make(map[string]string),
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

func processChunk(context *Chunk, template Template, doneChannel <-chan struct{}) <-chan Output {
	outChannel := make(chan Output)
	go func() {
		defer close(outChannel)
		out, err := template.Generate(context)
		outChannel <- Output{
			out: out,
			err: err,
		}
	}()

	return outChannel
}
