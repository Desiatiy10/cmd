package main

import "fmt"

func main() {
	accountingBooks := map[string]map[string][]string{
		"Anton M.": {
			"books":     {"Metro 2033", "Chack Palanick", "Computer Science"},
			"newspaper": {},
		},
		"Sergey L.": {
			"books":     {},
			"newspaper": {"WP", "RT", "some trash"},
		},
		"Mikhalil K.": {},
	}

	count := 0

	for _, categories := range accountingBooks {
		hasBookOrPaper := false

		for categori, items := range categories {
			if categori != "" || len(items) > 0 {
				hasBookOrPaper = true
				break
			}
		}
		if hasBookOrPaper {
			count++
		}
	}
	
	fmt.Printf("Кол-во читателей с изданиями на руках: %d\n ", count)

	for reader, categories := range accountingBooks {
		total := 0

		for _, items := range categories {
			total += len(items)
		}
		fmt.Printf("Читатель: %s Колличество изданий на руках: %d\n", reader, total)
	}
}
