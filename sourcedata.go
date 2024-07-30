package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"image/png"
	"os"
	gopath "path"
	"path/filepath"
	"strconv"
	"strings"
)

type SourceData struct {
	List []DataItem
}

func (s *SourceData) AddSourceData(item DataItem) {
	s.List = append(s.List, item)
}
func (s *SourceData) AddSourceDatas(items []DataItem) {
	s.List = append(s.List, items...)
}

func (s *SourceData) LoadTrainingSource(path string, config RuntimeConfig) {
	stat, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Unable to load file %s", path)
		panic(err)
	}
	if stat.IsDir() {
		fmt.Printf("Processing directory %s\n", path)
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			s.LoadTrainingSource(gopath.Join(path, f.Name()), config) // recurse
		}
	} else {
		fmt.Printf("Processing file %s\n", path)
		parts := strings.Split(path, ".")
		ext := parts[len(parts)-1]
		if ext == "csv" {
			s.LoadCsvFile(path, config)
		} else if ext == "png" {
			cs := filepath.Dir(path)
			if len(cs) == 1 {
				c := cs[len(cs)-1]
				s.LoadTrainingPngImage(c, path, config)
			}
		} else {
			fmt.Printf("W: Ignoring file %s\n", path)
		}
	}
}

/*
func (s *SourceData)  LoadRunningSource(path string, config RuntimeConfig)

	{
	    SourceData sourceData = appendTo ?? new SourceData();

	    if (System.IO.Directory.Exists(path))
	    {
	        foreach (var file in System.IO.Directory.GetDirectories(path))
	            LoadRunningSource(file, config, sourceData); // recurse
	        foreach (var file in System.IO.Directory.GetFiles(path))
	            LoadRunningSource(file, config, sourceData);
	    }
	    else if (System.IO.File.Exists(path))
	    {
	        switch (System.IO.Path.GetExtension(path).Trim().ToLower().Replace(".", ""))
	        {
	            case "csv": sourceData.AddRange(LoadCsvFile(path, config)); break;
	            case "png": sourceData.Add(LoadRuntimePngImage(path)); break;
	            default: Console.WriteLine("W: Ignoring file " + path); break;
	        }
	    }

	    return sourceData;
	}
*/
func (s *SourceData) LoadCsvFile(path string, config RuntimeConfig) {
	// path to mnist csv or emnist csv

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, data := range records {
		fieldMap, err := strconv.Atoi(data[0])
		if err != nil {
			panic(err)
		}
		pixelsInt := make([]int, len(data)-1)
		pixels := make([]float64, len(data))

		for i := 1; i < len(data); i++ {
			pixelsInt[i-1], _ = strconv.Atoi(data[i])
			pixels[i-1] = float64(pixelsInt[i-1] / 256)
		}

		di := DataItem{
			Character: config.DevectorizeKey(fieldMap),
			Pixels:    pixels,
		}
		s.AddSourceData(di)
	}
}

func (s *SourceData) LoadTrainingPngImage(c byte, path string, config RuntimeConfig) DataItem {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	image, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	idx := 0
	pixels := make([]float64, image.Bounds().Size().X*image.Bounds().Size().Y)
	for y := 0; y < image.Bounds().Size().Y; y++ {
		for x := 0; x < image.Bounds().Size().X; x++ {
			pixel := color.GrayModel.Convert(image.At(x, y)).(color.Gray)
			pixels[idx] = float64(pixel.Y) / 256 // assign and more to next pixel
			idx++
		}
	}
	return DataItem{
		Character: c,
		Pixels:    pixels,
	}
}
func LoadRuntimePngImage(path string) DataItem {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	image, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	idx := 0
	pixels := make([]float64, image.Bounds().Size().X*image.Bounds().Size().Y)
	for y := 0; y < image.Bounds().Size().Y; y++ {
		for x := 0; x < image.Bounds().Size().X; x++ {
			pixel := color.GrayModel.Convert(image.At(x, y)).(color.Gray)
			pixels[idx] = float64(pixel.Y) / 256 // assign and more to next pixel
			idx++
		}
	}
	return DataItem{
		Pixels: pixels,
	}
}
