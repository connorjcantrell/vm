
# vm

## About
Simulating a physical computer by implementing a simple, register-based virtual machine in Go

## The Computer
- 20 bytes of memory
- 3 registers: A pointer counter and 2 general purpose registers
- 5 instructions

### Memory
- Simulated by a fixed sized array to model memory
- 20 byte capacity

```
00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13
__ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __
INSTRUCTIONS ---------------------------^ OUT-^ IN-1^ IN-2^
```

### Instruction
The first 13 bytes are reserved for instructions for executing the program

#### Byte value mapping
```
halt    0x01
load    0x02
store   0x03
add     0x04
sub     0x05
```

#### Paramaters
```
halt 
load reg (addr)     # Load value at given address into register
store reg (addr)    # Store the value in register at the given address
add reg1 reg2       # Add reg1 and reg2 and store result into reg1
sub reg1 reg2       # Subtract reg2 from reg1 and store result into reg1
```

### Registers 
- 1 for the program counter
- 2 general purpose registers


Special thanks to invaluable lessons by Bradfield School of Computer Science


