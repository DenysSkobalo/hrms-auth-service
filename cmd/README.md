# cmd/
The entry point for our application. The directory name for each application must match the name of the executable file you want to build. You shouldn't put a lot of code in this directory. The most common practice is to use a small main function that imports and calls all the necessary code from the /internal and /pkg directories.