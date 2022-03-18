package sl

type (
    MapperFunc[T, E any] func(t T) (e E)

    Number interface {
        ~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 |
        float32 | float64
    }

    PseudoStream[T any] <-chan T
)

// CollectSlice collect elements from the channel
func (s PseudoStream[T]) CollectSlice() (slice []T) {
    for item := range s {
        slice = append(slice, item)
    }
    return
}

// Stream generate a pseudo-stream from slice
func Stream[T any](slice []T) PseudoStream[T] {
    s := make(chan T)
    go func() {
        defer close(s)
        for _, t := range slice {
            s <- t
        }
    }()
    return s
}

// Mapper return a PseudoStream of the results of applying the given MapperFunc
func Mapper[T, E any](sources []T, fn MapperFunc[T, E]) PseudoStream[E] {
    s := make(chan E, 1)
    go func() {
        defer close(s)
        for _, item := range sources {
            s <- fn(item)
        }
    }()
    return s
}

// MapperStream return a PseudoStream of the results of applying the given MapperFunc
func MapperStream[T, E any](ps PseudoStream[T], fn MapperFunc[T, E]) PseudoStream[E] {
    s := make(chan E, 1)
    go func() {
        defer close(s)
        for item := range ps {
            s <- fn(item)
        }
    }()
    return s
}

func Distinct[T comparable](ps PseudoStream[T]) PseudoStream[T] {
    s := make(chan T, 1)
    go func() {
        defer close(s)

        m := make(map[T]struct{})
        for item := range ps {
            if _, ok := m[item]; !ok {
                m[item] = struct{}{}
                s <- item
            }
        }
    }()
    return s
}

func DistinctField[T any, E comparable](sources []T, fn MapperFunc[T, E]) PseudoStream[E] {
    s := make(chan E, 1)
    go func() {
        defer close(s)

        m := make(map[E]struct{})
        for _, item := range sources {
            e := fn(item)
            if _, ok := m[e]; !ok {
                m[e] = struct{}{}
                s <- e
            }
        }
    }()
    return s
}

func ReduceField[T any, E Number]() (e E) {
    panic("unimplemented")
    return
}

func Reduce[T Number](source []T) (r T) {
    for _, item := range source {
        r += item
    }
    return
}
