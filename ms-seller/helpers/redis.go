package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-seller/model"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetProductCache(c *redis.Client, product *model.Product, ctx context.Context) error {
	p, err := json.Marshal(product)
	if err != nil {
		return err
	}

	key := "product" + strconv.Itoa(product.ID)

	err = c.Set(ctx, key, p, 1*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetProduct(c *redis.Client, productID int, ctx context.Context) (model.Product, error) {
	var product model.Product
	key := "product" + strconv.Itoa(productID)

	str, err := c.Get(ctx, key).Result()
	if err == redis.Nil {
		return model.Product{}, fmt.Errorf("product not found")
	}
	if err != nil {
		return model.Product{}, err
	}

	err = json.Unmarshal([]byte(str), &product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
