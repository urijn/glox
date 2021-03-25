package opcode

const (
	// Load the constant for use
	OpReturn byte = iota + 1
	OpConstant
	OpNegate
	OpAdd
	OpSubtract
	OpMultiply
	OpDivide
	OpNil
	OpTrue
	OpFalse
	OpNot
	OpEqual
	OpGreater
	OpLess
	OpPop
	OpPrint
	OpDefineGlobal
	OpGetGlobal
	OpSetGlobal
)