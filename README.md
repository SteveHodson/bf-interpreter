# bf-interpreter
I'm learning Golang so I thought I'd write an interpreter for a simple language as a starting project and BrainFuck being a simple Turing complete language makes an excellent choice.
[BrainFuck on Wikipeadia](https://en.wikipedia.org/wiki/Brainfuck)
I am updating the project in bits adding more and more language instructions.  The included hello_world.b file is for future testing as the interpreter does not yet support all the instructions used in that program.
## Usage
Make sure you have Golang installed and on your path
```
go run main.go hello_world.b
```

## Instructions Supported
| Instruction | Command                                             |
|:------------|:----------------------------------------------------|
|    +        | Increment value at a memory location by one         |
|    -        | decrement value at a memory location by one         |
|    <        | Move the memory location to the left by one         |
|    >        | Move the memory location to the right by one        |
|    [        | Jmp Fwd to associated ] if mem location == 0        |
|    ]        | Jmp Back to associated [ if mem location != 0       |
|    .        | Output the current memory location                  |
## Instructions Not Currently Supported
| Instruction | Command                                             |
|:------------|:----------------------------------------------------|
|    ,        | Input value to current memory location              |
