package classfile

import "fmt"

type ClassFile struct {
	//magic unit32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	fmt.Print("Parse 29")
	//fmt.Printf("%v", classData)
	cr := &ClassReader{classData}
	fmt.Print("Parse cr")
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)

	fmt.Print("reader self.constantpool")

	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.interfaces = reader.readUint16s()

	fmt.Print("reader self.constantpool49")

	self.fields = readMembers(reader, self.constantPool)
	fmt.Println("reader self.constantpool fields")
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)

	fmt.Println("reader self.constantpool 55")

}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0XCAFEBABE {
		panic("java.lang.ClassFormatError: magic !")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}

	}
	panic("java.lang.UnsupportedClassVersionError!")
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
