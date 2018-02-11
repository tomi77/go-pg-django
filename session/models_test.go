package session

import (
	"fmt"
	"time"
)

func ExampleSession_String() {
	session := Session{
		SessionKey:  "4eu6u2bygainvm9ihefne3zo20n2tjt6",
		SessionData: "qaz123",
		ExpireDate:  time.Now(),
	}

	fmt.Println(session)
	// Output:
	// 4eu6u2bygainvm9ihefne3zo20n2tjt6
}
