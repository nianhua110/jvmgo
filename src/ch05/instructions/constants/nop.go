package constants

import "ch05/instructions/base"
import "ch05/rtda"

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing todo
}
