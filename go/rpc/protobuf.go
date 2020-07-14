package rpc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func Write() {
	// MessageEnvelope是models.pb.go的结构体
	oldData := &MessageEnvelope{
		TargetId: 1,
		ID:       "1",
		Type:     "2",
		Payload:  []byte("ITDragon protobuf"),
	}

	data, err := proto.Marshal(oldData)
	if err != nil {
		fmt.Println("marshal error: ", err.Error())
	}
	fmt.Println("marshal data : ", data)

	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func Read() *MessageEnvelope {
	data, _ := ioutil.ReadFile("./test.txt")
	newData := &MessageEnvelope{}
	err := proto.Unmarshal(data, newData)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	return newData

}
