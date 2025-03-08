package mrlog

import (
	"log"
	"os"
	"time"
)

// getLogFileName mengembalikan nama file log berdasarkan tipe log (info atau error)
func getLogFileName(logType string) string {
	return time.Now().Format("2006-01-02") + "_" + logType + ".log"
}

// openLogFile membuka atau membuat file log
func openLogFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

// WriteInfo menulis pesan log ke file log info
func Info(message string) {
	writeLog("info", message)
}

// WriteError menulis pesan log ke file log error
func Error(message string) {
	writeLog("error", message)
}

// writeLog adalah fungsi umum untuk menulis log berdasarkan jenisnya
func writeLog(logType, message string) {
	logFileName := getLogFileName(logType)
	logFile, err := openLogFile(logFileName)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Println(message)
}
