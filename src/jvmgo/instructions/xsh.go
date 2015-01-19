package instructions

import "jvmgo/rtda"

// Shift left int
type ishl struct {NoOperandsInstruction}
func (self *ishl) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x1f
    result := v1 << s
    stack.PushInt(result)
}

// Arithmetic shift right int
type ishr struct {NoOperandsInstruction}
func (self *ishr) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x1f
    result := v1 >> s
    stack.PushInt(result)
}

// Logical shift right int
type iushr struct {NoOperandsInstruction}
func (self *iushr) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x1f
    result := int32(uint32(v1) >> s)
    stack.PushInt(result)
}


// Shift left long
type lshl struct {NoOperandsInstruction}
func (self *lshl) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x3f
    result := v1 << s
    stack.PushLong(result)
}

// Arithmetic shift right long
type lshr struct {NoOperandsInstruction}
func (self *lshr) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x3f
    result := v1 >> s
    stack.PushLong(result)
}

// Logical shift right long
type lushr struct {NoOperandsInstruction}
func (self *lushr) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x3f
    result := int64(uint64(v1) >> s)
    stack.PushLong(result)
}
