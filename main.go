package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

const (
	nbConsumers = 3
)

var gunzipCMD string
var teqcCMD string
var pathTmp = "./tmp"

func init() {
	gunzipCMD = "gunzip"
	if os.Getenv("GUNZIP_CMD") != "" {
		gunzipCMD = os.Getenv("GUNZIP_CMD")
	}

	teqcCMD = "teqc"
	if os.Getenv("TEQC_CMD") != "" {
		teqcCMD = os.Getenv("TEQC_CMD")
	}
}

func main() {
	if len(os.Args) != 4 {
		panic("Please specify the base station id, start time and end time.")
	}

	baseStationID := os.Args[1]
	startTimetStr := os.Args[2]
	endTimeStr := os.Args[3]

	log.Printf("Input:%s %s %s", baseStationID, startTimetStr, endTimeStr)

	start, err := convertTime(startTimetStr)
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

	end, err := convertTime(endTimeStr)
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

	if _, err := os.Stat(pathTmp); os.IsNotExist(err) {
		os.Mkdir(pathTmp, os.ModePerm)
	}

	var ch = make(chan fileData)
	wg := sync.WaitGroup{}
	wg.Add(nbConsumers)

	var files []string
	for i := 0; i < nbConsumers; i++ {
		go func(n int) {
			for data := range ch {
				file, err := processFile(data)
				if err != nil {
					log.Printf("error process file (%s): %s", data.fileName, err)
				} else {
					files = append(files, file)
				}
			}
			wg.Done()
		}(i)
	}

	var i int
	for start.Unix() < end.Unix() {
		path := getFilePath(baseStationID, "", start)
		log.Printf("start fetching file: %s", path)
		ff := fileData{
			fileName: path,
			number:   i,
		}
		ch <- ff

		start = start.AddDate(0, 0, 1)
		i++
	}

	close(ch)
	wg.Wait()

	if err := processFilesWithTeqc(files, baseStationID); err != nil {
		log.Fatal(err)
	}

	log.Println("file example.obs successfully created!")

	clean()
}

// processFile retrieve the file from the ftp server
// and unzip the file
func processFile(f fileData) (string, error) {
	fileName, err := retrieveFile(f)
	if err != nil {
		return "", err
	}

	return unzipFile(fileName)
}

// clean remove folder
func clean() {
	c := exec.Command("rm", "-R", pathTmp)

	if err := c.Run(); err != nil {
		log.Println(err)
	}
}
