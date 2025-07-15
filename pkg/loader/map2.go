package loader

import (
	"fmt"
	"os"

	"github.com/gden173/geotiff/geotiff"
)

func Load() {
	f, err := os.Open("../resourse/map/NE1_LR_LC.tif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read the geotiffPinned
	gtiff, err := geotiff.Read(f)
	if err != nil {
		panic(err)
	}

	// get the geotiff bounds
	bounds, err := gtiff.Bounds()
	if err != nil {
		panic(err)
	}
	fmt.Println(bounds)
}
