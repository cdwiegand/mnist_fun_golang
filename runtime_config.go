package main

import (
	"strconv"
	"strings"
)

type RuntimeConfig struct {
	TrainingPath             string
	RunFile                  string
	ModelFile                string
	Mode                     RuntimeMode
	Loops                    int
	RequestedHiddenLayers    []int
	StopIfEveryoneMinQuality int
	VECTOR_MAPPING           map[int]byte
}

type RuntimeMode int

const (
	RuntimeMode_Detect    = 1
	RuntimeMode_Training  = 2
	RuntimeMode_Running   = 3
	RuntimeMode_TrainMore = 4
)

func NewRuntimeConfig(args []string) (ret RuntimeConfig) {
	ret.Loops = 4
	ret.Mode = RuntimeMode_Training
	ret.RequestedHiddenLayers = []int{40, 10}

	ret.TrainingPath = getIfArg(args, "train")
	ret.RunFile = getIfArg(args, "run")
	ret.ModelFile = getIfArg(args, "model")
	Loops := getIfArg(args, "loops")
	if Loops != "" {
		LoopsI, err := strconv.Atoi(Loops)
		if err == nil {
			ret.Loops = LoopsI
		}
	}
	Hiddenlayers := getIfArg(args, "hiddenlayers")
	if Hiddenlayers != "" { // 100,40 or 800,80,14
		parts := strings.Split(Hiddenlayers, ",")
		RequestedHiddenLayers := make([]int, 0)
		for _, part := range parts {
			HiddenlayersI, err := strconv.Atoi(part)
			if err == nil {
				RequestedHiddenLayers = append(RequestedHiddenLayers, HiddenlayersI)
			}
		}
		ret.RequestedHiddenLayers = RequestedHiddenLayers
	}

	/*
		RuntimeMode mode = GetIfArg(args, "mode", p => Enum.TryParse(p, true, out RuntimeMode result) ? result : RuntimeMode.Detect);
		StopIfEveryoneMinQuality = GetIfArg<int?>(args, "stopAtQuality", p => int.TryParse(p, out int result) ? result : null);
		string? hiddenlayers = GetIfArg(args, "hiddenlayers");

		if (mode == RuntimeMode.Detect)
		{
			// can we determine?
			if (!string.IsNullOrEmpty(TrainingPath)) Mode = RuntimeMode.Training;
			else if (!string.IsNullOrEmpty(RunFile)) Mode = RuntimeMode.Running;
			else throw new Exception("No valid --mode: specified, and can't infer usage!");
		}

		if (!string.IsNullOrEmpty(hiddenlayers))
		{
			// means we want a specific size setup
			var tmp = hiddenlayers.Split(',', ':'); // 100:64 as an example or 40,10
			if (!tmp.All(p => int.TryParse(p, out _)))
				throw new Exception("All hidden layer values --hiddenlayers:x,y,z must be ints!");
			RequestedHiddenLayers = tmp.Select(p => int.Parse(p)).ToArray();
		}

		// valid?
		if (string.IsNullOrEmpty(TrainingPath) && string.IsNullOrEmpty(RunFile))
			throw new Exception("Must specify at least a training path or a run file!");
		if (!string.IsNullOrEmpty(RunFile) && (string.IsNullOrEmpty(ModelFile) || !System.IO.File.Exists(ModelFile)))
			throw new Exception("If running, must specify a VALID model file!");
	*/
	ret.VECTOR_MAPPING = make(map[int]byte)

	FILL_LETTERS := "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabdefghnqrt"
	for i := 0; i < len(FILL_LETTERS); i++ {
		ret.VECTOR_MAPPING[i] = FILL_LETTERS[i]
	}
	return
}

func getIfArg(args []string, wanted string) string {
	for i, j := range args {
		if strings.HasPrefix(j, "--"+wanted+":") {
			return j[3+len(wanted):]
		} else if strings.HasPrefix(j, "--"+wanted+"=") {
			return j[3+len(wanted):]
		} else if j == "--"+wanted && i < len(args)-1 {
			return args[i+1]
		}
	}
	return ""
}

// given an index (vector), what is the character?
func (c *RuntimeConfig) DevectorizeKey(key int) byte {
	return c.VECTOR_MAPPING[key]
}

// basicly, turn a character into its index (or vector) for comparison
func (c *RuntimeConfig) VectorizeKey(char byte) int {
	for idx := 0; idx < len(c.VECTOR_MAPPING); idx++ {
		if c.VECTOR_MAPPING[idx] == char {
			return idx
		}
	}
	return -1
}
