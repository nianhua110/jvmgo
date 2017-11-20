package math
type DNEG struct {
  base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame)  {
  stack := frame.OperandStack()
  v1 :=stack PopDouble()
  stack.PushDouble(-v1)
}


type FNEG struct {
  base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtda.Frame)  {
  stack := frame.OperandStack()
  v1 :=stac.PopFloat()
  stack.PushFloat(-v1)
}

type INEG struct {
  base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame)  {
  stack := frame.OperandStack()
  v1 :=stack.PopInt()
  stack.PushInt(-v1)
}


type LNEG struct {
  base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame)  {
  stack := frame.OperandStack()
  v1 :=stack PopLong()
  stack.PushLong(-v1)
}
