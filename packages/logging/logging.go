package log

import (
    "fmt"
    "time"
    "os"
)

var logLevel = "INFO"

func suppressLog(level string) bool {
    if ( logLevel != "VERBOSE" && level != "INFO" && level != "ERROR" ) {
        return true
    } 
    return false
}

// filter log level by 

// export printLog function
func PrintLog(data string, level string) {
    if suppressLog(level) {
        return
    }

    time := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s][%s] %s\n", level, time, data)
}

func Error(data string) {
    PrintLog(data, "ERROR")
}

func Warning(data string) {
    PrintLog(data, "WARNING")
}

func Info(data string) {
    PrintLog(data, "INFO")
}

func Debug(data string) {
    PrintLog(data, "DEBUG")
}

func Fatal(err error) {
    // print then exit
    PrintLog(err.Error(), "FATAL")
    os.Exit(1)
}

func EnableVerbose() {
    logLevel = "VERBOSE"
}