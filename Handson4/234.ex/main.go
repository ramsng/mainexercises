package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type structerr struct {
	error1 string
	error2 error
}

func (s structerr) Error() string {
	return fmt.Sprintf("Fatal error : %v %v", s.error1, s.error2)
}

type person struct {
	First   string
	Last    string
	Sayings string
}

func main() {
	p1 := person{
		"James", "Bond", "No time to die",
	}
	//	var errora structerr
	//fatalerr := fmt.Errorf(" Fatal Error: %v ; input %v ", error, p1)
	errorz := jsonmarshal(p1)
	fmt.Println(errorz)
	//	func()(error){
}

func jsonmarshal(p1 person) error {
	if json1, err := json.Marshal(p1); err != nil {
		errora := fmt.Errorf("Marshal data : ")
		errorb := fmt.Sprint(p1.First, p1.Last)
		return structerr{errorb, errora}
	} else {
		if errorc := os.WriteFile("output.txt", []byte(json1), 0777); errorc == nil {
			//return errors.New("Error writting file; error code")
			return fmt.Errorf("Error writting file; error code : %v", errorc)
		} else {
			return nil
		}
	}

}
