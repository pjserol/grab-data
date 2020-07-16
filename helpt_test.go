package main

import (
	"errors"
	"testing"
	"time"
)

func Test_convertTime(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput time.Time
		expectedErr    error
	}{
		{
			name:           "test happy path",
			input:          "2017-09-14T23:11:22Z",
			expectedOutput: time.Date(2017, 9, 14, 23, 11, 22, 0, time.UTC),
		},
		{
			name:        "test with error",
			input:       "",
			expectedErr: errors.New(`parsing time "" as "2006-01-02T15:04:05Z": cannot parse "" as "2006"`),
		},
	}
	for _, test := range tests {
		output, err := convertTime(test.input)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected %v, but got %v", test.name, test.expectedOutput, output)
		}

		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("for %s, expected error %v, but got error %v", test.name, test.expectedErr, err)
		}
	}
}

func Test_getFilePath(t *testing.T) {
	tests := []struct {
		name           string
		baseStationID  string
		hourBlock      string
		in             time.Time
		expectedOutput string
	}{
		{
			name:           "test happy path",
			baseStationID:  "nybp",
			hourBlock:      "b",
			in:             time.Date(2017, 9, 14, 23, 11, 22, 0, time.UTC),
			expectedOutput: "/cors/rinex/2017/257/nybp/nybp257b.17o.gz",
		},
		{
			name:           "test happy path with hour block empty",
			baseStationID:  "nybp",
			in:             time.Date(2017, 9, 14, 23, 11, 22, 0, time.UTC),
			expectedOutput: "/cors/rinex/2017/257/nybp/nybp2570.17o.gz",
		},
		{
			name:           "test with empty value",
			expectedOutput: "/cors/rinex/1/001//0010.1o.gz",
		},
		// case hour block not empty
	}
	for _, test := range tests {
		output := getFilePath(test.baseStationID, test.hourBlock, test.in)

		if test.expectedOutput != output {
			t.Errorf("for %s, expected %v, but got %v", test.name, test.expectedOutput, output)
		}
	}
}
