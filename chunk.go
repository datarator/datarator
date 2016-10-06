package main

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
	out []byte
	err error
}

type ChunkProcessor interface {
	Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output
}

type ChunkProcessorFactory struct {
}

func (cf ChunkProcessorFactory) CreateChunkProcessor() (ChunkProcessor, error) {
	return ChunkProcessorSerial{}, nil
}

func processChunk(chunk *Chunk, template Template, doneChannel <-chan struct{}) <-chan Output {
	outChannel := make(chan Output)
	go func() {
		defer close(outChannel)
		for chunk.index = chunk.from; chunk.index < chunk.to; chunk.index++ {
			out, err := template.Generate(chunk)
			outChannel <- Output{
				out: out,
				err: err,
			}
		}
	}()

	return outChannel
}
