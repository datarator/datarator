package main

type ChunkProcessorSerial struct {
}

func (cp ChunkProcessorSerial) Process(count int, template Template, doneChannel <-chan struct{}) <-chan Output {
	outChannel := make(chan Output)
	go func() {
		out, err := template.Generate(&Context{
			ToIndex:      count,
			CurrentIndex: []int{0},
		})

		outChannel <- Output{
			out: out,
			err: err,
		}
		close(outChannel)
	}()

	return outChannel

}
