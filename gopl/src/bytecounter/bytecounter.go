package bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // 转换int为ByteCounter类型
	return len(p), nil
}
