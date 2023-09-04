package slavea

// BinLogFileNamePos结构体
type BinLogFileNamePos struct {
	Name string
	Pos  int64
}

// BinLogFileNamePos结构体构造函数
func NewBinLogFileNamePos(name string, pos int64) *BinLogFileNamePos {
	return &BinLogFileNamePos{
		Name: name,
		Pos:  pos,
	}
}
