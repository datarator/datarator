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
