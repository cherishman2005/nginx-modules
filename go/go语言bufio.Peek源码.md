# go语言bufio.Peek源码

bufio.Peek(n)返回前N个未读字节 ，不会更改已读计数的值。在方法调用后，要查看返回参数error是否为nil，以确保操作满足要求。

```golang
func (b *Reader) Peek(n int) ([]byte, error) {
	if n < 0 {
		return nil, ErrNegativeCount
	}

	b.lastByte = -1
	b.lastRuneSize = -1

	// 当未读字节数小于n，且缓冲区不满(b.w-b.r < len(b.buf))，
	// 即缓冲区中从头到尾必须都是未读字节才算是缓冲区已经满
	// 且 b.err 为nil，这三者都满足时，开始调用b.fill()填充缓冲区
	// fill()方法会把未读数据移动到缓冲区头部，并把后面空出来的部分写满
	// for 保证了至少可以把缓冲区填满
	for b.w-b.r < n && b.w-b.r < len(b.buf) && b.err == nil {
		b.fill() // b.w-b.r < len(b.buf) => buffer is not full
	}

	// 当要读取的字节数大于缓冲区长度时，返回所有未读节字，并附带错误信息
	if n > len(b.buf) {
		return b.buf[b.r:b.w], ErrBufferFull
	}

	// 0 <= n <= len(b.buf)
	var err error
	// 当要读取的字节数大于所有未读节字数时，返回所有未读节字，并附带错误信息
	if avail := b.w - b.r; avail < n {
		// not enough data in buffer
		n = avail
		err = b.readErr()
		if err == nil {
			err = ErrBufferFull
		}
	}
	// 当 n 小于 未读字节数时，程序直接跳到这里
	return b.buf[b.r : b.r+n], err
}
```
