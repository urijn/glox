package stack

import (
	"bytes"
	"github.com/urijn/glox/backend/vm"
	"github.com/urijn/glox/opcode"
	"github.com/urijn/glox/value"
)

func (v *VM) incrementIP() {
	v.ip += 1
}

func (v *VM) readByte() byte {
	defer v.incrementIP()
	return v.chunk.Code[v.ip]
}

func (v *VM) readConstant() *value.Value {
	return v.chunk.Constants.Values[v.readByte()]
}

func (v *VM) binaryOperation(valType value.ValueType, op byte) vm.InterpretResult {
	b := v.Peek(0)
	a := v.Peek(1)

	if !b.Is(value.ValNumber) || !a.Is(value.ValNumber) {
		v.runtimeError("Operands must be numbers.")
		return vm.InterpretRuntimeError
	}

	bVal := v.Pop().Val.GetNumber()
	aVal := v.Pop().Val.GetNumber()

	var result interface{}
	switch op {
	case opcode.OpAdd:
		result = aVal + bVal
	case opcode.OpDivide:
		result = aVal / bVal
	case opcode.OpMultiply:
		result = aVal * bVal
	case opcode.OpSubtract:
		result = aVal - bVal
	case opcode.OpGreater:
		result = aVal > bVal
	case opcode.OpLess:
		result = aVal < bVal
	}

	v.Push(value.NewValue(valType, result))

	return vm.InterpretOk
}

func (v *VM) isFalsey(val *value.Value) bool {
	return val.Is(value.ValNil) ||
		(val.Is(value.ValBool) && !val.Val.GetBool())
}

func (v *VM) valuesEqual(a, b *value.Value) bool {
	if a.ValType != b.ValType {
		return false
	}

	switch a.ValType {
	case value.ValBool:
		return a.Val.GetBool() == b.Val.GetBool()
	case value.ValNil:
		return true
	case value.ValNumber:
		return a.Val.GetNumber() == a.Val.GetNumber()
	case value.ValObj:
		return a.Val.Get() == b.Val.Get()
	default:
		return false
	}
}

func (v *VM) concatenate()  {
	bStr := (v.Pop().Val.GetObject()).(*value.ObjectString)
	aStr := v.Pop().Val.GetObject().(*value.ObjectString)

	var buffer bytes.Buffer
	buffer.WriteString(aStr.String())
	buffer.WriteString(bStr.String())
	v.Push(value.NewObjectValueString(buffer.String()))
}