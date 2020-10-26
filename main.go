package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xxmdhs/tiebasign/sign"
)

func main() {
	var ok bool
	var ispush bool
	zanhao := 0
	sw := strings.Builder{}
	sw.WriteString("百度贴吧自动签到\n\n")
	msg := `一共需要给 ` + strconv.Itoa(len(BDUSS)+1) + ` 个账号签到。`
	log.Println(msg)
	sw.WriteString(msg + "\n\n")
	for _, v := range BDUSS {
		ok = false
		zanhao++
	finish:
		for i := 0; i < 3; i++ {
			tbs, err := sign.Getbs(v)
			if err != nil {
				log.Println(err)
				continue
			}
			list, err := sign.GetFollow(v)
			if err != nil {
				log.Println(err)
				continue
			}
			errCh := make(chan string, 10)
			limit := make(chan struct{}, 10)
			msgCh := make(chan string, 10)
			sum := len(list)
			msg := "第" + strconv.Itoa(zanhao) + "个账号需要给" + strconv.Itoa(sum) + "个贴吧签到。"
			sw.WriteString(msg + "\n")
			log.Println(msg)
			if sum == 0 {
				ok = true
				break
			}
			go func() {
				for _, name := range list {
					limit <- struct{}{}
					go toSign(name, v, tbs, errCh, limit, msgCh)
				}
			}()
			var s int
			for {
				select {
				case msg := <-msgCh:
					sw.WriteString(msg + "\n")
					log.Println(msg)
					ispush = true
					s++
					if s == sum {
						ok = true
						break finish
					}
				case err := <-errCh:
					panic(err)
				}
			}
		}
		if !ok {
			panic("签到失败")
		}
		msg := "第" + strconv.Itoa(zanhao) + "个账号签到完成。"
		log.Println(msg)
		sw.WriteString(msg + "\n\n")
	}
	msg = `全部账号签到完成。`
	log.Println(msg)
	sw.WriteString(msg + "\n\n")
	if ispush && tgkey != "" {
		msg := strings.TrimSpace(sw.String())
		var ok bool
		for i := 0; i < 3; i++ {
			err := sign.Push(msg, tgchatID, tgkey)
			if err != nil {
				log.Println(err)
				continue
			}
			ok = true
			break
		}
		if !ok {
			panic("推送失败")
		}
	}
}

func toSign(name, bduss, tbs string, errCh chan<- string, limit <-chan struct{}, msgCh chan<- string) {
	err := sign.Tosign(name, bduss, tbs)
	if err != nil {
		errCh <- err.Error()
	}
	msgCh <- name + "吧签到成功"
	<-limit
}

var (
	BDUSS    = make([]string, 0)
	tgkey    string
	tgchatID string
)

func init() {
	c := os.Getenv("BDUSS")
	err := json.Unmarshal([]byte(c), &BDUSS)
	if err != nil {
		panic(err)
	}
	tgkey = os.Getenv("tgkey")
	tgchatID = os.Getenv("tgchatID")
}
