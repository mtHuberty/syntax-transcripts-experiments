package main

import (
	"fmt"
	"github.com/AdamSiekierski/syntax-experiments/utils"
)

// Just print out the first episode

func main() {
	transcripts, err := utils.LoadTranscripts()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(utils.Format(transcripts[0]))
	}
}
