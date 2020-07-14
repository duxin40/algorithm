package rpc

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	Write()
	fmt.Println("write file successfully !")
}

func TestRead(t *testing.T) {
	x := Read()
	fmt.Println(fmt.Sprintf("get data: %+v", x.String()))
}
