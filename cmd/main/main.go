package main

import (
	"log"
)

// func Do(ctx context.Context, client *ent.Client) error {
// 	a8m, err := client.User.
// 		Create().
// 		SetAge(30).
// 		SetName("Mashraki").
// 		Save(ctx)
// 	if err != nil {
// 		return fmt.Errorf("creating user: %w", err)
// 	}

// 	log.Println("user:", a8m)

// 	card1, err := client.Card.
// 		Create().
// 		SetOwner(a8m).
// 		SetNumber("1020").
// 		SetExpired(time.Now().Add(time.Minute)).
// 		Save(ctx)
// 	if err != nil {
// 		return fmt.Errorf("creating card: %w", err)
// 	}

// 	log.Println("card:", card1)

// 	// Only returns the card of the user,
// 	// and expects that there's only one.
// 	card2, err := a8m.QueryCard().Only(ctx)
// 	if err != nil {
// 		return fmt.Errorf("querying card: %w", err)
// 	}

// 	log.Println("card:", card2)

// 	// The Card entity is able to query its owner using
// 	// its back-reference.
// 	owner, err := card2.QueryOwner().Only(ctx)

// 	if err != nil {
// 		return fmt.Errorf("querying owner: %w", err)
// 	}

// 	log.Println("owner:", owner)

// 	return nil
// }

func main() {
	// Do(context.Background(), driver.MysqlClient)
	log.Println("owner:")
}
