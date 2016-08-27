package main

type ChunkProcessorSerial struct {
}

func (cp ChunkProcessorSerial) Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output {
	return processChunk(&Context{
		ToIndex:      count,
		CurrentIndex: []int{0},
	}, template, doneChannel)
}
