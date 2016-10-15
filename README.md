[![Build Status](https://travis-ci.org/donutloop/command-pipeline.svg?branch=master)](https://travis-ci.org/donutloop/command-pipeline)

# Command pipeline

CommandPipeline is a lightweight high performance pipeline (using "concurrency" to execute the commands) for Go.

## Usage

This is just a quick introduction

Let's start with a trivial Hello World example:

```go
    package main

    import (
        "github.com/donutloop/command-provider"
        "bytes"
    )

    func main() {

	    buildCommand := func(letter string) func(input *bytes.Buffer) (*bytes.Buffer, error) {
        		return func(input *bytes.Buffer) (*bytes.Buffer, error) {
        
        			input.WriteString(letter)
        
        			return input, nil
        		}
        	}
        
        	firstInput := bytes.NewBufferString("H")
        
        	pipeline := command_pipeline.New(
        		firstInput,
        		buildCommand("E"),
        		buildCommand("L"),
        		buildCommand("L"),
        		buildCommand("O"),
        		buildCommand(" "),
        		buildCommand("W"),
        		buildCommand("O"),
        		buildCommand("R"),
        		buildCommand("L"),
        		buildCommand("D"),
        	)
        
        	data, err := pipeline.Execute()
        
        	data.String() //Output: HELLO WORLD
    }
```