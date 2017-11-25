package classfile

type ExceptionAttribute struct {
	exceptionIndextable []uint16
}

func (self *ExceptionAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndextable = reader.readUint16s()
}
func (self *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndextable
}
