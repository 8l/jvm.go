package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations 
type invokespecial struct {Index16Instruction}
func (self *invokespecial) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*class.ConstantMethodref)
    method := cMethodRef.SpecialMethod()

    if method.IsNative() {
        // exec native method
        nativeMethod := method.NativeMethod().(func(*rtda.Frame))
        nativeMethod(frame)
        return
    }

    thread.InvokeMethod(method)
}