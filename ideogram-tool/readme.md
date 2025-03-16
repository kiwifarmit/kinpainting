# Use the default config.json

ideogram.exe

# Specify a configuration file

ideogram.exe --config different\_config.json

# To compile the Go program on Windows:

Install Go from [https://golang.org/dl/](https://www.google.com/url?sa=E&source=gmail&q=https://golang.org/dl/) if not already installed.

Open a command prompt in the project directory.

Compile the program:

```
go build -o ideogram.exe
```

The ideogram.exe file will be created in the current directory and can be run from any location as long as the config.json file is in the same directory as the executable.