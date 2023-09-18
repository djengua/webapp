package raffle

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/djengua/webapp/models"
)

type Raffle struct {
	Id           int
	Uuid         string
	Name         string
	Participants []models.Participant
	CreatedAt    time.Time
}

// func (r Raffle) CreateRaffle() {
// 	ranNum := rand.New(rand.NewSource(99))
// 	// return Raffle{
// 	// 	Name:         fmt.Sprintf("Raffle %d", ranNum),
// 	// 	CreatedAt:    time.Now(),
// 	// 	Id:           ranNum.Int(),
// 	// 	Uuid:         "UUIDGEN",
// 	// 	Participants: make([]models.Participant, 0),
// 	// }
// 	r.Name = fmt.Sprintf("Raffle %d", ranNum)
// 	r.CreatedAt = time.Now()
// 	r.Id = ranNum.Int()
// 	r.Uuid = "UUIDGEN"
// 	r.Participants = make([]models.Participant, 0)
// }

func (r Raffle) AddParticipants(p models.Participant) {
	r.Participants = append(r.Participants, p)
}

func (r Raffle) SelectWinner() models.Participant {
	fmt.Println("len")
	fmt.Println(len(r.Participants))
	winner := rand.Intn(len(r.Participants))
	fmt.Println("winner")
	fmt.Println(winner)
	return r.Participants[winner]
}

func (r Raffle) ToString() string {
	return fmt.Sprintf("{ id: %s, name: %s, length: %d }", r.Name, len(r.Participants))
}
