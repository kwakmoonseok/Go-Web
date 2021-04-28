package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/Web9/cipher"
	"github.com/tuckersGo/goWeb/Web9/lzw"
)
type Component interface {
	Operator(string)
}

var sentData string
var recvData string 

type SendComponent struct {}

func (self *SendComponent) Operator(data string) {
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	value, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(value))
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	value, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(value))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	value, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(value))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	value, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(value))
}

type ReadComponent struct {}
func (self *ReadComponent) Operator(data string) {
	recvData = data
}


func main() {
	sender := &EncryptComponent{
		key: "abcde", 
		com: &ZipComponent{
			com: &SendComponent{}}}

	sender.Operator("Hello world")

	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{}}}
	
	receiver.Operator(sentData)
	fmt.Println(recvData)
}	