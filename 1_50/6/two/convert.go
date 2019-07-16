package two

import "bytes"

func Convert(s string, numRows int) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	if l == 1 || numRows == 1 {
		return s
	}

	// 判断有多少行
	row := numRows
	if l < numRows {
		row = l
	}
	down := false
	zs := make([]bytes.Buffer, row)
	cr := 0
	for i := 0; i < l; i++ {
		zs[cr].WriteByte(s[i])
		// 判断行数是否发生变化
		if cr == 0 || cr == numRows-1 {
			down = !down
		}

		if down {
			cr++
		} else {
			cr--
		}
	}
	buffer := zs[0]
	for i := 1; i < row; i++ {
		buffer.Write(zs[i].Bytes())
	}
	return buffer.String()
}
