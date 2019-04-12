package entities

// BookStatus 图片状态
type BookStatus int

const (
	// NotLend 未借出
	NotLend BookStatus = iota
	// Lended 已借出
	Lended
)

// RecordType 记录类型
type RecordType int

const (
	// Borrow 借
	Borrow RecordType = iota
	// Return 还
	Return
)
