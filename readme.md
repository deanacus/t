# t

t - the minimalist command line todo list

This is still a work in progress

Features:

- [x] List todo items  
- [x] Add todo items  
- [x] Delete todo items  
- [x] Save todo items in a file  
- [ ] Configure using a config file  
- [ ] Have a default config  
- [x] Colourise output

## Prerequisites

* [Go](https://golang.org)

## Installation

Clone the repo to your `$GOPATH`, then you can choose between running the app
from source or compiling the binary

#### Run from source

```go run t```

#### Compile

```go build t```

then move the binary to somewhere in your `$PATH`

## Usage

```
Usage: t [option] <Input>

Options:
  -l, --list                    List the tasks in the todo list
  -a, --add <Task>              Add a task to the todo list
  -d, --delete <TaskID>         Delete a task from the todo list
  -v, --version                 Show the version
  -h, --help                    Show this help information
```

## Contributing

Please read [CONTRIBUTING.md][cont] for details on how to contribute to this project, 
including coding standards, commit message requirements and the process for 
submitting pull requests.

## Versioning

SemVer all the things

[cont]: https://github.com/deanacus/t/tree/master/CONTRIBUTING.md