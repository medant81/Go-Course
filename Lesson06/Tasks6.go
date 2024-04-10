package main

import (
	"fmt"
)

// For Task 6.1 and 6.3
type contract struct {
	ID     int
	Number string
	Date   string
}

// For Task 6.3
func (c contract) print() string {
	return fmt.Sprintf("Договор № %s от %s", c.Number, c.Date)
}

func main() {

	// Task 6.1
	c := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v\n", c)

	// Task 6.2
	type contractLocal struct {
		ID     int
		Number string
		Date   string
	}

	cLocal := contractLocal{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}

	fmt.Printf("%+v\n", cLocal)

	// Task 6.3
	fmt.Printf("%s\n", c.print())

	// Task 6.4
	type contacts struct {
		Addss string
		Phone string
	}
	type user struct {
		ID   int
		Name string
		contacts
	}

	type employee struct {
		ID   int
		Name string
		contacts
	}

	u := user{
		ID:   1,
		Name: "Антон",
		contacts: contacts{
			Addss: "Планета земля",
			Phone: "112211",
		},
	}

	e := employee{
		ID:   2,
		Name: "Автомакон",
		contacts: contacts{
			Addss: "Москва",
			Phone: "221122",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)

}
