# CPU

## The model
The traditional model of a computer is a processing unit, which gets told what to do by a program of 
instructions; the program might be accessed with its own special memory, or it might be sitting in the same
area as normal memory, depending on the computer. Each instruction takes a short amount of time to run.
and they're all run one by one. From the CPU's perspective, a loop starts up as soon as the computer is
turned on, to fetch an instruction from memory, work out what it says, and execute it.

従来のコンピュータモデルは、プロセッシングユニットである。それは、命令コードによって何をすべきか指示を得る。
そのプログラムは、独自のメモリによってアクセスされるか、通常のメモリ領域に書かれているだろう。
各命令は、短時間で、一つ一つ実行される。CPUの観点では、コンピュータが起動するとすぐにループが始まり、
メモリから命令をフェッチし、それに実行する。

In order to keep track of where the CPU is within the program, a number is held by the CPU called the Program
Counter(PC). After an instructin is fetched from memory, the PC is advanced by however many bytes make up the 
instruction.

CPUがプログラムのどこにいるのかをトラックするために、Program Counter (PC) という数値が CPU によって保持される。
命令がメモリからフェッチされた後、PC は命令を構成する多くのバイト列によってインクリメントされる。

![The fetch-decode-execute loop](https://github.com/a2-ito/gist_images/blob/master/02_go-gameboy/01_the_fetch-decode-execute_loop.pngh)

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
  - A structure for retaining the current state of the registers;
  - The amount of time used to execute the last instruction;
  - The amount of time that the CPU has run in total;
- Functions to simulate each instruction;
- A table mapping said functions onto the opcode map;
- A known interface to talk to the simulated memory.

The flags register(F) is important to the functioning of the processor: it automatically calculates certain
bits, or flags, based on the result of the last operation. There are four flags in the Gameboy Z80:

- Zero(0x80): Set if the last operation produced a result of 0;
- Operation(0x40): Set if the last operation was a subtraction;
- Half-carry(0x20): Set if, in the result of the last operation, the lower half of the byte overflowed past 15;
- Carry(0x10): Set if the last operation produced a result over 255 (for additions) or under 0 (for subtractions).

Since the basic calculation registers are 8-bits, the carry flag allows for the software to work out what happend to a value if the result of a calculation overflowed the register. With these flag handling issues in mind, a few examples of instruction simulations are shown below. These examples are simplified, and don't calculate 
the half-carry flag.





