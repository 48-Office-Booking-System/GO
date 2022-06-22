package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hallo Gais")

	// websocket function
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {

		log.Println("on connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":3007", nil))
}

/*

1. pindah ke branch masing masing
2. pull dari development (biar dapet yg paling update)
3. ngerjain masing masing kerjaan di branch
4. kalo udah push ke branch masing masing tersebut

//bisa diurus sama ario yang no. 5
5. checkout ke development dan pull dari branch kita tadi lalu push ke development

asdaisbifao
asdoasjdoas
*/
