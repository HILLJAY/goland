package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client01() {

	dial, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {

		fmt.Println(err)
	} else {

		fmt.Println("连接成功", dial)

		//从终端中读取数据

		reader := bufio.NewReader(os.Stdin)

		for {
			readString, _ := reader.ReadString('\n')

			if err != nil {

				fmt.Println("读取终端问题错误")

			} else {

				_, err := dial.Write([]byte(readString))

				if readString=="exit" {

					break
				}

				if err != nil {

					fmt.Println(err)
				}

			}
		}

	}

}

/**
他 在 1 9 6 2 年的 “ S e n s o r am a S im u l a t o r , , 专利 已具有一 定的 v R 技术的思想 川 . 电子计算技术的发展和 计 算机的小型化 , 推动了仿真技术的发展 , 逐步形成了 计算机仿真科学技术学 科 . 1 9 6 5 年 , 计算机 图形学的重要 奠基人 S ut h er - 一a n d 博士发表了一篇短文 “ T h e u l t im a t e d i s p l盯 , , , 以其敏锐的洞 察力和 丰富的想象力描绘了一种新的 显示 技术 . 他设想在这种显示 技术支持下 , 观察者可以直接沉浸在计算机控制的虚拟 环境之 中 , 就如同 日 常生活在真实世界一样 . 同时 , 观察者还能以 自然的 方式 与虚拟环境中的对象进行交互 , 如触摸感知和 控 制虚拟 对象等 . S ut he lr an d 的文章从计算机显示和人 机交互的角度提出了模拟现实世界的思想 , 推动了计 算机图形 图像技术的发展 , 并启发了头盔显示器 、 数 据手套等新型人机交互设备的研究 。
*/