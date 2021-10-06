package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	//"os"
	//"time"
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
	//var data = getData()
	//
	//log.Println("Last Update ID: " + strconv.FormatInt(data.LastUpdateId, 10))
	//log.Println("BIDS:")
	//log.Println("Price          Qty        Sum")
	//log.Println("-----------------------------")
	//
	//for i := 0; i < 15; i++ {
	//	first, _ := strconv.ParseFloat(data.Bids[i][0], 64)
	//	second, _ := strconv.ParseFloat(data.Bids[i][1], 64)
	//	prod := first * second
	//
	//	log.Printf("%v  %s %v", data.Bids[i][0], data.Bids[i][1], prod)
	//}

	var t time.Duration
	t = 20 * time.Millisecond

	for x := range time.Tick(t) {
		var data = getData()
		log.Println(x)
		log.Println("Last Update ID: " + strconv.FormatInt(data.LastUpdateId, 10))
		log.Println("BIDS:")
		log.Println("Price          Qty        Sum")
		log.Println("-----------------------------")

		for i := 0; i < 15; i++ {
			first, _ := strconv.ParseFloat(data.Bids[i][0], 64)
			second, _ := strconv.ParseFloat(data.Bids[i][1], 64)
			prod := first * second

			log.Printf("%v  %s %v", data.Bids[i][0], data.Bids[i][1], prod)
		}
	}

	//writer := uilive.New()       // writer for the first line
	//writer2 := writer.Newline()  // writer for the second line
	//// start listening for updates and render
	//writer.Start()
	//
	//
	//i := 0
	//for true {
	//	fmt.Fprintf(writer, "Downloading File 1.. %d %%\n", i)
	//	fmt.Fprintf(writer2, "Downloading File 2.. %d %%\n", i)
	//	time.Sleep(time.Millisecond * 50)
	//	i++
	//}

	//fmt.Fprintln(writer, "Finished downloading both files :)")
	//writer.Stop() // flush and stop rendering

}
