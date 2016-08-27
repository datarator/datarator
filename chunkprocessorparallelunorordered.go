package main

type ChunkProcessorParallelUnordered struct {
}

func (cp ChunkProcessorParallelUnordered) Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output {
	return merge(doneChannel, cp.buildChannels(count, template, doneChannel))
}

func (cp ChunkProcessorParallelUnordered) buildChannels(count int, template Template, doneChannel <-chan struct{}) [](<-chan Output) {
	chunks := buildChunks(count)
	chunkChannels := [](<-chan Output){}

	for i := 0; i < len(chunks); i++ {
		chunkChannels = append(chunkChannels, processChunk(&chunks[i], template, doneChannel))
	}

	return chunkChannels
}
