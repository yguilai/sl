package sl

type (
    MapperFunc[T, E any] func(t T) (e E)

    Number interface {
        ~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 |
        float32 | float64
    }
)

// Mapper return a slice consisting of the results of applying the given MapperFunc
func Mapper[T, E any](sources []T, fn MapperFunc[T, E]) []E {
    return CollectSlice[E](mapper(sources, fn))
}

func mapper[T, E any](sources []T, fn MapperFunc[T, E]) <-chan E {
    s := make(chan E)
    go func() {
        defer close(s)
        for _, item := range sources {
            s <- fn(item)
        }
    }()
    return s
}

func Distinct[T ~comparable](sources []T) []T {
    s := make(chan T)
    go func() {
        defer close(s)

        m := make(map[T]struct{})
        for _, item := range sources {
            if _, ok := m[item]; !ok {
                m[item] = struct{}{}
                s <- item
            }
        }
    }()
    return CollectSlice[T](s)
}

func DistinctField[T any, E ~comparable](sources []T, fn MapperFunc[T, E]) []E {
    s := make(chan E)
    go func() {
        defer close(s)

        m := make(map[E]struct{})
        for _, item := range sources {
            e := fn(item)
            if _, ok := m[item]; !ok {
                m[item] = struct{}{}
                s <- e
            }
        }
    }()
    return CollectSlice[E](s)
}

func CollectSlice[T any](s <-chan T) []T {
    var slice []T
    for t := range s {
        slice = append(slice, t)
    }
    return slice
}

func ReduceField[T any, E ~Number]() (e E) {
    panic("unimplemented")
    return
}

func Reduce[T ~Number](source []T) (r T) {
    for _, item := range source {
        r += item
    }
    return
}
