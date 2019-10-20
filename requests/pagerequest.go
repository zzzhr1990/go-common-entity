package requests

// OrderByRequest ORDER
type OrderByRequest interface {
	GetName() string
	GetAsc() bool
}

// PageInfo info of pge
type PageInfo interface {
	GetPage() int64
	GetPageSize() int64
	GetTotalPage() int64
	GetTotalRecord() int64
}

// ListInfo info of list
type ListInfo interface {
	GetStart() int64
	GetLimit() int64
}
