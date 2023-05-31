package geecache

import "fmt"

// A ByteView holds an immutable view of bytes.

type ByteView struct {
	b []byte //存储真实的缓存值， byte类型可支持任意数据类型的存储
}

// Len returns the view's length
// 返回其所占的内存大小
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice
// 只读，返回一个拷贝，防止缓存对象被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

// cloneBytes
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func main() {
	fmt.Println("vim-go")
}
