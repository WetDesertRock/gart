package gart

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jaffee/commandeer"
	"github.com/spf13/pflag"
	"github.com/wetdesertrock/flexiconfig"
)

// ProgramVersionID is a string that can be set when building in order to
// identify the version ID of the program. For instance, to set it to the
// commit ID of your project you can run go build like this:
//   go build -ldflags "-X github.com/wetdesertrock/gart.ProgramVersionID=$GIT_COMMIT"
// gart will store this ID in the metadata and print it to stdout
var ProgramVersionID string

type CLIArgs struct {
	ConfigPaths  []string `flag:"config-path" short:"c" help:"Loads the specified config file"`
	OutputConfig string   `flag:"config-out" short:"O" help:"Saves the resulting config at the specified path"`
}

// InitProgramWithSimulator will setup the program (parse CLI args, retrive
// settings), and setup a Simulator with the supplied world.
func InitProgramWithSimulator(name string, world World) *Simulator {
	// Setup everything
	settings, _ := InitProgram(name)

	// Setup simulator
	return NewSimulator(world, settings)
}

// InitProgramWithPainting will setup the program (parse CLI args, retrive
// settings), and setup a Painting to use.
func InitProgramWithPainting(name string) *Painting {
	// Setup everything
	settings, _ := InitProgram(name)

	// Setup simulator
	return NewPainting(settings)
}

// InitProgram does all of the mundane tasks needed to run a program. You
// shouldn't need to invoke this directly. It is exposed incase somone needs to
// initialize the program manually.
func InitProgram(name string) (flexiconfig.Settings, *CLIArgs) {
	if ProgramVersionID == "" {
		fmt.Printf("Running %s\n", name)
	} else {
		fmt.Printf("Running %s (version: %s)\n", name, ProgramVersionID)
	}

	// Parse the CLI arguments
	cliArgs := InitCLIArgs()

	// Load our settings
	settings := InitSettings(cliArgs)

	initMetadata(settings)

	// Save our settings if needed
	if cliArgs.OutputConfig != "" {
		// print to stdout if set to '-'
		if cliArgs.OutputConfig == "-" {
			fmt.Printf("### CONFIG:\n%s\n### END CONFIG\n", settings.GetPrettyJSON("", "  "))
		} else {
			fmt.Printf("Saving config file to: %s\n", cliArgs.OutputConfig)
			err := ioutil.WriteFile(cliArgs.OutputConfig, settings.GetJSON(), 0644)
			if err != nil {
				fmt.Printf("Error saving config file: %v\n", err)
			}
		}
	}

	// Init our FS:
	outPath, _ := settings.GetString("gart:output:OutPath", "./output")
	InitFS(outPath)

	// Calculate our output path if our OutFileName isn't defined
	if _, err := settings.GetString("gart:output:OutFileName", ""); err != nil {
		outName, err := GetOutputFileName(outPath, name)
		if err != nil {
			panic(err)
		}

		settings.RawSet(false, "gart:output:OutFileName", outName)
	}

	return settings, cliArgs
}

// InitCLIArgs parses the CLI arguments and returns them. You shouldn't need to
// invoke this directly. It is exposed incase somone needs to initialize the
// program manually.
func InitCLIArgs() *CLIArgs {
	cliArgs := &CLIArgs{}
	flags := pflag.CommandLine
	err := commandeer.Flags(flags, cliArgs)
	if err != nil {
		fmt.Printf("calling Flags: %v\n", err)
	}

	err = flags.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("parsing flags: %v\n", err)
	}

	return cliArgs
}

// InitSettings loads the settings files (reading them from the args if needed)
// You shouldn't need to invoke this directly. It is exposed incase somone
// needs to initialize the program manually.
func InitSettings(cliArgs *CLIArgs) flexiconfig.Settings {
	// Setup our settings
	settings := flexiconfig.NewSettings()
	settings.AddLuaLoader("gart", luaLibLoad)

	// If we don't specify which config files to use just default to config.json/config.lua.
	if len(cliArgs.ConfigPaths) == 0 {
		if err := settings.LoadJSONFile("./config.json"); err != nil {
			println(err.Error())
		}
		if err := settings.LoadLuaFile("./config.lua"); err != nil {
			println(err.Error())
		}
	} else {
		// Load all of the specified config files (ignoring our defaults)
		for _, configPath := range cliArgs.ConfigPaths {
			if err := settings.LoadFile(configPath); err != nil {
				// Since all of these files are manually specified we want to crash if we can't load any
				panic(err)
			}
		}
	}

	return settings
}

func initMetadata(settings flexiconfig.Settings) {
	if ProgramVersionID != "" {
		settings.RawSet(true, "metadata:Version", ProgramVersionID)
	}
}
