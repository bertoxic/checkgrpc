package chat

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/net/context"
)

type Server struct {
}

// mustEmbedUnimplementedChatServiceServer implements ChatServiceServer.
func (s *Server) mustEmbedUnimplementedChatServiceServer() {
	//
}

func (s *Server) SayHello(ctx context.Context, message *Request) (*Response, error) {

	log.Printf("receive message request from client %s", message.Request)

	return &Response{Response: fmt.Sprintf("ok hello %s, its nice to finally meet you, \n is this your message'%s'", message.Who, message.Request), Who: "server"}, nil

}

func (s *Server) Subhello(req *Request, servx ChatService_SubhelloServer) error {
	for {
		stars := generateRandomSentence()
		err := servx.Send(&Response{Response: fmt.Sprintf(" %s", stars)})
		if err != nil {
			return err
		}
		awaitedTime := rand.Intn(15)
		//time.Sleep(time.Duration(awaitedTime)* time.Second)
		time.Sleep(time.Duration(awaitedTime) * time.Second)
	}

}

func generateRandomSentence() string {

	// byteArray := []byte{}
	// randword := rand.Intn(6) + 2
	// words := []string{}
	// for i := 0; i < 4; i++ {

	// 	for i := 0; i < randword; i++ {
	// 		randbyte := rand.Intn(25) + 97
	// 		byteArray = append(byteArray, byte(randbyte))
	// 	}
	// 	words = append(words, string(byteArray[:]))
	// 	byteArray = []byte{}
	// }

	words := []string{"apple", "banana", "cherry", "dog", "elephant", "happy", "sun", "moon", "blue", "green"}

	rand.Seed(time.Now().UnixNano())

	numWords := rand.Intn(6) + 3

	sentenceWords := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		randomIndex := rand.Intn(len(words))
		sentenceWords[i] = words[randomIndex]
	}

	sentence := strings.Title(strings.Join(sentenceWords, " ")) + "."

	return sentence
}
