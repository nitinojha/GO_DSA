//Explanation:
//Subsystems (CPU, Memory, HardDrive): Represents the complex components of a computer system with specific responsibilities.
//Facade (Computer): Provides a Start method that internally manages and calls the operations of these subsystems without exposing them to the client. The Facade simplifies the usage of these complex components by offering a unified interface.
//Client Code: Interacts with the complex system through the Computer facade, simplifying the boot process with a single high-level method call.

package main

import "fmt"

// CPU is a subsystem for processing tasks
type CPU struct{}

// Freeze freezes the CPU
func (c *CPU) Freeze() {
	fmt.Println("Freezing CPU.")
}

// Jump jumps to a specific position
func (c *CPU) Jump(position int) {
	fmt.Printf("Jumping to position %d.\n", position)
}

// Execute executes commands
func (c *CPU) Execute() {
	fmt.Println("Executing instructions.")
}

// Memory is a subsystem for data storage
type Memory struct{}

// Load loads data to the memory
func (m *Memory) Load(position int, data string) {
	fmt.Printf("Loading %s to position %d.\n", data, position)
}

// HardDrive is a subsystem for permanent data storage
type HardDrive struct{}

// Read reads data from the hard drive
func (h *HardDrive) Read(position int, size int) string {
	fmt.Printf("Reading %d bytes from position %d.\n", size, position)
	return "BootLoader"
}

// Computer is the Facade providing a simple interface to complex subsystems
type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

// NewComputer initializes the subsystems
func NewComputer() *Computer {
	return &Computer{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// Start starts the computer; the Facade method
func (c *Computer) Start() {
	c.cpu.Freeze()
	c.memory.Load(0, c.hardDrive.Read(0, 256))
	c.cpu.Jump(0)
	c.cpu.Execute()
}

func main() {
	computer := NewComputer()
	computer.Start()
}