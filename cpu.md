# CPU

## The model
The traditional model of a computer is a processing unit, which gets told what to do by a program of 
instructions; the program might be accessed with its own special memory, or it might be sitting in the same
area as normal memory, depending on the computer. Each instruction takes a short amount of time to run.
and they're all run one by one. From the CPU's perspective, a loop starts up as soon as the computer is
turned on, to fetch an instruction from memory, work out what it says, and execute it.

�]���̃R���s���[�^���f���́A�v���Z�b�V���O���j�b�g�ł���B����́A���߃R�[�h�ɂ���ĉ������ׂ����w���𓾂�B
���̃v���O�����́A�Ǝ��̃������ɂ���ăA�N�Z�X����邩�A�ʏ�̃������̈�ɏ�����Ă��邾�낤�B
�e���߂́A�Z���ԂŁA�����s�����BCPU�̊ϓ_�ł́A�R���s���[�^���N������Ƃ����Ƀ��[�v���n�܂�A
���������疽�߂��t�F�b�`���A������Ɏ��s����B

In order to keep track of where the CPU is within the program, a number is held by the CPU called the Program
Counter(PC). After an instructin is fetched from memory, the PC is advanced by however many bytes make up the 
instruction.

CPU���v���O�����̂ǂ��ɂ���̂����g���b�N���邽�߂ɁAProgram Counter (PC) �Ƃ������l�� CPU �ɂ���ĕێ������B���߂�����������t�F�b�`���ꂽ��APC �͖��߂��\�����鑽���̃o�C�g��ɂ���ăC���N�������g�����B

<img src="https://github.com/a2-ito/gist_images/blob/master/02_go-gameboy/01_the_fetch-decode-execute_loop.png" width="320">

The CPU in the original GameBoy is a modified Zilog Z80m, so the following things are pertinent:

- The Z80 is an 8-bit chip, so all the internal workings operate on one byte at a time;
- The memory interface can address up to 65,5536 bytes (a 16-bit address bus);
- Programs are accessed through the same address bus as normal memory;
- An instruction can be anywhere between one and three bytes.

In addition to the PC, other numbers are held inside the CPU that can be used for calculation, and they're
referred to as registers: A, B, C, D, E, H, and L. Each of them is one byte, so each one can hold a value from 0
to 255. Most of the instructions in the Z80 are used to handle values in these registers: loading a value from
memory into a register, adding or subtracting values, and so forth.

PC �ɉ����āA���̐��l���v�Z���p�̂��߂�CPU���ɕێ������BA, B, C, D, E, H, L �Ƃ������W�X�^�Ƃ��ĎQ�Ƃ����B���ꂼ��1�o�C�g�ł���A0����255�܂ł̐��l��ێ����邱�Ƃ��ł���BZ80���̂قƂ�ǂ̖��߂͂����̃��W�X�^���������߂ɗ��p�����B���������烌�W�X�^�Ƀ��[�h����A���Z����A�l�𒊏o����A�ȂǂȂǁB

If there are 256 possible values in the first byte of an instruction, that makes for 256 possible instructions
in the basic table. That table is detailed in the Gameboy Z80 opcode map released on this site. 
Each of these can be simulated by a JavaScript function, that operates on an internal model of the registers,
and produces effects on an internal model of the memory interface.

256 �̉\�Ȓl

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

Since the basic calculation registers are 8-bits, the carry flag allows for the software to work out what happend to a value if the result of a calculation overflowed the register. With these flag handling issues in mind, a few examples of instruction simulations are shown below. These examples are simplified, and don't calculate the half-carry flag.

## Memory interfacing
A processor that can manipulate registers within itself is all well and good, but it must be able to put results into memory to be useful. In the same way, the above CPU emulation requires an interface to emulated memory; this can be provided by a memory management unit (MMU). Since the Gameboy itself doesn't contain a complicated MMU, the emulated unit can be quite simple.

���W�X�^�𑀍�ł���v���Z�b�T�͊�{�ǂ����A�������Ɍ��ʂ��������߂Ȃ���΂Ȃ�Ȃ��B���l�ɁA��q��CPU�G�~�����[�V�����́A�G�~�����[�g���ꂽ�������ւ̃C���^�[�t�F�[�X���K�v�ł���B����́AMemory Management Unit(MMU) �ɂ���Ē񋟂����B�Q�[���{�[�C���g�͕��G�� MMU �͊܂�ł��Ȃ����A���̃G�~�����[�g���ꂽ���j�b�g�͔��ɃV���v���ɂȂ蓾��B

At this point, the CPU only needs to know that an interfaces is present; the details of how the Gameboy maps banks of memory and hardware onto the address bus are inconsequential to the processor's operation. Four operatoins are required by the CPU:

���̓_�ɂ����āACPU �̓C���^�[�t�F�[�X�����݂��Ă��邱�Ƃ�m�邾���ł悢�B�Q�[���{�[�C���������o���N�ɂǂ��}�b�v���Ă��邩�̏ڍׂ�n�[�h�E�F�A�A�h���X�o�X�́A�v���Z�b�T�̃I�y���[�V�����ɂ����ďd�v�ł͂Ȃ��BCPU�ɂ́A4�̃I�y���[�V�������K�v�Ƃ����B

# Dispatch and reset
With the instructions in place, the remaining pieces of the puzzle for the CPU are to reset the CPU when it starts up, and to feed instructions to the emulations routines. Having a reset routine allows for the CPU to be stopped and "rewound" to the start of execution.

���̖��߂��z�u����Ă���ƁACPU�͋N������CPU�����Z�b�g���A���[�e�B�����G�~�����[�g����B���Z�b�g���[�e�B���������Ƃ́ACPU����~���ăX�^�[�g���s��"�����߂�"���Ƃ����e����B

In order for the emulation to run, it has to emulate the fetch-decode-execute sequence detailed earlier. "Execute" is taken care of by the instruction emulation fuctions, but fetch and decode require a specialist piece of code, known as a "dispatch loop". This loop takes each instruction, decodes where it must be sent for execution, and dispatches it to the function in question.

���s�̃G�~�����[�V�����̂��߁Afetch-decode-execute �̃V�[�P���X���G�~�����[�g���Ȃ���΂Ȃ�Ȃ��B"Execute"�͖��߃G�~���[���[�V�����֐��ɂ���ăP�A����邪�AFetch �� Decode �͓��ʂȃR�[�h���K�v�ƂȂ�B"dispatch loop" �Œm����B���̃��[�v�͊e���߂��Ăяo���A���s�̂ǂ��ɑ����邩���f�R�[�h���A������f�B�X�p�b�`����B

## Usage in a system emulation 
Implementing a Z80 emulation is useless without an emulation to run it. In the next part of this series, the work of emulating the Gameboy begins: I'll be looking at the Gameboy's memory map, and how a game image can be loaded into the emulator over the Web.

Z80�G�~�����[�^�̎����́A��������s���邽�߂̃G�~�����[�V�����ł���B���̃p�[�g�ł́A�Q�[���{�[�C���G�~�����[�V�������郏�[�N���n�߂�B�Q�[���{�[�C�̃������}�b�v�y�сA�Q�[���C���[�W���ǂ̂悤�Ƀ��[�h����邩�𒭂߂�B

The complete Z80 core is availavle: http://imrannazar.com/content/files/jsgb.z80.js; please feel free to let me know if you encounter any bugs in the implementation.

