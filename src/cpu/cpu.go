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


func Init() {
	fmt.Printf("hello")
	_c := clock{0, 0}
	_Z80 := Z80{clock{0, 0}, register{0,0,0,0,0,0,0,0, 0,0, 0,0}}
  fmt.Println(_c)
  fmt.Println(_Z80._c)
  fmt.Println(_Z80._r)
}

// Add E to A, leaving results in A (ADD A, E)
func ADD(_Z80 *Z80) {
  _Z80._r.a += _Z80._r.e                    // Perform addition
  _Z80._r.f = 0                             // Clear flags
  //if(!(_Z80._r.a & 255)) _Z80._r.f |= 0x80  // Check for zero
  //if(_Z80._r.a > 255) _Z80._r.f |= 0x10     // Check for carry
  //_Z80._r.a &= 255                          // Mask to 8-bits
  _Z80._r.m = 1
  _Z80._r.t = 4               // 1 M-time taken
}

// Compare B to A, setting flags (CP A, B)
func CP(_Z80 *Z80) {
  var i = _Z80._r.a                     // Temp copy of A
  i -= _Z80._r.b                        // Subtract B
  //_Z80._r.f |= 0x40                     // Set subtraction flag
  //if(!(i & 255)) _Z80._r.f |= 0x80       // Check for zero
  //if(i < 0) Z80._r.f |= 0x10             // Check for underflow
  _Z80._r.m = 1
	_Z80._r.t = 4           // 1 M-time taken
}

// No-operation (NOP)
func NOP(_Z80 *Z80) {
  _Z80._r.m = 1
  _Z80._r.t = 4                // 1 M-time taken
}

