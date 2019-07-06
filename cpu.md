# CPU

## The model
The traditional model of a computer is a processing unit, which gets told what to do by a program of 
instructions; the program might be accessed with its own special memory, or it might be sitting in the same
area as normal memory, depending on the computer. Each instruction takes a short amount of time to run.
and they're all run one by one. From the CPU's perspective, a loop starts up as soon as the computer is
turned on, to fetch an instruction from memory, work out what it says, and execute it.

In order to keep track of where the CPU is within the program, a number is held by the CPU called the Program
Counter(PC). After an instructin is fetched from memory, the PC is advanced by however many bytes make up the 
instruction.

![The fetch-decode-execute loop](https://github.com/a2-ito/gist_images/blob/master/01_spanner/spanner_figure1.png)

The CPU in the original GameBoy is a modified Zilog Z80m, so the following things are pertinent:

- The Z80 is an 8-bit chip, so all the internal workings operate on one byte at a time;
- The memory interface can address up to 65,5536 bytes (a 16-bit address bus);
- Programs are accessed through the same address bus as normal memory;
- An instruction can be anywhere between one and three bytes.

In addition to the PC, other numbers are held inside the CPU that can be used for calculation, and they're
referred to as registers: A, B, C, D, E, H, and L. Each of them is one byte, so each one can hold a value from 0
to 255. Most of the instructions in the Z80 are used to handle values in these registers: loading a value from
memory into a register, adding or subtracting values, and so forth.

If there are 256 possible values in the first byte of an instruction, that makes for 256 possible instructions in 
the basic table. That table is detailed in the Gameboy Z80 opcode map released on this site. Each of these can 
be simulated by a JavaScript function, that operates on an internal model of the registers, and produces
effects on an internal model of the memory interface.

There are other registers in the Z80. that deal with holding status: the flags register(F), whose operation is 
discussed below; and the stack pointer (SP) which is used alongside the PUSH and POP instructions for basic
LIFO handling of values. The baic model of the Z80 emulation would therefore require the following 
components:

- An internal state:
  - A structure for retaining the current state of the 


