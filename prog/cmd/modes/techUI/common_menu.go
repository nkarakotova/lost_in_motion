package techUI

import (
	"prog/cmd/registry"
	"prog/internal/models"
	"prog/pkg/errors/menuErrors"
	"fmt"
	"time"
)

func printTrainingsOnWeek(a *registry.AppServiceFields) error {
	trainings, err := a.TrainingService.GetAllBetweenDateTime(time.Now(), time.Now().AddDate(0, 0, 7))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(trainings) == 0 {
		fmt.Println("На неделю ещё не поставлены тренеровки!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренировки на неделе:")
	for _, t := range trainings {
		fmt.Printf("\nid: %d\nНазвание: %s\nДата и время: %s\n\n", t.ID, t.Name, t.DateTime)
	}

	return nil
}

func printAllDirections(a *registry.AppServiceFields)  error {
	directions, err := a.DirectionService.GetAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(directions) == 0 {
		fmt.Println("Направления ещё не добавлены!")
		return menuErrors.ErrorMenu
	}

	var gender string

	fmt.Printf("Все направления: ")
	for _, d := range directions {
		switch d.AcceptableGender {
		case models.Unknown:
			gender = "любой"
		case models.Male:
			gender = "мужской"
		case models.Female:
			gender = "женский"
		default:
			fmt.Println("Некорректно заданный пол!")
			return menuErrors.ErrorMenu
		}

		fmt.Printf("\nНазвание: %s\nОписание: %s\nДопустимый пол: %s\n\n", d.Name, d.Description, gender)
	}
	return nil
}

func getDirection(a *registry.AppServiceFields) (*models.Direction, error) {
	printAllDirections(a)
	var directionName string
	fmt.Printf("Введите название направления: ")
	_, err := fmt.Scanf("%s", &directionName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	direction, err := a.DirectionService.GetByName(directionName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return direction, nil
}

func printCoaches(a *registry.AppServiceFields) error {
	direction, err := getDirection(a)
	if err != nil {
		return err
	}

	coaches, err := a.CoachService.GetAllByDirection(direction.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(coaches) == 0 {
		fmt.Println("По данному направлению нет тренеров!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренера по данному направлению: ")
	for _, c := range coaches {
		fmt.Printf("\nИмя: %s\nОписание: %s\n\n", c.Name, c.Description)
	}

	return nil
}

func printCoachesByDirection(a *registry.AppServiceFields, direction *models.Direction) error {
	coaches, err := a.CoachService.GetAllByDirection(direction.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(coaches) == 0 {
		fmt.Println("По данному направлению нет тренеров!")
		return menuErrors.ErrorMenu
	}

	fmt.Printf("Тренера по данному направлению: ")
	for _, c := range coaches {
		fmt.Printf("\nИмя: %s\nОписание: %s\n\n", c.Name, c.Description)
	}

	return nil
}