package regnum

import (
	"fmt"
)

// The mapping between hardware registers and DWARF registers, See
// https://loongson.github.io/LoongArch-Documentation/LoongArch-Vol1-EN.html
// https://loongson.github.io/LoongArch-Documentation/LoongArch-ELF-ABI-EN.html

const (
	// General-purpose Register
	LOONG64_R0  = 0
	LOONG64_R1  = 1
	LOONG64_R2  = 2
	LOONG64_R3  = 3
	LOONG64_R4  = 4
	LOONG64_R5  = 5
	LOONG64_R6  = 6
	LOONG64_R7  = 7
	LOONG64_R8  = 8
	LOONG64_R9  = 9
	LOONG64_R10 = 10
	LOONG64_R11 = 11
	LOONG64_R12 = 12
	LOONG64_R13 = 13
	LOONG64_R14 = 14
	LOONG64_R15 = 15
	LOONG64_R16 = 16
	LOONG64_R17 = 17
	LOONG64_R18 = 18
	LOONG64_R19 = 19
	LOONG64_R20 = 20
	LOONG64_R21 = 21
	LOONG64_R22 = 22
	LOONG64_R23 = 23
	LOONG64_R24 = 24
	LOONG64_R25 = 25
	LOONG64_R26 = 26
	LOONG64_R27 = 27
	LOONG64_R28 = 28
	LOONG64_R29 = 29
	LOONG64_R30 = 30
	LOONG64_R31 = 31

	// Floating-point Register
	LOONG64_F0  = 32
	LOONG64_F1  = 33
	LOONG64_F2  = 34
	LOONG64_F3  = 35
	LOONG64_F4  = 36
	LOONG64_F5  = 37
	LOONG64_F6  = 38
	LOONG64_F7  = 39
	LOONG64_F8  = 40
	LOONG64_F9  = 41
	LOONG64_F10 = 42
	LOONG64_F11 = 43
	LOONG64_F12 = 44
	LOONG64_F13 = 45
	LOONG64_F14 = 46
	LOONG64_F15 = 47
	LOONG64_F16 = 48
	LOONG64_F17 = 49
	LOONG64_F18 = 50
	LOONG64_F19 = 51
	LOONG64_F20 = 52
	LOONG64_F21 = 53
	LOONG64_F22 = 54
	LOONG64_F23 = 55
	LOONG64_F24 = 56
	LOONG64_F25 = 57
	LOONG64_F26 = 58
	LOONG64_F27 = 59
	LOONG64_F28 = 60
	LOONG64_F29 = 61
	LOONG64_F30 = 62
	LOONG64_F31 = 63

	// Floating condition flag register
	LOONG64_FCC0 = 64
	LOONG64_FCC1 = 65
	LOONG64_FCC2 = 66
	LOONG64_FCC3 = 67
	LOONG64_FCC4 = 68
	LOONG64_FCC5 = 69
	LOONG64_FCC6 = 70
	LOONG64_FCC7 = 71

	// Extra, not defined in ELF-ABI specification
	LOONG64_ERA  = 72
	LOONG64_BADV = 73

	_LOONG64_MaxRegNum = LOONG64_BADV

	// See golang src/cmd/link/internal/loong64/l.go
	LOONG64_LR = LOONG64_R1  // ra: address fro subroutine
	LOONG64_SP = LOONG64_R3  // sp: stack pointer
	LOONG64_FP = LOONG64_R22 // fp: frame pointer
	LOONG64_PC = LOONG64_ERA // rea : exception program counter
)

func LOONG64ToName(num uint64) string {
	switch {
	case num <= LOONG64_R31:
		return fmt.Sprintf("R%d", num)

	case num >= LOONG64_F0 && num <= LOONG64_F31:
		return fmt.Sprintf("F%d", num-32)

	case num >= LOONG64_FCC0 && num <= LOONG64_FCC7:
		return fmt.Sprintf("FCC%d", num-64)

	case num == LOONG64_ERA:
		return fmt.Sprintf("ERA")

	case num == LOONG64_BADV:
		return fmt.Sprintf("BADV")

	default:
		return fmt.Sprintf("Unknown%d", num)
	}
}

func LOONG64MaxRegNum() uint64 {
	return _LOONG64_MaxRegNum
}

var LOONG64NameToDwarf = func() map[string]int {
	r := make(map[string]int)
	for i := 0; i <= 31; i++ {
		r[fmt.Sprintf("R%d", i)] = LOONG64_R0 + i
	}

	for i := 0; i <= 31; i++ {
		r[fmt.Sprintf("F%d", i)] = LOONG64_F0 + i
	}

	for i := 0; i <= 7; i++ {
		r[fmt.Sprintf("FCC%d", i)] = LOONG64_FCC0 + i
	}

	r[fmt.Sprintf("ERA")] = LOONG64_ERA
	r[fmt.Sprintf("BADV")] = LOONG64_BADV

	return r
}()
