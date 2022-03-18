package sl

import (
    "fmt"
    "testing"
)

type User struct {
    Id int64
}

func TestMapper(t *testing.T) {
    var users []User
    for i := 0; i < 5; i++ {
        users = append(users, User{Id: int64(i + 1)})
    }
    var expectIdMap = map[int64]struct{}{
        1: {},
        2: {},
        3: {},
        4: {},
        5: {},
    }
    idList := Mapper(users, func(u User) int64 { return u.Id }).CollectSlice()

    for _, id := range idList {
        if _, ok := expectIdMap[id]; !ok {
            panic(fmt.Sprintf("unexpected %d", id))
        }
    }
}

func TestDistinct(t *testing.T) {
    var users []User
    for i := 0; i < 5; i++ {
        users = append(users, User{Id: int64(i + 1)})
    }
    users = append(users, User{Id: int64(5)})
    idList := Distinct(Mapper(users, func(u User) int64 { return u.Id })).CollectSlice()
    fmt.Println(idList)
}
