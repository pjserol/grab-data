// +build integration

package main

import (
	"os"
	"os/exec"
	"testing"
)

func Test_unzipFile(t *testing.T) {
	pathTmp = "./test"
	os.Mkdir(pathTmp, os.ModePerm)

	file1 := "./test_data/file1.gz"
	file2 := "./test/fileTmp.gz"

	c := exec.Command("cp", file1, file2)

	if err := c.Run(); err != nil {
		t.Errorf("error retrieve file (delete folder): %v", err)
	}

	fileName, err := unzipFile(file2)

	if fileName != "./test/fileTmp" {
		t.Errorf("wrong fileName: %s", fileName)
	}

	if err != nil {
		t.Errorf("error retrieve file: %v", err)
	}

	//time.Sleep(1 * time.Second)

	c = exec.Command("rm", "-R", pathTmp)

	if err := c.Run(); err != nil {
		t.Errorf("error retrieve file (delete folder): %v", err)
	}
}

func Test_processFilesWithTeqc(t *testing.T) {
	file1 := "./test_data/file1"
	bs := "example"
	files := []string{file1}

	err := processFilesWithTeqc(files, bs)
	if err != nil {
		t.Errorf("error process files with teqc: %v", err)
	}

	//time.Sleep(1 * time.Second)

	c := exec.Command("rm", bs+".obs")

	if err := c.Run(); err != nil {
		t.Errorf("error process files with teqc (delete file): %v", err)
	}
}
