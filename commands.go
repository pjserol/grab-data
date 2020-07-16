package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// unzipFile run the command gunzip
// return the fileName without the extension
func unzipFile(fileName string) (string, error) {
	log.Println("Command::", append([]string{gunzipCMD}, fileName))
	cmd := exec.Command(gunzipCMD, fileName)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)), nil
}

// processFilesWithTeqc run the command teqc
// and create a file example.obs
func processFilesWithTeqc(fileNames []string, baseStationID string) error {
	cmdArgs := fileNames
	cmdArgs = append(cmdArgs, ">")
	cmdArgs = append(cmdArgs, "example.obs")

	log.Println("Command::", append([]string{teqcCMD}, cmdArgs...))
	cmd := exec.Command(teqcCMD, cmdArgs...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("Error:%v", err)
	}

	return ioutil.WriteFile("./"+baseStationID+".obs", stdout.Bytes(), 0644)
}
