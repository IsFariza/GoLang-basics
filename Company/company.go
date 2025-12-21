package Company

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Employee interface {
	GetDetails()
	getID() uint64
	setID(id uint64)
}

type Company struct {
	employee      map[uint64]Employee
	incrementedID uint64
}
type FullTimeEmployee struct {
	id         uint64
	firstName  string
	lastName   string
	experience float64
	salary     float64
}

func (fullTimeEmpl *FullTimeEmployee) GetDetails() {
	fmt.Printf("Full Time Employee: %s %s\n", fullTimeEmpl.firstName, fullTimeEmpl.lastName)
	fmt.Printf("- Employee id: %d\n", fullTimeEmpl.id)
	fmt.Printf("- experience in Month: %.2f\n", fullTimeEmpl.experience)
	fmt.Printf("- salary: %f$\n", fullTimeEmpl.salary)
}
func (fullTimeEmpl *FullTimeEmployee) getID() uint64 {
	return fullTimeEmpl.id
}
func (fullTimeEmpl *FullTimeEmployee) setID(id uint64) { fullTimeEmpl.id = id }

type PartTimeEmployee struct {
	id           uint64
	firstName    string
	lastName     string
	hourlySalary float64
	hoursWorked  uint64
}

func (partTimeEmpl *PartTimeEmployee) GetDetails() {
	fmt.Printf("Part Time Employee: %s %s\n", partTimeEmpl.firstName, partTimeEmpl.lastName)
	fmt.Printf("- Employee id: %d\n", partTimeEmpl.id)
	fmt.Printf("- Hourly salary: %.2f\n", partTimeEmpl.hourlySalary)
	fmt.Printf("- Hours Worked this Week: %v\n", partTimeEmpl.hoursWorked)
}
func (partTimeEmpl *PartTimeEmployee) getID() uint64 {
	return partTimeEmpl.id
}
func (partTimeEmpl *PartTimeEmployee) setID(id uint64) { partTimeEmpl.id = id }

func (company *Company) AddEmployee(employee Employee) {

	id := company.incrementedID
	company.incrementedID++
	employee.setID(id)
	company.employee[id] = employee
	fmt.Printf("Employee id %v added successfully\n", id)
}

func (company *Company) ListEmployees() {
	if len(company.employee) == 0 {
		fmt.Println("No employees in Company")
	} else {
		for _, employee := range company.employee {
			employee.GetDetails()
		}
	}
}

func printMenu() {
	fmt.Println("\n=== Menu (enter the number of operation): ===")
	fmt.Println("1. Add Employee ")
	fmt.Println("2. List employee ")
	fmt.Println("3. Exit")
	fmt.Println()
}

func CompanyMenu() {
	company := &Company{
		employee:      make(map[uint64]Employee),
		incrementedID: 1,
	}

	var choice int
	for {
		printMenu()
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Error reading user choice")
			continue
		}
		switch choice {
		case 1:
			employee := company.getEmployeeInfo()
			if employee != nil {
				company.AddEmployee(employee)
			}
		case 2:
			company.ListEmployees()
		case 3:
			return
		default:
			fmt.Println("Error reading user choice")
		}
	}
}
func (company *Company) getEmployeeInfo() Employee {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)
	fmt.Print("Last Name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	var typeEmpl uint8
	fmt.Println("Full Time Employee or Part Time Employee? Enter the number:")
	fmt.Println("1. Full Time")
	fmt.Println("2. Part Time")
	fmt.Scan(&typeEmpl)
	switch typeEmpl {
	case 1:
		var experience, salary float64
		fmt.Print("experience in Months: ")
		fmt.Scan(&experience)
		fmt.Print("salary: ")
		fmt.Scan(&salary)

		return &FullTimeEmployee{
			firstName:  firstName,
			lastName:   lastName,
			experience: experience,
			salary:     salary,
		}
	case 2:
		var hourlySalary float64
		var hoursWorked uint64
		fmt.Print("Hourly salary: ")
		fmt.Scan(&hourlySalary)
		fmt.Print("Hours Worked this Week: ")
		fmt.Scan(&hoursWorked)

		return &PartTimeEmployee{
			firstName:    firstName,
			lastName:     lastName,
			hourlySalary: hourlySalary,
			hoursWorked:  hoursWorked,
		}
	default:
		fmt.Println("Invalid type")
		return nil
	}

}
