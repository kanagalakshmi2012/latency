package main
import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
)
func setupDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	createTable := `
	CREATE TABLE dns_records (
		hostname TEXT PRIMARY KEY
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		return nil, err
	}

	records := []string{
		"node1.company.local",
		"node2.company.local",
		"node3.company.local",
	}
	for _, hostname := range records {
		_, err = db.Exec("INSERT INTO dns_records(hostname) VALUES (?)", hostname)
		if err != nil {
			return nil, err
		}
	}
return db, nil
}
func dnsLookup(db *sql.DB, hostname string) (bool, float64, error) {
	start := time.Now()
	var result string
	err := db.QueryRow("SELECT hostname FROM dns_records WHERE hostname = ?", hostname).Scan(&result)
	elapsed := time.Since(start).Seconds() * 1000
	if err == sql.ErrNoRows {
		return false, elapsed, nil
	} else if err != nil {
		return false, elapsed, err
	}
	return true, elapsed, nil
}
func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalf("Failed to setup DB: %v", err)
	}
	defer db.Close()
	hostnames := []string{"node1.company.local", "node4.company.local", "node2.company.local"}
	for _, hostname := range hostnames {
		found, duration, err := dnsLookup(db, hostname)
		if err != nil {
			log.Printf("Error looking up %s: %v", hostname, err)
			continue
		}
		if found {
			fmt.Printf("Lookup for %s: Found, Time taken = %.3f ms\n", hostname, duration)
		} else {
			fmt.Printf("Lookup for %s: Not found, Time taken = %.3f ms\n", hostname, duration)
		}
	}
}
