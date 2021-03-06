package xueqiu

import "io"
import "net/http"
import "encoding/json"

var charturi = "/v5/stock/chart/minute.json?period=1d&symbol="

type ChartMessage struct {
	Data Chart
}

type Chart struct {
	//Chart 价格图表结构体
	Last_close float32
	Items      []Item
}

type Item struct {
	//Chart 价格图表item结构体
	Current   float32 //当前价
	Volume    int     //成交量
	Avg_price float32 //均价
	Chg       float32 //涨跌额
	Percent   float32 //涨跌幅度
	Timestamp int64   //时间戳
	Amount    float64
}

func Getchartmessage(symbol string) (ChartMessage, error) {
	var chartmessage ChartMessage
	curl := stockhost + charturi + symbol
	cookie, err := GetCookie()

	if err != nil {
		return chartmessage, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", curl, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return chartmessage, err
	}

	dec := json.NewDecoder(resp.Body)

	if err != nil {
		return chartmessage, err
	}

	for {
		var m ChartMessage
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			return chartmessage, err
		}

		chartmessage = m
	}

	return chartmessage, err
}

//价格获取接口 get quote
func Getchart(symbol string) (Chart, error) {
	chartmessage, err := Getchartmessage(symbol)

	if err != nil {
		return chartmessage.Data, err
	}

	return chartmessage.Data, nil
}
