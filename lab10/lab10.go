package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"context"
	"unicode/utf8"

	"github.com/gorilla/websocket"
	"github.com/reactivex/rxgo/v2"
)

type client chan<- string // an outgoing message channel

var (
	entering      = make(chan client)
	leaving       = make(chan client)
	messages      = make(chan rxgo.Item) // all incoming client messages
	ObservableMsg = rxgo.FromChannel(messages)
	swearWords []string
	sensativeName []string
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	MessageBroadcast := ObservableMsg.Observe()
	for {
		select {
		case msg := <-MessageBroadcast:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg.V.(string)
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func clientWriter(conn *websocket.Conn, ch <-chan string) {
	for msg := range ch {
		conn.WriteMessage(1, []byte(msg))
	}
}

func wshandle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "你是 " + who + "\n"
	messages <- rxgo.Of(who + " 來到了現場" + "\n")
	entering <- ch

	defer func() {
		log.Println("disconnect !!")
		leaving <- ch
		messages <- rxgo.Of(who + " 離開了" + "\n")
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		messages <- rxgo.Of(who + " 表示: " + string(msg))
	}
}

func init() {
	file, err := os.Open("swear_word.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		swearWords = append(swearWords, scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}

	file2, err := os.Open("sensitive_name.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file2.Close()

	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		sensativeName = append(sensativeName, scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}
}

func InitObservable() {
	// TODO: Please create an Observable to handle the messages
	/*
		ObservableMsg = ObservableMsg.Filter(...) ... {
		}).Map(...) {
			...
		})
	*/
	ObservableMsg = rxgo.FromChannel(messages).Filter(func(item interface{}) bool {
        msg := item.(string)
        for _, word := range swearWords {
            if strings.Contains(msg, word) {
                return false
            }
        }
        return true
    }).Map(func(_ context.Context, item interface{}) (interface{}, error) {
        msg := item.(string)
        for _, name := range sensativeName {
            if strings.Contains(msg, name) {
                // Count the number of characters (runes) in the sensitive name
                runeCount := utf8.RuneCountInString(name)
                if runeCount > 1 {
                    // Convert the sensitive name to a slice of runes
                    runes := []rune(name)
                    // Replace the second character with "*"
                    runes[1] = '*'
                    // Convert the slice of runes back to a string
                    replacement := string(runes)
                    msg = strings.ReplaceAll(msg, name, replacement)
                }
            }
        }
        return msg, nil
    })
}


func main() {
	InitObservable()
	go broadcaster()
	http.HandleFunc("/wschatroom", wshandle)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("server start at :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}