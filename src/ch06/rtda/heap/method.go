package heap

import "ch06/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethods := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethods)
		methods[i].copyAttributes(cfMethods)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethods *classfile.MemberInfo) {
	if codeAttr := cfMethods.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}

func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}


func (self *Method) MaxLocals() uint {
	return self.maxLocals
}


func (self *Method) Code() []byte {
	return self.code
}
