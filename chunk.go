package main

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

func buildChunks(count int) []*Chunk {
	chunkSize := opts.ChunkSize
	chunks := []*Chunk{}

	chunksCount := count / chunkSize
	for i := 0; i < chunksCount; i++ {
		chunks = append(chunks, &Chunk{
			from:   i * chunkSize,
			to:     (i + 1) * chunkSize,
			index:  i * chunkSize,
			values: make(map[string]string),
		})
	}

	oneMoreChunkNeeded := count%chunkSize > 0
	if oneMoreChunkNeeded {
		chunks = append(chunks, &Chunk{
			from:   chunksCount * chunkSize,
			to:     count,
			index:  chunksCount * chunkSize,
			values: make(map[string]string),
		})
	}
	return chunks
}
