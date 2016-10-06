package main

import "bytes"

type ChunkProcessorSerial struct {
}

func (cp ChunkProcessorSerial) Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output {
	return cp.processChunks(buildChunks(count), template, doneChannel)
}

func (cp ChunkProcessorSerial) processChunks(chunks []*Chunk, template Template, doneChannel <-chan struct{}) <-chan Output {
	outChannel := make(chan Output)
	go func() {
		defer close(outChannel)

		for i := 0; i < len(chunks); i++ {
			chunk := chunks[i]
			var buffer bytes.Buffer
			for chunk.index = chunk.from; chunk.index < chunk.to; chunk.index++ {
				out, err := template.Generate(chunk)
				if err != nil {
					outChannel <- Output{
						out: out,
						err: err,
					}
				}
				buffer.Write(out)
			}

			outChannel <- Output{
				out: buffer.Bytes(),
				err: nil,
			}
		}
	}()

	return outChannel
}
