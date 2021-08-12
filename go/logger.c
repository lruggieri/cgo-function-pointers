typedef void (*loggerFunc) (const char* message, int level);

// this is the trick that allows us to call a C function from Go. Since 'loggerFunc' cannot be called as a function in Go,
// we just call it from C. In Go, we can instead call C.bridge_logger.
void bridge_logger(loggerFunc f, const char* message, int level)
{
    f(message, level);
}