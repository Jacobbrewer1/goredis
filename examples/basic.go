// This file shows how the configuration works for the goredis package.

package examples

import (
	"time"

	goredis "github.com/Jacobbrewer1/goredis/redis"
	"github.com/gomodule/redigo/redis"
)

func main() {
	if err := goredis.NewPool(
		goredis.WithDefaultPool(),
		goredis.WithMaxIdle(5),
		goredis.WithMaxActive(10),
		goredis.WithIdleTimeout(int(5*time.Minute)),
		goredis.WithAddress("localhost:6379"),
		goredis.WithNetwork(goredis.NetworkTCP),
		goredis.WithDialOpts(
			redis.DialDatabase(0),
			redis.DialUsername("username"),
			redis.DialPassword("password"),
		),
	); err != nil {
		panic(err)
	}
}
