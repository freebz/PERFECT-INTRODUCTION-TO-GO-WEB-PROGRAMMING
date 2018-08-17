// 별칭

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/lann/squirrel"
)

func findUser(identifier string) (*User, error) {
	var (
		id   int
		name string
	)

	err := sq.Select("id, name").From("users").
		Where(sq.Eq{"identifier": identifier}).
		RunWith(db()).QueryRow().Scan(&id, &name)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Username is %s\n", username)
	}

	u := &User{id, name, identifier, avatar_revision.String}
	return u, nil
}
