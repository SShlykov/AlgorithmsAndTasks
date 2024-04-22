package cache

type DataAdapter interface {
	Put(key string, value any) error
	Get(key string) (any, error)
}

type InMemoryCacheV2 struct {
	adapter DataAdapter
}
