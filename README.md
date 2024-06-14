# Brainfuck Interpreter

This project provides an interpreter for the Brainfuck language. It reads Brainfuck source code from the first parameter and executes it.

## Commands

Brainfuck consists of the following commands:

* `>`: Increment the memory pointer.
* `<`: Decrement the memory pointer.
* `+`: Increment the byte at the memory pointer.
* `-`: Decrement the byte at the memory pointer.
* `.`: Output the byte at the memory pointer as an ASCII character.
* `[`: Jump forward to the command after the corresponding `]` if the byte at the memory pointer is zero.
* `]:` Jump back to the command after the corresponding `[` if the byte at the memory pointer is non-zero.
* Any other character is ignored and treated as a comment.

## Memory Model

* The interpreter uses an array of 2048 bytes, all initialized to 0, with a pointer starting at the first byte.
Usage

* To run the Brainfuck interpreter, use the following command line syntax:

```bash
$ go run . "Your Brainfuck code here"
```

### Examples

1. Hello World Program:

```bash
$ go run . "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
Hello World!$
```

2. Short Program Outputting 'Hi':

```bash
$ go run . "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>."
Hi$
```

3. Program Outputting 'abc':

```bash
$ go run . "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++."
abc$
```

4. Empty Input:
```bash
$ go run .
$
```