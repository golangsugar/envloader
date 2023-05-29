package envloader

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	configRegex  = `^([a-zA-Z][a-zA-Z0-9_]+)=([[\S ]*]?)$`
	rxGroupKey   = 1
	rxGroupValue = 2
)

var rxConfig = regexp.MustCompile(configRegex)
var logOnFileCloseError = false

// LogFileClosingError enables logging of errors when closing the file.
func LogFileClosingError() {
	logOnFileCloseError = true
}

func fileClose(f *os.File) {
	if err := f.Close(); err != nil && logOnFileCloseError {
		log.Printf("error closing file %v", err)
	}
}

func readLine(line string) (key, value string, valid bool) {
	if line == "" || strings.HasPrefix(line, "#") { // line is commented or empty
		return
	}

	if matches := rxConfig.FindStringSubmatch(line); len(matches) > 0 {
		valid = true
		key = matches[rxGroupKey]

		if len(matches) > 2 {
			value = matches[rxGroupValue]
		}
	}

	return
}

// LoadFromFile loads environment variables values from a given text file in to a map[string]string.
// configFile is the file name, with the complete path if necessary.
// if errorIfFileDoesntExist, the function returns with an error in case of the given file doesn't exist.
// valid lines must comply with regex ^([A-Z][A-Z0-9_]+)([=]{1})([[\S ]*]?)$.
// Examples of valid lines:
// ABC=prd
// XYZ=
// ABC="42378462%&&3 178964@"
// mnoPQR=42378462%&&3 ###
//
// Examples of *invalid* lines:
// Commented/ignored: #XYZ=4334343434 ( starts with # ).
// invalid/ignored: opt=ler0ler0 ( has to be all caps/uppercase ).
// Invalid/Ignored: _LETTERS=4334343434 ( has to start with a letter ).
// Invalid/Ignored: X=4334343434 ( should contain 2 or more chars ).
// Environment variables reference for curious: https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap08.html.
func LoadFromFile(configFile string, errorIfFileDoesntExist bool) error {
	f, err := os.Open(configFile)
	if err != nil {
		if err == os.ErrNotExist && !errorIfFileDoesntExist {
			return nil
		}

		return err
	}

	defer fileClose(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if key, value, valid := readLine(line); valid {
			if err = os.Setenv(key, value); err != nil {
				return err
			}
		}
	}

	return nil
}
