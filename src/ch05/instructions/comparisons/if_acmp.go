package comparisons

import (
	"ch05/instructions/base"
	"ch05/rtda"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmpPop(frame) {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmpPop(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmpPop(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	val2 := stack.PopRef()
	val1 := stack.PopRef()
	return val1 == val2
}
