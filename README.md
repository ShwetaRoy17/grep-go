# go-grep

A lightweight, fast, and minimal implementation of the classic `grep` command-line utility written in **Go**.  
It searches for patterns in text files or standard input using regular expressions.

## Features

- Supports basic and extended regular expressions  
- Matches patterns across multiple files  
- Case-sensitive and case-insensitive search options  
- Prints matching lines with file names and line numbers  
- Handles piped input (e.g., from `cat`, `echo`, etc.)

## Installation

```bash
git clone https://github.com/ShwetaRoy17/grep-go.git
cd grep-go
go build -o grep-go
