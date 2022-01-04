package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"work-place/src/chatroom/common"
)

/**
声明一个转换结构体，使其与下面方法挂在一起
*/
type TransFer struct {
	Conn net.Conn
	Buf  [8024]byte
}

func (this *TransFer) HandleRequest() (mes common.Message, err error) {

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {

		return
	}

	//获取packageLength
	var pkgLength uint32
	pkgLength = binary.BigEndian.Uint32(this.Buf[:4])

	//根据packageLength获取信息
	n, err := this.Conn.Read(this.Buf[:pkgLength])
	if uint32(n) != pkgLength || err != nil {

		return
	}

	//序列化信息，封装为一个mes,注意这里一定要穿地址值，如果没传则默认为值拷贝，无法修改message
	err = json.Unmarshal(this.Buf[:pkgLength], &mes)
	if err != nil {
		return
	}

	return
}

func (this *TransFer) HandleResponse(mesBytes []byte) (err error) {

	//1.先获取数据长度
	var pkgLen uint32
	pkgLen = uint32(len(mesBytes))

	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)

	//2.先发送数据长度
	n, err := this.Conn.Write(this.Buf[:4])

	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//3.发送数据本身
	_, err = this.Conn.Write(mesBytes)

	if err != nil {
		return err
	}

	return
}
