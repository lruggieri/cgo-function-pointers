package main

/*
#include <string.h>
#include <stdlib.h> // for 'free' function

typedef void (*loggerFunc) (const char* message, int level);
void bridge_logger(loggerFunc f, const char* message, int level);

*/
import "C"
import "unsafe"

var loggingFunctionPointer C.loggerFunc

func main() {}

//export RegisterLogger
func RegisterLogger(iFunctionPointer C.loggerFunc) {
	loggingFunctionPointer = iFunctionPointer
}

//export LogAllLevels
func LogAllLevels(){
	if loggingFunctionPointer != nil {
		logWrapper := func(iLoggerFp C.loggerFunc, iLevel C.int) func(format string) {
			return func(format string) {
				message := C.CString(format)

				// this actually calls the registered C function pointer and logs
				C.bridge_logger(iLoggerFp, message, iLevel)

				C.free(unsafe.Pointer(message))
			}
		}

		logger := NewLogger(
			logWrapper(loggingFunctionPointer, 5),
			logWrapper(loggingFunctionPointer, 4),
			logWrapper(loggingFunctionPointer, 3),
			logWrapper(loggingFunctionPointer, 2),
			logWrapper(loggingFunctionPointer, 1),
		)

		logger.Trace("This is a TRACE message")
		logger.Debug("This is a DEBUG message")
		logger.Info("This is a INFO message")
		logger.Warning("This is a WARNING message")
		logger.Error("This is a ERROR message")
	}
}
