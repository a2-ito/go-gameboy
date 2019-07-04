package memory
//import "fmt"

// Memory interface
type MMU struct {
	// Flag indicating BIOS is mapped in 
	// BIOS is unmapped with the first instruction above 0x00FF
	_inbios uint8

	// Memory regions (initialised at reset time)
	_bios []uint8
	//_rom uint8[]
	//_wram uint8[]
	//_eram uint8[]
	//_zram uint8[]

  // Read 8-bit byte from a given address
	//
	rb func(addr uint16)
//		switch(addr & 0xF000){
//			// BIOS
//		}
	//rw func(addr uint8) // Read 16-bit word from a given address
  //wb func(addr uint8, val uint8) // Write 8-bit byte to a given address
  //ww func(addr uint8, val uint8) // Write 16-bit word to a given address
}

func NewMMU() {
	_MMU := new(MMU)
	_MMU._inbios = 0
}

func Rb(addr uint16) {
		// BIOS (256b)/ROM0
		switch(addr & 0xF000){
			case 0x0000:
			//if(MMU._inbios){
			//}
				//return MMU._rom[addr]
			case 0x1000:
			case 0x2000:
			case 0x3000:
			case 0x4000:
			case 0x5000:
			case 0x6000:
			case 0x7000:
			case 0x8000:
			case 0x9000:
			case 0xA000:
			case 0xB000:
			case 0xC000:
			case 0xD000:
			case 0xE000:
			case 0xF000:
			default: 
		}
}

