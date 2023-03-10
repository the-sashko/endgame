package logger

import (
	"endgame/src/settings"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const TypeDefault = "info"
const TypeDebug = "debug"
const TypeError = "error"

const logDirPath = "../logs"

func LogMessage(text string, logType string) {
	logFilePath := getLogFilePath(logType)
	oldLogFilePath := getOldLogFilePath(logType)

	currTime := time.Now().Format("[2006-01-02 15:04:05]")
	text = fmt.Sprintf("%s %s\n", currTime, text)

	logDir := fmt.Sprintf("%s/%s", logDirPath, logType)

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)

		if err != nil {
			doError(err)
		}
	}

	if _, err := os.Stat(oldLogFilePath); !os.IsNotExist(err) {
		err := os.Remove(oldLogFilePath)

		if err != nil {
			doError(err)
		}
	}

	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		_, err := os.Create(logFilePath)

		if err != nil {
			doError(err)
		}
	}

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := logFile.Close(); err != nil {
			if err != nil {
				doError(err)
			}
		}
	}()

	if _, err := logFile.WriteString(text); err != nil {
		if err != nil {
			doError(err)
		}
	}

	text = fmt.Sprintf("[%s] %s", strings.ToUpper(logType), text)

	print(text)
}

func LogDebug(debugMessage string) {
	if !settings.GetSettings().IsDebug() {
		return
	}

	LogMessage(debugMessage, TypeDebug)
}

func LogError(errorType string, errorEntity error, isExit bool) {
	errorMessage := fmt.Sprintf("[%s error] %s", errorType, errorEntity.Error())

	LogMessage(errorMessage, TypeError)

	if isExit {
		os.Exit(0)
	}
}

func getLogFilePath(logType string) string {
	currDate := time.Now().Format("2006-01-02")

	return fmt.Sprintf("%s/%s/%s.log", logDirPath, logType, currDate)
}

func getOldLogFilePath(logType string) string {
	currentYear := time.Now().Format("2006")

	previousYear, _ := strconv.Atoi(currentYear)
	previousYear--

	previousYearString := strconv.Itoa(previousYear)

	oldDate := previousYearString + time.Now().Format("-01-02")

	return fmt.Sprintf("%s/%s/%s.log", logDirPath, logType, oldDate)
}

func doError(errorEntity error) {
	errorMessage := fmt.Sprintf("[Logger Error]: %s", errorEntity.Error())
	fmt.Println(errorMessage)
	os.Exit(1)
}
