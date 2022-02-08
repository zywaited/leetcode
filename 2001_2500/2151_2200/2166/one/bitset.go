package one

type Bitset struct {
	revert  byte
	fixes   []byte
	oneNums int
	size    int
}

func Constructor(size int) Bitset {
	return Bitset{
		fixes: make([]byte, size),
		size:  size,
	}
}

func (bs *Bitset) Fix(idx int) {
	if bs.fixes[idx]^bs.revert == 0 {
		bs.fixes[idx] ^= 1
		bs.oneNums++
	}
}

func (bs *Bitset) Unfix(idx int) {
	if bs.fixes[idx]^bs.revert == 1 {
		bs.fixes[idx] ^= 1
		bs.oneNums--
	}
}

func (bs *Bitset) Flip() {
	bs.revert ^= 1
	bs.oneNums = bs.size - bs.oneNums
}

func (bs *Bitset) All() bool {
	return bs.Count() == bs.size
}

func (bs *Bitset) One() bool {
	return bs.Count() > 0
}

func (bs *Bitset) Count() int {
	return bs.oneNums
}

func (bs *Bitset) ToString() string {
	vs := make([]byte, 0, bs.size)
	for idx := 0; idx < bs.size; idx++ {
		vs = append(vs, (bs.fixes[idx]^bs.revert)+'0')
	}
	return string(vs)
}
