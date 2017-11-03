package rtda
type Frame struct {
  lower *Frame
  localVars LocalVars
  operandStack *OperandStack
}

func newFrame(maxLocals, maxStack uint) *Frame  {
  return &Frame{
    localVars: new LocalVars(maxLocals),
    operandStack: newOperandStack(maxStack),
  }
}
