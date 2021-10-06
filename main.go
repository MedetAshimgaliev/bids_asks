package main

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uilive"
	"net/http"
	"strconv"
	"time"
)

type DataModel struct {
	LastUpdateId int64
	Bids         [][]string
	Asks         [][]string
}

func getData() DataModel {
	var url = "https://api.binance.com/api/v3/depth?symbol=ETHBUSD&limit=20"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return DataModel{}
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var data DataModel
	err = decoder.Decode(&data)
	return data
}

func printBids() {
	writer := uilive.New()
	writer2 := writer.Newline()
	writer.Start()

	var uID int64
	for true {
		var data = getData()
		uID = data.LastUpdateId

		fmt.Fprintf(writer, "Last Update ID:  %d \n", uID)

		for i := 0; i < 15; i++ {
			first, _ := strconv.ParseFloat(data.Bids[i][0], 64)
			second, _ := strconv.ParseFloat(data.Bids[i][1], 64)
			prod := first * second
			fmt.Fprintf(writer2, "#%d: Price: %s | Qty: %s | Sum: %v \n", i, data.Bids[i][0], data.Bids[i][1], prod)
		}
		time.Sleep(time.Millisecond * 1000)
	}
	fmt.Fprintln(writer, "Finished")
	writer.Stop()
}

const socketAddress string = "wss://testnet.binancefuture.com"

//func initWebsocketClient() {
//	fmt.Println("Starting Client")
//	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/ws", address), "", fmt.Sprintf("http://%s/", address))
//	if err != nil {
//		fmt.Printf("Dial failed: %s\n", err.Error())
//		os.Exit(1)
//	}
//	incomingMessages := make(chan string)
//	go readClientMessages(ws, incomingMessages)
//	i := 0
//	for {
//		select {
//		case <-time.After(time.Duration(2e9)):
//			i++
//			response := new(Message)
//			response.RequestID = i
//			response.Command = "Eject the hot dog."
//			err = websocket.JSON.Send(ws, response)
//			if err != nil {
//				fmt.Printf("Send failed: %s\n", err.Error())
//				os.Exit(1)
//			}
//		case message := <-incomingMessages:
//			fmt.Println(`Message Received:`,message)
//
//
//		}
//	}
//}
//
//func readClientMessages(ws *websocket.Conn, incomingMessages chan string) {
//	for {
//		var message string
//		// err := websocket.JSON.Receive(ws, &message)
//		err := websocket.Message.Receive(ws, &message)
//		if err != nil {
//			fmt.Printf("Error::: %s\n", err.Error())
//			return
//		}
//		incomingMessages <- message
//	}
//}

func main() {
	fmt.Println("*****BIDS & ASKS*****")
	printBids()
}
