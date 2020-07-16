package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

type fileData struct {
	fileName string
	number   int
}

// retrieveFile download the file from the ftp server
// and copy the file locally in a tmp/ folder
func retrieveFile(f fileData) (string, error) {
	c, err := ftp.Dial("ftp.ngs.noaa.gov:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return "", err
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		return "", err
	}

	res, err := c.Retr(f.fileName)
	if err != nil {
		return "", err
	}

	defer res.Close()

	fileName := fmt.Sprintf(pathTmp+"/file%d.gz", f.number)
	outFile, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, res)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
