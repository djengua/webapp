package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// func main() {
// 	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, World! Eliott")
// 	})
// 	fmt.Printf("Server running (port=8081), route: http://localhost:8081/helloworld\n")
// 	if err := http.ListenAndServe(":8081", nil); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("finish")
// }

/// 1 - Example 1
// type EliottHandler struct{}

//	func (h *EliottHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello Eliott!")
//	}
//
// / 1 - Example 1
func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintf(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}

		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
	fmt.Fprintln(w, "Show message")

}

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello world!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
	fmt.Fprintf(w, "Set Message")
}

// func helloEliott(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	fmt.Fprintf(w, "Hello %s! | ", p.ByName("name"))

// 	ranNum := rand.New(rand.NewSource(99))
// 	fmt.Println(fmt.Sprintf("Raffle %d", ranNum.Intn(99)))
// 	rf := raffle.Raffle{
// 		Name:         fmt.Sprintf("Raffle %d", ranNum.Intn(99)),
// 		CreatedAt:    time.Now(),
// 		Id:           ranNum.Intn(99),
// 		Uuid:         "UUIDGEN",
// 		Participants: make([]models.Participant, 0),
// 	}

// 	fmt.Fprintf(w, "Inicia la Rifa: $%s | ", rf.Name)

// 	p1 := models.Participant{
// 		Id:        1,
// 		Uuid:      "David",
// 		Active:    true,
// 		CreatedAt: time.Now(),
// 	}

// 	p2 := models.Participant{
// 		Id:        2,
// 		Uuid:      "Karen",
// 		Active:    true,
// 		CreatedAt: time.Now(),
// 	}

// 	p3 := models.Participant{
// 		Id:        3,
// 		Uuid:      "Eliott",
// 		Active:    true,
// 		CreatedAt: time.Now(),
// 	}

// 	rf.Participants = append(rf.Participants, p1)
// 	rf.Participants = append(rf.Participants, p2)
// 	rf.Participants = append(rf.Participants, p3)

// 	fmt.Println(rf.ToString())

// 	winner := rf.SelectWinner()
// 	fmt.Println("El ganador es: %s", winner.Uuid)
// 	fmt.Fprintf(w, "El ganador es: %s", winner.Uuid)
// }

func main() {
	// mux := httprouter.New()
	// mux.GET("/hello/:name", helloEliott)
	// mux.GET("/set-message", setMessage)
	// mux.GET("/show-message", showMessage)
	// mux.GET("/get-message", helloEliott)

	server := http.Server{
		Addr: "127.0.0.1:8081",
		// Handler: mux,
	}
	// http.HandleFunc("/hello/:name", helloEliott)
	http.HandleFunc("/set-message", setMessage)
	http.HandleFunc("/show-message", showMessage)
	server.ListenAndServe()
}
