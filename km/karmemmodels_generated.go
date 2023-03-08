package km

import (
	karmem "karmem.org/golang"
	"unsafe"
)

var _ unsafe.Pointer

var _Null = [37]byte{}
var _NullReader = karmem.NewReader(_Null[:])

type (
	PacketIdentifier uint64
)

const (
	PacketIdentifierKarmemA = 2488188364914332533
)

type KarmemA struct {
	Name     string
	Birthday int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    float64
}

func NewKarmemA() KarmemA {
	return KarmemA{}
}

func (x *KarmemA) PacketIdentifier() PacketIdentifier {
	return PacketIdentifierKarmemA
}

func (x *KarmemA) Reset() {
	x.Read((*KarmemAViewer)(unsafe.Pointer(&_Null[0])), _NullReader)
}

func (x *KarmemA) WriteAsRoot(writer *karmem.Writer) (offset uint, err error) {
	return x.Write(writer, 0)
}

func (x *KarmemA) Write(writer *karmem.Writer, start uint) (offset uint, err error) {
	offset = start
	size := uint(37)
	if offset == 0 {
		offset, err = writer.Alloc(size)
		if err != nil {
			return 0, err
		}
	}
	__NameSize := uint(1 * len(x.Name))
	__NameOffset, err := writer.Alloc(__NameSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+0, uint32(__NameOffset))
	writer.Write4At(offset+0+4, uint32(__NameSize))
	__NameSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Name)), __NameSize, __NameSize}
	writer.WriteAt(__NameOffset, *(*[]byte)(unsafe.Pointer(&__NameSlice)))
	__BirthdayOffset := offset + 8
	writer.Write8At(__BirthdayOffset, *(*uint64)(unsafe.Pointer(&x.Birthday)))
	__PhoneSize := uint(1 * len(x.Phone))
	__PhoneOffset, err := writer.Alloc(__PhoneSize)
	if err != nil {
		return 0, err
	}
	writer.Write4At(offset+16, uint32(__PhoneOffset))
	writer.Write4At(offset+16+4, uint32(__PhoneSize))
	__PhoneSlice := [3]uint{*(*uint)(unsafe.Pointer(&x.Phone)), __PhoneSize, __PhoneSize}
	writer.WriteAt(__PhoneOffset, *(*[]byte)(unsafe.Pointer(&__PhoneSlice)))
	__SiblingsOffset := offset + 24
	writer.Write4At(__SiblingsOffset, *(*uint32)(unsafe.Pointer(&x.Siblings)))
	__SpouseOffset := offset + 28
	writer.Write1At(__SpouseOffset, *(*uint8)(unsafe.Pointer(&x.Spouse)))
	__MoneyOffset := offset + 29
	writer.Write8At(__MoneyOffset, *(*uint64)(unsafe.Pointer(&x.Money)))

	return offset, nil
}

func (x *KarmemA) ReadAsRoot(reader *karmem.Reader) {
	x.Read(NewKarmemAViewer(reader, 0), reader)
}

func (x *KarmemA) Read(viewer *KarmemAViewer, reader *karmem.Reader) {
	__NameString := viewer.Name(reader)
	if x.Name != __NameString {
		__NameStringCopy := make([]byte, len(__NameString))
		copy(__NameStringCopy, __NameString)
		x.Name = *(*string)(unsafe.Pointer(&__NameStringCopy))
	}
	x.Birthday = viewer.Birthday()
	__PhoneString := viewer.Phone(reader)
	if x.Phone != __PhoneString {
		__PhoneStringCopy := make([]byte, len(__PhoneString))
		copy(__PhoneStringCopy, __PhoneString)
		x.Phone = *(*string)(unsafe.Pointer(&__PhoneStringCopy))
	}
	x.Siblings = viewer.Siblings()
	x.Spouse = viewer.Spouse()
	x.Money = viewer.Money()
}

type KarmemAViewer [37]byte

func NewKarmemAViewer(reader *karmem.Reader, offset uint32) (v *KarmemAViewer) {
	if !reader.IsValidOffset(offset, 37) {
		return (*KarmemAViewer)(unsafe.Pointer(&_Null[0]))
	}
	v = (*KarmemAViewer)(unsafe.Add(reader.Pointer, offset))
	return v
}

func (x *KarmemAViewer) size() uint32 {
	return 37
}
func (x *KarmemAViewer) Name(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 0+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *KarmemAViewer) Birthday() (v int64) {
	return *(*int64)(unsafe.Add(unsafe.Pointer(x), 8))
}
func (x *KarmemAViewer) Phone(reader *karmem.Reader) (v string) {
	offset := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 16))
	size := *(*uint32)(unsafe.Add(unsafe.Pointer(x), 16+4))
	if !reader.IsValidOffset(offset, size) {
		return ""
	}
	length := uintptr(size / 1)
	slice := [3]uintptr{
		uintptr(unsafe.Add(reader.Pointer, offset)), length, length,
	}
	return *(*string)(unsafe.Pointer(&slice))
}
func (x *KarmemAViewer) Siblings() (v int32) {
	return *(*int32)(unsafe.Add(unsafe.Pointer(x), 24))
}
func (x *KarmemAViewer) Spouse() (v bool) {
	return *(*bool)(unsafe.Add(unsafe.Pointer(x), 28))
}
func (x *KarmemAViewer) Money() (v float64) {
	return *(*float64)(unsafe.Add(unsafe.Pointer(x), 29))
}
