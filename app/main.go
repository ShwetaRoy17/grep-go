package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
	"regexp"
)

// Ensures gofmt doesn't remove the "bytes" import above (feel free to remove this!)
var _ = bytes.ContainsAny

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchReg(line []byte,pattern string) (bool,error){ 

	match,err :=regexp.Match(pattern,line)
	// if(len(pattern)==0){
	// 	return true
	// }
	
	// if(ind>=len(line)){
	// 	return len(pattern)==0
	// }



	// if(strings.HasPrefix(pattern,"\\d")){
	// 	if	line[ind]>='0' && line[ind]<='9'{
	// 		return matchReg(line,ind+1,pattern[0:len(pattern)])
	// 	}else{
	// 		return matchReg(line,ind+1,pattern)
	// 	}
	// }else if(strings.HasPrefix(pattern,"\\w")){
	// 	if (line[ind]>='a' && line[ind]<='z') || (line[ind]>='A' && line[ind]<='Z') || (line[ind]>='0' && line[ind]<='9') || (line[ind]=='_'){
	// 		return matchReg(line,ind+1,pattern[0:len(pattern)])
	// 	}else{
	// 		return matchReg(line,ind+1,pattern)
	// 	}
	// }
	


	return match,err
}

func matchLine(line []byte, pattern string) (bool, error) {
	var ok bool
	if utf8.RuneCountInString(pattern)==1 {
		ok = bytes.ContainsAny(line,pattern)
	} 
	ok, err:= matchReg(line,pattern)
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println(ok,err)
	return ok, err
}
