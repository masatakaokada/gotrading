package main

import (
	"gotrading/app/controllers"
	"gotrading/config"
	"gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}

// apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

// tickerChannel := make(chan bitflyer.Ticker)
// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
// for ticker := range tickerChannel {
// 	fmt.Println(ticker)
// 	fmt.Println(ticker.GetMidPrice())
// 	fmt.Println(ticker.DateTime())
// 	fmt.Println(ticker.TruncateDateTime(time.Second))
// 	fmt.Println(ticker.TruncateDateTime(time.Minute))
// 	fmt.Println(ticker.TruncateDateTime(time.Hour))
// }

// ビットコインの買い注文を出す
// order := &bitflyer.Order{
// 	ProductCode:     config.Config.ProductCode,
// 	ChildOrderType:  "MARKET",
// 	Side:            "BUY",
// 	Size:            0.01,
// 	MinuteToExpires: 1,
// 	TimeInForce:     "GTC",
// }
// res, _ := apiClient.SendOrder(order)
// fmt.Println(res.ChildOrderAcceptanceID)

// 注文一覧を取得
// i := "注文ID"
// params := map[string]string{
// 	"product_code":              config.Config.ProductCode,
// 	"child_order_acceptance_id": i,
// }
// r, _ := apiClient.ListOrder(params)
// fmt.Println(r)

// func main() {
// 	s := models.NewSignalEvents()
// 	df, _ := models.GetAllCandle("BTC_USD", time.Minute, 10)
// 	c1 := df.Candles[0]
// 	c2 := df.Candles[5]
// 	s.Buy("BTC_USD", c1.Time.UTC(), c1.Close, 1.0, true)
// 	s.Sell("BTC_USD", c2.Time.UTC(), c2.Close, 1.0, true)
// 	// 最新の取引１件だけ取得
// 	fmt.Println(models.GetSignalEventsByCount(1))
// 	// 現在時刻以降に取引されたレコードを取得
// 	fmt.Println(s.CollectAfter(time.Now().UTC()))
// 	// c1以降に取引されたレコードを取得
// 	fmt.Println(s.CollectAfter(c1.Time))
// }
