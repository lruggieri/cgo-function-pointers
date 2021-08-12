### Medium article
This repo is a working example of [this Medium article](https://medium.com/@luc4.ruggieri/cgo-calling-a-c-function-pointer-from-go-dab4dd0fc9e2).

### How to compile
```
cd go
go build — buildmode=c-shared ./

cd ..
gcc -Wall -g main.c -o main go/function_pointers
```
and then run

`DYLD_LIBRARY_PATH="$(pwd)/go" ./main` on Mac

` LD_LIBRARY_PATH="$(pwd)/go" ./main` on Linux

Or, build a proper CMake file :)