package server

import (
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	pb "github.com/im-auld/users/proto"
	"log"
	"os"
	"strconv"
)

type DBClient struct {
	*redis.Client
}

var logger *log.Logger = log.New(os.Stdout, "cLi", log.Lshortfile)

func (db DBClient) GetUser(key string) (*pb.User, bool) {
	user := new(pb.User)
	found := false
	val, err := db.Get(key).Result()
	if err == redis.Nil {
		logger.Println("Not found")
	} else if err != nil {
		logger.Printf("Ooops: %s", err)
	} else {
		logger.Println("found")
		err = proto.Unmarshal([]byte(val), user)
		if err == nil {
			found = true
		}
	}
	return user, found
}

func (db DBClient) SetUser(user *pb.User) {
	value, _ := proto.Marshal(user)
	key := strconv.FormatInt(user.UserId, 10)
	err := db.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func NewDBClient() *DBClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	db := &DBClient{client}
	return db
}
