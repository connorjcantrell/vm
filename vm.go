package main

import "fmt"

// Op codes
const LOAD = 0x01
const STORE = 0x02
const ADD = 0x03
const SUB = 0x04
const HALT = 0xff

type VM struct {
	registers [3]uint
	memory    [20]uint
}

// Load 20 byte array of memory into VM
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __
// INSTRUCTIONS ---------------------------^ OUT-^ IN-1^ IN-2^
func (vm *VM) Load(memory [20]uint) {
	vm.memory = memory
}

// Run the stored program to completion,
// mutating the memory array in place.
func (vm *VM) Compute() {
	for {
		pc := vm.registers[0]
		op_code := vm.memory[pc]

		if op_code == HALT {
			return
		}

		arg1 := vm.memory[pc+1]
		arg2 := vm.memory[pc+2]

		switch op_code {
		case LOAD:
			/*
				Build a mask to select the lower 8 bits:
				vm.memory[arg2+1]:     0000 0000 0000 0001 : 1
				vm.memory[arg2+1]<<8 : 0000 0001 0000 0000 : 256

				bitwise "|" this and another bit string

				vm.memory[arg2+1]<<8 : 0000 0001 0000 0000
				     vm.memory[arg2] : 0000 0000 XXXX XXXX
				==========================================
							       | : 0000 0001 XXXX XXXX
			*/
			vm.registers[arg1] = vm.memory[arg2+1]<<8 | vm.memory[arg2]
		case STORE:
			/*
				Build a mask to select the lower 8 bits:
				 00000001 : 1
				100000000 : 1 << 8
				 11111111 : (1 << 8) - 1

				You can then bitwise "&" this and another bit string

				vm.memory[arg2] : XXXX XXXX XXXX XXXX
				   (1 << 8) - 1 : 0000 0000 1111 1111
				======================================
							  & : 0000 0000 XXXX XXXX
			*/
			vm.memory[arg2] = vm.registers[arg1] & ((1 << 8) - 1)
			/*
				// Higher significant byte
				vm.memory[arg2+1] = vm.registers[arg1] >> 8
				// same as
				vm.memory[arg2+1] = vm.registers[arg1] // 256;
				// same as
				vm.memory[arg2+1] = vm.registers[arg1] / 2**8;
				// same as
				vm.memory[arg2+1] = vm.registers[arg1] / (2 / 2 / 2 / 2 / 2 / 2 / 2 / 2);

				// Each shift right ( >> ) is a divide by 2
				// Each shift left ( << ) is a multiply by 2

				Remove mask to select the lower 8 bits:
				vm.memory[arg2+1]    : 0000 0001 0000 0000 : 256
				vm.memory[arg2+1]>>8 : 0000 0000 0000 0001 : 1
			*/
			vm.memory[arg2+1] = vm.registers[arg1] >> 8
		case ADD:
			vm.registers[arg1] = vm.registers[arg1] + vm.registers[arg2]
		case SUB:
			vm.registers[arg1] = vm.registers[arg1] - vm.registers[arg2]

		}
		// update program counter
		vm.registers[0] += 3
	}
}

func (vm *VM) String(message string) {
	fmt.Println(message)
	for _, v := range vm.memory {
		fmt.Printf("%02x ", v)
	}
	fmt.Println("\nINSTRUCTIONS ---------------------------^ OUT-^ IN-1^ IN-2^\n")
}

func main() {
	var memory = [20]uint{
		0x01, 0x01, 0x10, // 0x00: load A 0x10
		0x01, 0x02, 0x12, // 0x03: load B 0x12
		0x03, 0x01, 0x02, // 0x06: add A B
		0x02, 0x01, 0x0e, // 0x09: store A 0x0e
		0xff,       // 0x0c: halt
		0x00,       // <<unused>>
		0x00, 0x00, // 0x0e: output
		0xff, 0x00, // 0x10: input X = 255
		0x03, 0x00, // 0x12: input Y = 3
	}

	vm := VM{}
	vm.Load(memory)
	vm.String("Before:")
	vm.Compute()
	vm.String("After:")
}
