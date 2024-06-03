package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type lodgers struct {
	Number   int
	Landlord string
	Tenat    string
}

type contract struct {
	Number   int    `json:"number"`
	Landlord string `json:"landlord"`
	Tenat    string `json:"tenat"`
}

type contracts struct {
	List []contract
}

type contractXML struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contractsXML struct {
	List []contractXML `xml:"contract"`
}

type contract1 struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts1 struct {
	List []contract1
}

func main() {
	// Task 13.1
	strJSON := `{"number":1, "landlord":"Остап Бендер", "tenat":"Шура Балаганов"}`
	l := lodgers{}

	err := json.Unmarshal([]byte(strJSON), &l)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	fmt.Printf("%+v", l)
	fmt.Println("")

	// Task 13.2
	c := contract{
		Number:   2,
		Landlord: "Остап",
		Tenat:    "Шура",
	}
	res, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error save JSON: ", err)
	}
	fmt.Printf("%+v", string(res))
	fmt.Println("")

	// Task 13.3
	strXML := `<?xml version="1.0" encoding="UTF-8"?>
		<contracts>
			<contract>
				<number>1</number>
				<sign_date>2023-09-02</sign_date>
				<landlord>Остап</landlord>
				<tenat>Шура</tenat>
			</contract>
			<contract>
				<number>2</number>
				<sign_date>2023-09-03</sign_date>
				<landlord>Бендер</landlord>
				<tenat>Балаганов</tenat>
			</contract>
		</contracts>`

	cXML := contractsXML{}
	err = xml.Unmarshal([]byte(strXML), &cXML)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%+v", cXML)

	// Task 13.4
	/*Необходимо представить в виде xml структуру contracts

	type contracts struct {
		List []contract `xml:"contract"`
	}
	contract{
		Number: 1,
		Landlord: "Остап Бендер",
		tenat: "Шура Балаганов",

	}*/
	fmt.Println("\n")
	c1 := contract1{
		Number:   1,
		Landlord: "Остап Бендер",
		Tenat:    "Шура Балаганов",
	}
	d := contracts1{}
	d.List = append(d.List, c1)

	res, err = xml.MarshalIndent(d, "", " ")
	if err != nil {
		fmt.Println("Error xml: ", err)
	}
	fmt.Printf("%s%s", xml.Header, string(res))

	// Task 13.5
}
