package cpu
import "fmt"

// Time clock: The Z80 holds two types of clock (m and t)
type clock struct {
	m uint8
	t uint8
}

// Register set
type register struct {
  a, b, c, d, e, h, l, f uint8 // 8-bit register
  pc, sp int
  m, t int                    // Clock for last instr
}

type Z80 struct {
  _c clock
	_r register
}

// Memory interface
type MMU struct {
  rb func(addr uint8) // Read 8-bit byte from a given address
  rw func(addr uint8) // Read 16-bit word from a given address
  //wb func(addr uint8, val uint8) // Write 8-bit byte to a given address
  //ww func(addr uint8, val uint8) // Write 16-bit word to a given address
}

func Init() {
	fmt.Printf("hello")
	_c := clock{0, 0}
	_Z80 := Z80{clock{0, 0}, register{0,0,0,0,0,0,0,0, 0,0, 0,0}}
	fmt.Println(_c)
  fmt.Println(_Z80._c)
  fmt.Println(_Z80._r)
	_Z80.ADD()
}

// Add E to A, leaving results in A (ADD A, E)
func (Z80 *Z80) ADD() {
  Z80._r.a += Z80._r.e                    // Perform addition
  Z80._r.f = 0                             // Clear flags
	// Check for zero
  if((Z80._r.a & 255)==0){
	  Z80._r.f = 0x80
	}
	if(Z80._r.a > 255){
		Z80._r.f = 0x10     // Check for carry
	}
  var s = fmt.Sprintf("%x", Z80._r.f)
  fmt.Println(s)
	Z80._r.a &= 255                      // Mask to 8-bits
  Z80._r.m = 1
  Z80._r.t = 4               // 1 M-time taken
}

// Compare B to A, setting flags (CP A, B)
func (Z80 *Z80) CP() {
  var i = Z80._r.a                     // Temp copy of A
  i -= Z80._r.b                        // Subtract B
  Z80._r.f = 0x40                     // Set subtraction flag
  if((i & 255)==0){
		Z80._r.f = 0x80       // Check for zero
	}
	if(i < 0){
		Z80._r.f = 0x10             // Check for underflow
	}
	Z80._r.m = 1
	Z80._r.t = 4           // 1 M-time taken
}

// No-operation (NOP)
func (Z80 *Z80) NOP() {
  Z80._r.m = 1
  Z80._r.t = 4                // 1 M-time taken
}

// Push registers B and C to the stack (PUSH BC)
/*
func (Z80 *Z80) PUSHBC() {
	Z80._r.sp--;                               // Drop through the stack
	MMU.wb(Z80._r.sp, Z80._r.b);               // Write B
	Z80._r.sp--;                               // Drop through the stack
	MMU.wb(Z80._r.sp, Z80._r.c);               // Write C
	Z80._r.m = 3; Z80._r.t = 12;               // 3 M-times taken
}
*/


