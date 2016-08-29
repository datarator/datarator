package main

type ChunkProcessorSerial struct {
}

func (cp ChunkProcessorSerial) Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output {
	return processChunk(&Chunk{
		to:     count,
		values: make(map[string]string),
	}, template, doneChannel)
}
