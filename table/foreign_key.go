package table

// 外键
type ForeignKey struct {
	Name      string
	RefSchema string
	RefTable  string
	RefColumn string
}
