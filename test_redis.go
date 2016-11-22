package main

import (
    "encoding/json"
    "github.com/garyburd/redigo/redis"
    "log"
)

type User struct {
    Id    int32
    Name  string
    Score int32
}

func main() {

    c, err := redis.Dial("tcp", "192.168.99.100:6379")
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // struct to JSON
    user := &User{Id: 1, Name: "name", Score: 2}
    serialized, _ := json.Marshal(user)
    log.Println("serialized : ", string(serialized))

    // set
    c.Do("SET", "test", serialized)

    // get
    data, _ := redis.Bytes(c.Do("GET", "test"))
    log.Println("data : ", data)

    // JSON to struct
    if data != nil {
        deserialized := new(User)
        json.Unmarshal(serialized, deserialized)
        log.Println("deserialized : ", deserialized)
    }

}
