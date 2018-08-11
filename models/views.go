package models

import (
	"math/rand"

	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// View struct

type View struct {
	gorm.Model
	Text string
}

// ViewManager struct
type ViewManager struct {
	db *DB
}

// NewViewManager - Create a view manager that can be used for retrieving views
func NewViewManager(db *DB) (*ViewManager, error) {

	db.AutoMigrate(&View{})

	viewmgr := ViewManager{}

	viewmgr.db = db

	return &viewmgr, nil
}

// RandomView - return a random view
func (qm *ViewManager) RandomView() *View {
	view := View{
		Text: views[rand.Intn(len(views))],
	}
	return &view
}

var views = []string{
	"Shanto pagla doesn't call the wrong number. You answer the wrong phone.",
	"Shanto pagla has already been to Mars; that's why there are no signs of life.",
	"Shanto pagla and Superman once fought each other on a bet. The loser had to start wearing his underwear on the outside of his pants.",
	"Some magicans can walk on water, Shanto pagla can swim through land.",
	"Shanto pagla once urinated in a semi truck's gas tank as a joke....that truck is now known as Optimus Prime.",
	"Shanto pagla doesn't flush the toilet, he scares the sh*t out of it",
	"Shanto pagla counted to infinity - twice.",
	"Shanto pagla can cut through a hot knife with butter",
	"Shanto pagla is the reason why Waldo is hiding.",
	"Death once had a near-Shanto pagla experience",
	"When the Boogeyman goes to sleep every night, he checks his closet for Shanto pagla.",
	"Shanto pagla can slam a revolving door.",
	"Shanto pagla once kicked a horse in the chin. Its decendants are known today as Giraffes.",
	"Shanto pagla will never have a heart attack. His heart isn't nearly foolish enough to attack him.",
	"Shanto pagla once got bit by a rattle snake........ After three days of pain and agony ..................the rattle snake died",
	"Shanto pagla can win a game of Connect Four in only three moves.",
	"When Shanto pagla does a pushup, he isn't lifting himself up, he's pushing the Earth down.",
	"There is no theory of evolution. Just a list of animals Shanto pagla allows to live.",
	"Shanto pagla can light a fire by rubbing two ice-cubes together.",
	"Shanto pagla doesn’t wear a watch. HE decides what time it is.",
	"The original title for Alien vs. Predator was Alien and Predator vs Shanto pagla.",
	"The film was cancelled shortly after going into preproduction. No one would pay nine dollars to see a movie fourteen seconds long.",
	"Shanto pagla doesn't read books. He stares them down until he gets the information he wants.",
	"Shanto pagla made a Happy Meal cry.",
	"Outer space exists because it's afraid to be on the same planet with Shanto pagla.",
	"If you spell Shanto pagla in Scrabble, you win. Forever.",
	"Shanto pagla can make snow angels on a concrete slab.",
	"Shanto pagla destroyed the periodic table, because Shanto pagla only recognizes the element of surprise.",
	"Shanto pagla has to use a stunt double when he does crying scenes.",
	"Shanto pagla' hand is the only hand that can beat a Royal Flush.",
	"There is no theory of evolution. Just a list of creatures Shanto pagla has allowed to live.",
	"Shanto pagla does not sleep. He waits.",
	"Shanto pagla tells a GPS which way to go.",
	"Some people wear Superman pajamas. Superman wears Shanto pagla pajamas.",
	"Shanto pagla's tears cure cancer ..... to bad he has never cried",
	"Shanto pagla doesn't breathe, he holds air hostage.",
	"Shanto pagla had a staring contest with Medusa, and won.",
	"When life hands Shanto pagla lemons, he makes orange juice.",
	"When Shanto pagla goes on a picnic, the ants bring him food.",
	"Shanto pagla gives Freddy Krueger nightmares.",
	"They once made a Shanto pagla toilet paper, but there was a problem: It wouldn't take shit from anybody.",
	"Shanto pagla can punch a cyclops between the eyes.",
	"Shanto pagla doesn't mow his lawn, he stands on the porch and dares it to grow",
	"Shanto pagla put out a forest fire. using only gasoline",
	"Shanto pagla CAN believe it's not butter.",
	"Custom t-shirts provided by Spreadshirt",
}
