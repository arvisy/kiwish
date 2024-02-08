package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-user/model"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetUserCache(c *redis.Client, User *model.User, ctx context.Context) error {
	p, err := json.Marshal(User)
	if err != nil {
		return err
	}

	key := "User" + strconv.Itoa(User.Id)

	err = c.Set(ctx, key, p, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUser(c *redis.Client, UserID int, ctx context.Context) (model.User, error) {
	var User model.User
	key := "User" + strconv.Itoa(UserID)

	str, err := c.Get(ctx, key).Result()
	if err == redis.Nil {
		return model.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return model.User{}, err
	}

	err = json.Unmarshal([]byte(str), &User)
	if err != nil {
		return model.User{}, err
	}

	return User, nil
}
