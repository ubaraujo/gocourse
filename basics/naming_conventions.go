package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	/*
		- PascalCase: Structs, interfaces, enums.
			Ex: CalculateArea, UserInfo, NewHTTPRequest

		- snake_case: Variable names, constants, file names.
			Ex: user_id, first_name, http_request

		- UPPERCASE: Constants
			Ex: PI

		- mixedCase: Variables, identificadores específicos, bibliotecas externas.
			Ex: javaScript, htmlDocument, isValid
	*/

	const MAXRETRIES = 5

	var EmployeeID = 1001
	fmt.Println("EmployeeID: ", EmployeeID)

}
