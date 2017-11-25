package constants

import "ch06/instructions/base"
import "ch06/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing todo
}
