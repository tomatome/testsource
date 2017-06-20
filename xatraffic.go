// test
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type BusLine struct {
	ROUTEID           string
	ROUTENAME         string
	COMPARELINEID     string
	COMPARELINESERIAL string
	IsLive            bool
}

type BusData struct {
	run_nums     int    //总车辆
	avg_interval int    //平均发车时间
	has_real_bus bool   //真实车辆
	TicketSystem string //票价
	IsSwipe      bool   //可刷卡
	message      string //请求成功
	busLine      BusLine
	inUp         bool
}

func main() {
	resp, err := http.Get("https://www.xaglkp.com.cn/Bus/GetBusLineByName?buslinename=311")
	if err != nil {
		// handle error
		fmt.Println("err:", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("err:", err)
		os.Exit(-1)
	}

	fmt.Println(string(body))

	resp1, err1 := http.Post("https://www.xaglkp.com.cn/Bus/GetRealBusLine",
		"application/x-www-form-urlencoded",
		strings.NewReader("routeid=70"))
	if err1 != nil {
		fmt.Println(err1)
	}

	defer resp1.Body.Close()
	body1, err1 := ioutil.ReadAll(resp1.Body)
	if err1 != nil {
		// handle error
	}

	fmt.Println(string(body1))
}
