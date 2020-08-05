package database

import (
	"database/sql"
)

// CreateNewUser creates a new user, who starts at level 1.
func (db *DB) CreateNewUser(uuid string, name string) (sql.Result, error) {
	return db.Exec("insert into duser values($1,$2)", uuid, name)
}

// GetUser gets the details of a user
func (db *DB) GetUser(currUser *User, uuid string) error {
	return db.Get(currUser, "select name from duser where id = $1", uuid)
}

// // GetHints gets the hints released for a question
// func (db *DB) GetHints(currLev int, hints *[]string) error {
// 	return db.Select(hints, "select content from hints where number = $1", currLev)
// }

// // LogAnswerAttempt logs every answer attempt
// func (db *DB) LogAnswerAttempt(uuid string, currUser User, answer string) (sql.Result, error) {
// 	return db.Exec("insert into answer_logs values($1, $2, $3, $4)", uuid, currUser.Name, answer, time.Now())
// }

// // GetCorrectAns gets the correct answer for a level from the DB to check if the user's attempt is correct
// func (db *DB) GetCorrectAns(currUser User, correctAns *string) error {
// 	return db.Get(correctAns, "select answer from levels where number = $1", currUser.CurrLevel)
// }

// // CorrectAnswerSubmitted increments the user level on submission of correct answer
// func (db *DB) CorrectAnswerSubmitted(uuid string) (sql.Result, error) {
// 	return db.Exec("update duser set curr_level = curr_level + 1 where id = $1", uuid)
// }

// GetLeaderboard gets the users list in the descending order of level,
// and for users on the same level, in the ascending order of last submission time.
func (db *DB) GetLeaderboard(users *[]User) error {
	return db.Select(users, "select name, curr_level from duser order by curr_level desc, last_anstime")
}
