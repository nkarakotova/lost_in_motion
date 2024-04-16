package techUI

import (
	"fmt"
	registry "prog/cmd/registry"
)

const menu_string = `Меню гостя:
	0 -- выйти
	1 -- зарегистрироваться
	2 -- авторизироваться как клиент
	3 -- авторизироваться как администратор
	4 -- посмотреть расписание на неделю
	5 -- посмотреть все направления
	6 -- посмотреть тренеров по направлению
Выберите пункт меню: `

func RunMenu(a *registry.AppServiceFields) {

	var num int = 1

	for num != 0 {
		fmt.Printf("\n\n%s", menu_string)

		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Printf("\nПункт меню введён некорректно!\n\n")
			continue
		}

		if num == 0 {
			return
		}

		switch num {
		case 0:
			return
		case 1:
			client, err := createClient(a)
			if err == nil {
				fmt.Printf("\nКлиент успешно добавлен!\n\n")
				clientMenu(a, client)
			}
		case 2:
			client, err := loginClient(a)
			if err == nil {
				fmt.Printf("\nАвторизация успешна!\n\n")
				clientMenu(a, client)
			}
		case 3:
			_, err := loginAdmin(a)
			if err == nil {
				fmt.Printf("\nАвторизация успешна!\n\n")
				adminMenu(a)
			}
		case 4:
			printTrainingsOnWeek(a)
		case 5:
			printAllDirections(a)
		case 6:
			printCoaches(a)
		default:
			fmt.Printf("\nНеверный пункт меню!\n\n")
		}
	}
}
