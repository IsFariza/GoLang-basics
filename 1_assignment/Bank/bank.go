package Bank

import "fmt"

type BankAccount struct {
	balance      float64
	transactions []string
}

func (account *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Amount cannot be negative or zero")
		return
	}
	account.balance += amount
	account.transactions = append(account.transactions,
		fmt.Sprintf(" - Deposit (+%v): %.2f\n", amount, account.balance))
	fmt.Println("Deposit successfully")
}
func (account *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Amount cannot be negative or zero")
		return
	}
	if amount > account.balance {
		fmt.Println("Your bank account does not have enough balance")
		return
	}
	account.balance -= amount
	account.transactions = append(account.transactions,
		fmt.Sprintf(" - Withdraw (-%v): %.2f\n", amount, account.balance))
	fmt.Println("Withdraw successfully")
}
func (account *BankAccount) GetTransactions() {
	if len(account.transactions) == 0 {
		fmt.Println("No transactions")
		return
	}
	fmt.Println("Transactions:")
	for _, transaction := range account.transactions {
		fmt.Print(transaction)
	}
}
func (account *BankAccount) GetBalance() {
	fmt.Printf("Balance: %.2f\n", account.balance)
}
func BankAccountMenu() {
	account := &BankAccount{
		balance:      0,
		transactions: []string{},
	}
	var choice int
	for {
		printMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			account.GetBalance()
		case 2:
			var amount float64
			fmt.Println("Enter amount to Deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 3:
			var amount float64
			fmt.Println("Enter amount to Withdraw: ")
			account.GetBalance()
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 4:
			account.GetTransactions()
		case 5:
			return
		default:
			fmt.Println("Invalid choice")

		}
	}
}
func printMenu() {
	fmt.Println("\n=== Menu (enter the number of operation): ===")
	fmt.Println("1. Get Current Balance ")
	fmt.Println("2. Deposit ")
	fmt.Println("3. Withdraw ")
	fmt.Println("4. Get Transactions ")
	fmt.Println("5. Exit")
	fmt.Println()
}
