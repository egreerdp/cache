package main

import (
	"context"
	"fmt"
	"time"

	"github.com/egreerdp/cache"
	"github.com/redis/go-redis/v9"
)

type User struct {
	ID   uint
	Name string
}

func (u *User) Key() string {
	return fmt.Sprintf("%d_%s", u.ID, u.Name)
}

func (u *User) Prefix() string {
	return "user"
}

func main() {
	c, close, err := cache.NewCache(&redis.Options{Addr: "localhost:6379"}, 1*time.Minute, func(ctx context.Context, key string) (*User, error) {
		return &User{
			ID:   2,
			Name: "CallbackUser",
		}, nil
	})
	if err != nil {
		panic(err)
	}

	defer close()

	user := &User{
		ID:   1,
		Name: "User1",
	}

	c.Set(context.TODO(), user)

	u, err := c.Get(context.TODO(), user.Key())
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
}
