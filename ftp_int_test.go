// +build integration

package main

import (
	"os"
	"os/exec"
	"testing"
)

func Test_retrieveFile(t *testing.T) {
	pathTmp = "./test"
	os.Mkdir(pathTmp, os.ModePerm)

	fileName, err := retrieveFile(fileData{
		fileName: "/cors/rinex/2020/187/nypb/nypb1870.20o.gz",
		number:   1,
	})

	if fileName != "./test/file1.gz" {
		t.Errorf("wrong fileName: %s", fileName)
	}

	if err != nil {
		t.Errorf("error retrieve file: %v", err)
	}

	c := exec.Command("rm", "-R", pathTmp)

	if err := c.Run(); err != nil {
		t.Errorf("error retrieve file (delete folder): %v", err)
	}
}
