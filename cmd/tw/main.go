package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// Initialize a new tabwriter
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, '\t', 0)

	// Define the table headers
	fmt.Fprintln(tw, "Name\tAge\tOccupation\t")

	// Define the table data
	fmt.Fprintln(tw, "Alice\t23\tEngineer\t")
	fmt.Fprintln(tw, "Bob\t30\tDoctor\t")
	fmt.Fprintln(tw, "Charlie\t35\tTeacher\t")

	// Flush writes any buffered data to the underlying io.Writer
	tw.Flush()
}
