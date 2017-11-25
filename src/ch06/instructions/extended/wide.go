package extended

import (
	"ch06/instructions/base"
	"ch06/instructions/loads"
	"ch06/instructions/math"
	"ch06/instructions/stores"
	"ch06/rtda"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36:
		inst := &strores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37:
		inst := &strores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38:
		inst := &strores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39:
		inst := &strores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3A:
		inst := &strores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0xa9:
		panic("Unsupported code: 0xa9")
	}
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
