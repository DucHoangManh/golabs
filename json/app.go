package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Job string `json:"job"`
	City string `json:"city"`
	Salary int `json:"salary"`
	Birthday string `json:"birthdate"`
}
func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}
func main() {
	jsonFile, err := os.Open("persons.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var persons []Person
	json.Unmarshal(byteValue, &persons)
	// for i:=0; i<10; i++ {
	// 	fmt.Println(persons[i].Birthday,getAge(persons[i].Birthday))
	// }
	fmt.Println(TopJobsByNumerInEachCity(persons))
}