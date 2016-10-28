// Copyright 2016 Marcel Edmund Franke. All rights reserved.
// license that can be found in the LICENSE file.

package command_pipeline_test

import (
	"bytes"
	"github.com/donutloop/command-pipeline"
	"testing"
)

func TestCommandPipeline_Execute(t *testing.T) {

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

	expected := "HELLO WORLD"

	if err != nil {
		t.Errorf("Expected \"%s\" got \"%v\"", expected, err)
	}

	if data.String() != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, data.String())
	}
}

func TestCommandPipeline_Add(t *testing.T) {

	buildCommand := func(letter string) func(input *bytes.Buffer) (*bytes.Buffer, error) {
		return func(input *bytes.Buffer) (*bytes.Buffer, error) {

			input.WriteString(letter)

			return input, nil
		}
	}

	firstInput := bytes.NewBufferString("H")

	pipeline := command_pipeline.New(firstInput)

	pipeline.Add(
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

	expected := "HELLO WORLD"

	if err != nil {
		t.Errorf("Expected \"%s\" got \"%v\"", expected, err)
	}

	if data.String() != expected {
		t.Errorf("Expected \"%s\" got \"%s\"", expected, data.String())
	}
}

func TestCommandPipeline_AddedToLessCommands(t *testing.T) {

	firstInput := bytes.NewBufferString("H")

	pipeline := command_pipeline.New(firstInput)

	_, err := pipeline.Execute()

	expected := "Added to less commands (Count: 0)"

	if err == nil || err.Error() != expected {
		t.Errorf("Expected \"%s\" got \"%v\"", expected, err)
	}
}

func TestCommandPipeline_InputIsMissing(t *testing.T) {

	pipeline := command_pipeline.New(nil)

	_, err := pipeline.Execute()

	expected := "Input is missing (Value: <nil>)"

	if err == nil || err.Error() != expected {
		t.Errorf("Expected \"%s\" got \"%v\"", expected, err)
	}
}

func BenchmarkCommandPipeline_Execute(b *testing.B) {

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

	for n := 0; n < b.N; n++ {
		pipeline.Execute()
	}
}
