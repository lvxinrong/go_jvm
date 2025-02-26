package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

/*
*
读取byte字节数组里面的一个8字节元素，也就是data数组的第0个元素，
读取之后将data数组删除第0个元素
*/
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[:1]
	return val
}

func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

// 读取uint16表，表的开头由开头的uint16数据指出(也就是data[0]位置记录的是后面uint16数组的大小)
func (cr *ClassReader) readUint16S() []uint16 {
	n := cr.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.readUint16()
	}
	return s
}

func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}
