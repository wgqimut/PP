
/*
Write a program that displays temperature conversion tables.
The tables should use equals signs (=) and pipes (|) to draw lines,
with a header section:
 =======================
 | oC       | oF       |
 =======================
 | -40.0    | -40.0    |
 | ...      | ...      |
 =======================
The program should draw two tables. The first table has two columns,
with C in the first column and F in the second column.
Loop from –40 C through 100 C in steps of 5 using the temperature
conversion methods from lesson 13 to fill in both columns.
After completing one table, implement a second table with the columns
reversed, con- verting from F to C.
Drawing lines and padding values is code you can reuse for any data
that needs to be displayed in a two-column table. Use functions to separate
the table drawing code from the code that calculates the temperatures
for each row.
Implement a drawTable function that takes a first-class function as a
parameter and calls it to get data for each row drawn. Passing a different
function to drawTable should result in different data being displayed.
*/
package main

import "fmt"

const (
	header   = "======================="
	format   = "| %-9s| %-9s|\n"
	tempForm = "| %-9.1f| %-9.1f|"
)

type celsius float64

type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

type getRowFn func(row int) string

func ctof(row int) string {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()

	return fmt.Sprintf(tempForm, c, f)
}

func ftoc(row int) string {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()

	return fmt.Sprintf(tempForm, f, c)
}

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(header)
	fmt.Printf(format, hdr1, hdr2)
	fmt.Println(header)
	for i := 0; i < rows; i++ {
		fmt.Println(getRow(i))
	}
	fmt.Println(header)
}

func main() {
	drawTable("°C", "°F", 30, ctof)
	fmt.Println()
	drawTable("°F", "°C", 29, ftoc)
}
