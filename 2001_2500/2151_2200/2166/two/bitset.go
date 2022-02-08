package two

type Bitset struct {
	data    []byte
	revert  []byte
	oneNums int
	size    int
}

func Constructor(size int) Bitset {
	bs := Bitset{
		data:   make([]byte, size),
		revert: make([]byte, size),
		size:   size,
	}
	for idx := 0; idx < bs.size; idx++ {
		bs.data[idx] = 0 + '0'
		bs.revert[idx] = 1 + '0'
	}
	return bs
}

func (bs *Bitset) Fix(idx int) {
	if bs.data[idx]-'0' == 0 {
		bs.oneNums++
		bs.data[idx] = 1 + '0'
		bs.revert[idx] = 0 + '0'
	}
}

func (bs *Bitset) Unfix(idx int) {
	if bs.data[idx]-'0' == 1 {
		bs.oneNums--
		bs.data[idx] = 0 + '0'
		bs.revert[idx] = 1 + '0'
	}
}

func (bs *Bitset) Flip() {
	bs.data, bs.revert = bs.revert, bs.data
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
	return string(bs.data)
}
