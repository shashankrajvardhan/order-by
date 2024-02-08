package main

import (
	"fmt"
	"log"
)

func join() {
	query := `
	SELECT COUNT (customerid), customername
	FROM customer
	GROUP BY customername
	ORDER BY COUNT(customerid) DESC, customername asc;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			customerid   int
			customername string
		)

		err := rows.Scan(&customerid, &customername)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("customerid: %d, customername: %s\n", customerid, customername)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
