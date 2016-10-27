package command_pipeline

import (
	"bytes"
	"fmt"
)

type Command func(*bytes.Buffer) (*bytes.Buffer, error)

type CommandPipeline struct {
	commands    []Command
	input       *bytes.Buffer
}

func New(input *bytes.Buffer, command ...Command) *CommandPipeline {
	return &CommandPipeline{
		commands: append([]Command{}, command...),
		input:    input,
	}
}

func (cp *CommandPipeline) Execute() (*bytes.Buffer, error) {

	if cp.input == nil {
		return nil, fmt.Errorf("Input is missing (Value: %v)", cp.input)
	}

	if len(cp.commands) < 2 {
		return nil, fmt.Errorf("Added to less commands (Count: %d)", len(cp.commands))
	}

	errChan := make(chan error)

	connectoren := cp.createConnectors(len(cp.commands) + 1)

	for index, command := range cp.commands {
		go cp.convertToCommandWrapper(commandWrapperParam{command, connectoren[index], connectoren[index+1], errChan})()
	}

	go cp.initialStart(connectoren[0])()

	var err error
	var output *bytes.Buffer

	select {
	case err = <-errChan:
	case output = <- connectoren[len(cp.commands)]:
	}

	cp.closeConnectors(connectoren)

	return output, err
}

func (cp *CommandPipeline) createConnectors(commandCount int) []chan *bytes.Buffer {

	connectoren := make([]chan *bytes.Buffer, commandCount)

	for index := 0; index < commandCount; index++ {
		connectoren[index] = make(chan *bytes.Buffer)
	}

	return connectoren
}

func (cp *CommandPipeline) closeConnectors(connectoren []chan *bytes.Buffer) {
	for _, connector := range connectoren {
		close(connector)
	}
}

func (cp *CommandPipeline) initialStart(connector chan *bytes.Buffer) func() {
	return func() {
		connector <- cp.input
	}
}

func (cp *CommandPipeline) Add(c ...Command) {
	cp.commands = append(cp.commands, c...)
}

func (cp *CommandPipeline) Clear() {
	cp.commands = []Command{}
}

type commandWrapperParam struct {
	command Command
	input   chan *bytes.Buffer
	output  chan *bytes.Buffer
	err     chan error
}

func (cp *CommandPipeline) convertToCommandWrapper(param commandWrapperParam) func() {
	return func() {
		input := <-param.input

		if output, err := param.command(input); err != nil {
			param.err <- err
		} else {
			param.output <- output
		}
	}
}
