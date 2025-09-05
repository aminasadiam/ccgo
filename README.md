# CCGo Compiler

A simple C compiler written in Go that translates a subset of C to x86-64 assembly. This is an educational project to learn compiler design.

## Features

### Supported

- Function definitions (`int main()`)
- Return statements with simple expressions
- Integer literals
- Basic arithmetic (`+`, `-`)
- Block statements `{}`

### Planned

- More operators (`*`, `/`)
- Variable declarations
- Control flow (`if`, `while`)
- Function parameters

## Architecture

1. **Lexical Analysis** (`lexer/lexer.go`): Tokenizes source code.
2. **Parsing** (`parser/parser.go`): Builds AST using recursive descent.
3. **Code Generation** (`codegen/codegen.go`): Outputs x86-64 assembly.

## Usage

### Prerequisites

- Go compiler
- GNU Assembler (`as`)
- GNU Linker (`ld`)

### Build

```bash
go build -o ccgo main.go
```

### Compile a C file

```bash
./ccgo source.c
```

### Run

```bash
as -o out.o out.asm
ld -o out out.o
./out
echo $?  # Check return value
```

### Example

```c
int main() {
    return (2 + 3);
}
```

Run:

```bash
./ccgo test.c
as -o test.o out.asm
ld -o test test.o
./test
echo $?  # Outputs: 5
```

## Contributing

This is a learning project. Contributions welcome:

- Improve error messages
- Add expression types
- Write tests
- Expand documentation

## License

MIT
