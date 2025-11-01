package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// Ensures gofmt doesn't remove the "bytes" import above (feel free to remove this!)
var _ = bytes.ContainsAny

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) 
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) 
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

}



func matchReg(line []byte,pattern string) (bool,error){ 


	if len(pattern) == 0 {
		return true, nil
	}

	if len(line) == 0 {
		return false, nil
	}

	if strings.HasPrefix(pattern,"\\d") {
		if(line[0]>='0' && line[0]<='9'){
			return matchReg(line[1:],pattern[2:])
		}else{
			return false,nil
		}
	}else if strings.HasPrefix(pattern,"\\w"){
		if (line[0]>='a' && line[0]<='z') || (line[0]>='A' && line[0]<='Z') || (line[0]>='0' && line[0]<='9') || line[0]=='_' {
			return matchReg(line[1:],pattern[2:])
		}
	}else if pattern[0] == '['{
		closingIndex := strings.Index(pattern, "]")
		if closingIndex == -1 {
			return false, fmt.Errorf("unclosed character class")
		}
		charClass := pattern[1:closingIndex]
		match := false	
		for i := 0; i < len(charClass); i++ {
			if line[0] == charClass[i] {
				match = true
				break
			}
		}
		if match {
			return matchReg(line[1:], pattern[closingIndex+1:])
		} else {
			return false, nil
		}	
	}
	ok:= pattern[0]==line[0]
	
	if !ok{
		return false,nil
	}
	ok,err:= matchReg(line[1:],pattern[1:])
	

	
	return ok,err
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
