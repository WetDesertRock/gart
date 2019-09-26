package gart

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"time"

	"github.com/wetdesertrock/flexiconfig"
)

// Painting is the core struct that allows you to render and save your image
type Painting struct {
	Renderer *Renderer
	Settings flexiconfig.Settings
}

// NewPainting creates a new painting struct using the settings object. You
// shouldn't need to ever use this yourself for normal use. Instead you should
// use InitProgramWithPainting.
func NewPainting(settings flexiconfig.Settings) *Painting {
	result := Painting{
		Settings: settings,
	}

	result.Renderer = NewRenderer(result.Settings)

	seed, err := result.Settings.GetInt("gart:Seed", 0)
	if err != nil {
		now := time.Now()
		seed = now.UnixNano()
	}
	rand.Seed(seed)
	fmt.Printf("Seed: %d\n", seed)

	return &result
}

// Save saves the painting. It will determine where to save the file based on
// the gart:output:OutFileName and gart:output:OutPath settings.
func (this *Painting) Save() (string, error) {
	outFile, _ := this.Settings.GetString("gart:output:OutFileName", "painting.png")
	outPath, _ := this.Settings.GetString("gart:output:OutPath", "./output")
	outFilePath := path.Join(outPath, outFile)

	fmt.Println("Writing to ", outFilePath)
	this.Renderer.SavePNG(outFilePath)

	// Save our config as <outFile>.json if requested
	saveConfig, _ := this.Settings.GetBool("gart:output:SaveConfig", false)
	if saveConfig {
		configPath := outFilePath + ".json"
		fmt.Printf("Saving config file to: %s\n", configPath)
		err := ioutil.WriteFile(configPath, this.Settings.GetJSON(), 0644)
		if err != nil {
			fmt.Printf("Error saving config file: %v\n", err)
		}
	}

	return outFilePath, nil
}
