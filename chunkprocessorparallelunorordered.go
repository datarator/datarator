package main

type ChunkProcessorParallelUnordered struct {
}

func (cp ChunkProcessorParallelUnordered) Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output {
	return merge(doneChannel, cp.buildChannels(count, template))
}

func (cp ChunkProcessorParallelUnordered) buildChannels(count int, template Template) [](<-chan Output) {
	chunks := buildChunks(count)
	chunkChannels := [](<-chan Output){}

	for i := 0; i < len(chunks); i++ {
		chunkChannels = append(chunkChannels, cp.processChunk(&chunks[i], template))
	}

	return chunkChannels
}

func (cp ChunkProcessorParallelUnordered) processChunk(context *Context, template Template) <-chan Output {
	outChannel := make(chan Output)
	go func() {
		out, err := template.Generate(context)
		outChannel <- Output{
			out: out,
			err: err,
		}
		close(outChannel)
	}()

	return outChannel
}
