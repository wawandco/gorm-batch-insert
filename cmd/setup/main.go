// *Requires to have psql installed*
// It is used to setup the database for the exercise
package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Printf("== Setting up Role and Database ==\n")
	roleAndDBCmd := exec.CommandContext(context.Background(), "psql", "-f", "cmd/setup/migrations/role_and_db.sql")
	out, err := roleAndDBCmd.Output()
	if err != nil {
		log.Fatalf("failed to run SQL script to set the role and the database: %s", err.Error())
	}

	fmt.Printf("%s\n", string(out))
	fmt.Printf("== Role and Database ✅ ==\n")

	fmt.Printf("== Setting up Contacts Table ==\n")
	createTableCmd := exec.CommandContext(context.Background(), "psql", "-d", "notebook", "-f", "cmd/setup/migrations/create_contacts_table.sql")
	out, err = createTableCmd.Output()
	if err != nil {
		log.Fatalf("failed to create contacts table in notebook database: %s", err.Error())
	}

	fmt.Printf("%s\n", string(out))
	fmt.Printf("== Created Contacts Table ✅ ==\n")
}
