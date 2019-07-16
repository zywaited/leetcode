package one

func ReverseBits(num uint32) uint32 {
	r := uint32(0)
	i := uint8(32)
	for i > 0 {
		r <<= 1
		r += num & 1
		num >>= 1
		i--
	}
	return r
}
