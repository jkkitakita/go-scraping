package config

import (
	"fmt"
	"os"
	"time"
)

var (
	outputDir = os.Getenv("HOME") + "/Downloads"
)

// CsvFilename returns filename
func CsvFilename(name string) string {
	return fmt.Sprintf(`%+v/%+v_%+v.csv`, outputDir, name, time.Now().Format("2006-01-02"))
}
