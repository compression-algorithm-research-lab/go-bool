package bool_array

import (
	"fmt"
	if_expression "github.com/golang-infrastructure/go-if-expression"
)

type BoolArray struct {
	values   []byte
	capacity int
}

// New 创建一个布尔数组
func New(capacity int) *BoolArray {

	if capacity < 0 {
		panic(fmt.Errorf("capacity %d is invalid", capacity))
	}

	return &BoolArray{
		// 预分配容量
		values:   make([]byte, (capacity+7)/8),
		capacity: capacity,
	}
}

// 检查下标是否合法，比如越界检查之类的
func (x *BoolArray) checkIndex(indexes ...int) {
	for _, index := range indexes {
		if index < 0 || index >= x.capacity {
			panic(fmt.Errorf("index %d out of bound", index))
		}
	}
}

// FillAll 把布尔数组所有的值都填充为给定的值
func (x *BoolArray) FillAll(value bool) *BoolArray {
	return x.Fill(0, x.capacity, value)
}

// Fill 往给定的左开右闭区间填充给定的值
func (x *BoolArray) Fill(fromIndex, toIndex int, value bool) *BoolArray {

	x.checkIndex(fromIndex, toIndex)

	for fromIndex < toIndex {
		x.Set(fromIndex, value)
	}
	return x
}

// Set 设置给定位置的值
func (x *BoolArray) Set(index int, v bool) *BoolArray {

	x.checkIndex(index)

	targetIndex := index / 8
	offset := index % 8
	byteValue := byte(if_expression.Return(v, 1, 0))
	// 除了offset其它位置都原样拷贝
	oldByte := x.values[targetIndex] ^ (1 << offset)
	// 然后再把offset位置的新的值设置上
	newByte := oldByte | (byteValue << offset)
	x.values[targetIndex] = newByte
	return x
}

// Get 获取给定下标位置当前存储的布尔值
func (x *BoolArray) Get(index int) bool {

	x.checkIndex(index)

	targetIndex := index / 8
	offset := index % 8
	intValue := (int(x.values[targetIndex]) >> offset) & 0x01
	return if_expression.Return(intValue == 1, true, false)
}

// ToBinaryString 把数组转为二进制字符串形式
func (x *BoolArray) ToBinaryString() string {
	runes := make([]rune, x.capacity)
	// TODO 效率更高的遍历方式
	for index := 0; index < x.capacity; index++ {
		runes[index] = if_expression.Return(x.Get(index), '1', '0')
	}
	return string(runes)
}
