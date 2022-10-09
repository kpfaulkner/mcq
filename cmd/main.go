package main

import (
	"fmt"

	"github.com/alaingilbert/mcq"
)

func main() {
	world := mcq.NewWorld("C:/Users/kenfa/projects/mc-world-example/test/")

	/*
		bbox := mcq.New2DBBox(mc.Overworld, 0, 0, 100, 100)
		mcq.Q(world).BBox(bbox).Targets(mc.DiamondOreID).Find(func(result mcq.Result) {
			fmt.Printf("Found diamond ore at %s\n", result.Coord())
		}, mcq.WithBlocks) */

	//bbox := mcq.New2DBBox(mc.Overworld, 0, 0, 1000, 1000)

	mcq.Q(world).FindBlocks(func(result mcq.Result) {
		fmt.Printf("Found something at  %s\n", result.Coord())
	}, mcq.WithBlocks)
}
