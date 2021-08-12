#import <stdio.h>
#include "go/function_pointers.h"

void CLoggerSingle(const char* message, int level)
{
    char *sLevel;
    switch (level) {
        case 5:
            sLevel = "TRACE";
            break;
        case 4:
            sLevel = "DEBUG";
            break;
        case 3:
            sLevel = "INFO";
            break;
        case 2:
            sLevel = "WARNING";
            break;
        case 1:
            sLevel = "ERROR";
            break;
        default:
            sLevel = "UNDEFINED";
            break;
    }
    printf("[%s] %s\n", sLevel, message);
}

int main() {
    loggerFunc loggerSingle = &CLoggerSingle;

    RegisterLogger(loggerSingle);
    LogAllLevels();
}