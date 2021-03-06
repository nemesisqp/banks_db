# Banks DB

Community driven database to get bank info (name, brand, color, etc.) by bankcard prefix (BIN)

> This is golang port of [ramoona's banks-db](https://github.com/ramoona/banks-db).

### Install

```
go get -u github.com/KhasanovBI/banks_db/banks_db
```

### Usage

Below is an example which shows some common use cases for banks_db:

```go
package main

import (
	"fmt"
	"github.com/khasanovbi/banks_db/banks_db"
)

func main() {
	for _, creditCard := range []string{"5275940000000000", "4111111111111111"} {
		bank := banks_db.FindBank(creditCard)
		paymentSystem := banks_db.FindPaymentSystem(creditCard)
		fmt.Printf("CreditCard: %s\n", creditCard)
		fmt.Printf("Bank info: %#v\n", bank)
		if paymentSystem != nil {
			fmt.Printf("Payment system: %s\n", *paymentSystem)
		}
		fmt.Println()
	}
}

```

Output:
```
CreditCard: 5275940000000000
Bank info: &banks_db.Bank{Name:"citibank", Country:"ru", LocalTitle:"Ситибанк", EngTitle:"Citibank", URL:"https://www.citibank.ru/", Color:"#0088cf", Prefixes:[]int{419349, 427760, 427761, 520306, 527594}}
Payment system: mastercard

CreditCard: 4111111111111111
Bank info: (*banks_db.Bank)(nil)
Payment system: visa
```