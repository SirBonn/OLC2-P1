// registers.go
package arm64

import (
	"fmt"
)

// Register representa un registro ARM64
type Register string

// Registros ARM64 disponibles
const (
	// Registros de propósito general (64-bit)
	X0  Register = "x0"  // Resultado/argumento 0
	X1  Register = "x1"  // Argumento 1
	X2  Register = "x2"  // Argumento 2
	X3  Register = "x3"  // Argumento 3
	X4  Register = "x4"  // Argumento 4
	X5  Register = "x5"  // Argumento 5
	X6  Register = "x6"  // Argumento 6
	X7  Register = "x7"  // Argumento 7
	X8  Register = "x8"  // Indirect result location
	X9  Register = "x9"  // Temporal
	X10 Register = "x10" // Temporal
	X11 Register = "x11" // Temporal
	X12 Register = "x12" // Temporal
	X13 Register = "x13" // Temporal
	X14 Register = "x14" // Temporal
	X15 Register = "x15" // Temporal
	X16 Register = "x16" // IP0 (intra-procedure scratch)
	X17 Register = "x17" // IP1 (intra-procedure scratch)
	X18 Register = "x18" // Platform register
	X19 Register = "x19" // Callee-saved
	X20 Register = "x20" // Callee-saved
	X21 Register = "x21" // Callee-saved
	X22 Register = "x22" // Callee-saved
	X23 Register = "x23" // Callee-saved
	X24 Register = "x24" // Callee-saved
	X25 Register = "x25" // Callee-saved
	X26 Register = "x26" // Callee-saved
	X27 Register = "x27" // Callee-saved
	X28 Register = "x28" // Callee-saved
	X29 Register = "x29" // Frame pointer (FP)
	X30 Register = "x30" // Link register (LR)

	// Registros especiales
	SP  Register = "sp"  // Stack pointer
	XZR Register = "xzr" // Zero register

	// Registros de 32-bit (versiones W de los X)
	W0  Register = "w0"
	W1  Register = "w1"
	W2  Register = "w2"
	W3  Register = "w3"
	W4  Register = "w4"
	W5  Register = "w5"
	W6  Register = "w6"
	W7  Register = "w7"
	W8  Register = "w8"
	W9  Register = "w9"
	W10 Register = "w10"
	W11 Register = "w11"
	W12 Register = "w12"
	W13 Register = "w13"
	W14 Register = "w14"
	W15 Register = "w15"
)

// RegisterAllocator maneja la asignación de registros
type RegisterAllocator struct {
	available   []Register
	used        map[Register]bool
	calleeSaved []Register
	tempRegs    []Register
}

// NewRegisterAllocator crea un nuevo allocator de registros
func NewRegisterAllocator() *RegisterAllocator {
	ra := &RegisterAllocator{
		used: make(map[Register]bool),
		// Registros temporales que podemos usar libremente
		tempRegs: []Register{
			X9, X10, X11, X12, X13, X14, X15,
		},
		// Registros que debemos preservar si los usamos
		calleeSaved: []Register{
			X19, X20, X21, X22, X23, X24, X25, X26, X27, X28,
		},
	}

	// Inicializar registros disponibles con los temporales
	ra.available = make([]Register, len(ra.tempRegs))
	copy(ra.available, ra.tempRegs)

	return ra
}

// Allocate obtiene un registro disponible
func (ra *RegisterAllocator) Allocate() Register {
	if len(ra.available) == 0 {
		// Si no hay registros disponibles, necesitamos hacer spill
		// Por ahora, solo retornamos un error
		panic("out of registers - spilling not yet implemented")
	}

	// Tomar el primer registro disponible
	reg := ra.available[0]
	ra.available = ra.available[1:]
	ra.used[reg] = true

	return reg
}

// AllocateSpecific intenta asignar un registro específico
func (ra *RegisterAllocator) AllocateSpecific(reg Register) error {
	if ra.used[reg] {
		return fmt.Errorf("register %s already in use", reg)
	}

	// Remover de disponibles si está ahí
	for i, r := range ra.available {
		if r == reg {
			ra.available = append(ra.available[:i], ra.available[i+1:]...)
			break
		}
	}

	ra.used[reg] = true
	return nil
}

// Free libera un registro
func (ra *RegisterAllocator) Free(reg Register) {
	if !ra.used[reg] {
		return // Ya está libre
	}

	delete(ra.used, reg)

	// Solo agregar a disponibles si es un registro temporal
	for _, temp := range ra.tempRegs {
		if reg == temp {
			ra.available = append(ra.available, reg)
			break
		}
	}
}

// FreeAll libera todos los registros
func (ra *RegisterAllocator) FreeAll() {
	ra.used = make(map[Register]bool)
	ra.available = make([]Register, len(ra.tempRegs))
	copy(ra.available, ra.tempRegs)
}

// IsCalleeSaved verifica si un registro debe ser preservado
func (ra *RegisterAllocator) IsCalleeSaved(reg Register) bool {
	for _, saved := range ra.calleeSaved {
		if reg == saved {
			return true
		}
	}
	return false
}

// GetArgumentRegister obtiene el registro para el argumento n-ésimo
func GetArgumentRegister(n int) Register {
	if n < 0 || n > 7 {
		panic(fmt.Sprintf("invalid argument number: %d", n))
	}

	argRegs := []Register{X0, X1, X2, X3, X4, X5, X6, X7}
	return argRegs[n]
}

// GetReturnRegister obtiene el registro de retorno
func GetReturnRegister() Register {
	return X0
}

// RegisterInfo contiene información sobre el uso de un registro
type RegisterInfo struct {
	Register    Register
	InUse       bool
	Variable    string // Variable asociada (si aplica)
	IsArgument  bool   // Si contiene un argumento de función
	IsTemporary bool   // Si es un valor temporal
}

// RegisterMap mapea variables a registros
type RegisterMap struct {
	varToReg map[string]Register
	regToVar map[Register]string
}

// NewRegisterMap crea un nuevo mapa de registros
func NewRegisterMap() *RegisterMap {
	return &RegisterMap{
		varToReg: make(map[string]Register),
		regToVar: make(map[Register]string),
	}
}

// Assign asigna un registro a una variable
func (rm *RegisterMap) Assign(variable string, reg Register) {
	// Limpiar asignación previa si existe
	if oldReg, exists := rm.varToReg[variable]; exists {
		delete(rm.regToVar, oldReg)
	}

	// Limpiar variable previa del registro si existe
	if oldVar, exists := rm.regToVar[reg]; exists {
		delete(rm.varToReg, oldVar)
	}

	rm.varToReg[variable] = reg
	rm.regToVar[reg] = variable
}

// GetRegister obtiene el registro asignado a una variable
func (rm *RegisterMap) GetRegister(variable string) (Register, bool) {
	reg, exists := rm.varToReg[variable]
	return reg, exists
}

// GetVariable obtiene la variable asignada a un registro
func (rm *RegisterMap) GetVariable(reg Register) (string, bool) {
	variable, exists := rm.regToVar[reg]
	return variable, exists
}

// Clear limpia todas las asignaciones
func (rm *RegisterMap) Clear() {
	rm.varToReg = make(map[string]Register)
	rm.regToVar = make(map[Register]string)
}
