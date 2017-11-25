package conversions

import (
	"ch06/instructions/base"
	"ch06/rtda"
)

type L2F struct {
	base.NoOperandsInstruction
}

type L2I struct {
	base.NoOperandsInstruction
}
type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := int32(d)
	stack.PushInt(i)
}

func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := float32(d)
	stack.PushFloat(i)
}

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := float64(d)
	stack.PushDouble(i)
}
