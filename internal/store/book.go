package store

var (
	ListScheduleBook map[string][]string
)

type BookStore interface {
	CheckManyUserAtTimeRange(key string) int
}

type bookStoreImpl struct {
}

func NewBookStore() BookStore {
	return &bookStoreImpl{}
}

func (b *bookStoreImpl) CheckManyUserAtTimeRange(key string) int {
	return len(ListScheduleBook[key])
}
