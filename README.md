# go-gameboy

## The model
The traditional model of a computer is a processing unit, which gets told what to do by a program of 
instructions; the program might be accessed with its own special memory, or it might be sitting in the same
area as normal memory, depending on the computer. Each instruction takes a short amount of time to run.
and they're all run one by one. From the CPU's perspective, a loop starts up as soon as the computer is
turned on, to fetch an instruction from memory, work out what it says, and execute it.

In order to keep track of where the CPU is within the program, a number is held by the CPU called the Program
Counter(PC). After an instructin is fetched from memory, the PC is advanced by however many bytes make up the 
instruction.

### usage
```
$ go run xxx.go
$ go build xxx.go
```



