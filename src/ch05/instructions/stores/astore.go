package strores

import "ch05/rtda"
import "ch05/instructions/base"

type ASTORE struct {
	base.Index8Instruction
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func _atore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_atore(frame, uint(self.index))
}

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_atore(frame, 0)
}

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_atore(frame, 1)
}
func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_atore(frame, 2)
}
func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_atore(frame, 3)
}
