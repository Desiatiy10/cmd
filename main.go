package main

import "fmt"

func main() {
	week := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	weekend := week[5:7]
	fmt.Println("Выходные: ", weekend)
	week = week[:5]
	fmt.Println("Рабочие дни: ", week)

	week = append(week, weekend...)
	fmt.Println("Полная неделя: ", week)
}
