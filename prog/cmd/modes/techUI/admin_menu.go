package techUI

import (
	"bufio"
	"fmt"
	"os"
	"prog/cmd/registry"
	"prog/internal/models"
	"prog/pkg/errors/menuErrors"
	"time"
)

var in = bufio.NewReader(os.Stdin)

func loginAdmin(a *registry.AppServiceFields) (*models.Client, error) {
	var telephone string
	fmt.Printf("Введите телефон: ")
	_, err := fmt.Scanf("%s", &telephone)
	if err != nil {
		return nil, err
	}

	var password string
	fmt.Printf("Введите пароль: ")
	_, err = fmt.Scanf("%s", &password)
	if err != nil {
		return nil, err
	}

	client, err := a.ClientService.Login(telephone, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func createCoach(a *registry.AppServiceFields) {
	var name string
	fmt.Printf("Введите имя: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите описание: ")
	description, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	coach := &models.Coach{Name: name,
		Description: description}

	err = a.CoachService.Create(coach)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nТренер успешно добавлен!\n\n")
}

func createDirection(a *registry.AppServiceFields) {
	var name string
	fmt.Printf("Введите название: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите описание: ")
	description, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	var acceptableGender models.Gender
	fmt.Printf("Введите допустимый пол (0 - не указано, 1 - мужской, 2 - женский): ")
	_, err = fmt.Scanf("%d", &acceptableGender)
	if err != nil {
		fmt.Println(err)
		return
	}

	direction := &models.Direction{Name: name,
		Description:      description,
		AcceptableGender: acceptableGender}

	err = a.DirectionService.Create(direction)
	if err != nil {
		fmt.Println(err)
	}
}

func createHall(a *registry.AppServiceFields) {
	var number uint64
	fmt.Printf("Введите номер: ")
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println(err)
		return
	}

	var capacity uint64
	fmt.Printf("Введите вместительность: ")
	_, err = fmt.Scanf("%d", &capacity)
	if err != nil {
		fmt.Println(err)
		return
	}

	hall := &models.Hall{Number: number,
		Capacity: capacity}

	err = a.HallService.Create(hall)
	if err != nil {
		fmt.Println(err)
	}
}

func addDirectionToCoach(a *registry.AppServiceFields) {
	var coachName string
	fmt.Printf("Введите имя тренера: ")
	_, err := fmt.Scanf("%s", &coachName)
	if err != nil {
		fmt.Println(err)
		return
	}

	coach, err := a.CoachService.GetByName(coachName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var directionName string
	fmt.Printf("Введите название направления: ")
	_, err = fmt.Scanf("%s", &directionName)
	if err != nil {
		fmt.Println(err)
		return
	}

	direction, err := a.DirectionService.GetByName(directionName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = a.CoachService.AddDirection(coach.ID, direction.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createTraining(a *registry.AppServiceFields) {
	printTrainingsOnWeek(a)
	err := printAllDirections(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	var directionName string
	fmt.Printf("Введите название направления: ")
	_, err = fmt.Scanf("%s", &directionName)
	if err != nil {
		fmt.Println(err)
		return
	}

	direction, err := a.DirectionService.GetByName(directionName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = printCoachesByDirection(a, direction)
	if err != nil {
		return
	}

	var coachName string
	fmt.Printf("Введите имя тренера: ")
	_, err = fmt.Scanf("%s", &coachName)
	if err != nil {
		fmt.Println(err)
		return
	}

	coach, err := a.CoachService.GetByName(coachName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var year, day int
	var month time.Month
	fmt.Printf("Введите дату (в формате YYYY.MM.dd): ")
	_, err = fmt.Scanf("%d.%d.%d", &year, &month, &day)
	if err != nil {
		fmt.Println(err)
		return
	}

	times, err := a.CoachService.GetFreeTimeOnDate(coach.ID, time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(times) == 0 {
		fmt.Println("В выбранную дату выбранный тренер занят")
		return
	}

	fmt.Println("Время -> свободные залы: вместительность")
	for _, t := range times {
		hour, _, _ := t.Clock()
		halls, err := a.HallService.GetFreeOnDateTime(time.Date(year, month, day, hour, 0, 0, 0, time.UTC))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(hour)

		if len(halls) == 0 {
			fmt.Println("Нет свободных залов в данное время!")
			return
		}

		for _, h := range halls {
			fmt.Printf("-> %d: %d\n", h.Number, h.Capacity)
		}
	}

	var hour int
	var hallNum, availablePlacesNum uint64
	fmt.Printf("Выберете время, зал и количество человек(через пробел): ")
	_, err = fmt.Scanf("%d %d %d", &hour, &hallNum, &availablePlacesNum)
	if err != nil {
		fmt.Println(err)
		return
	}

	var acceptableAge uint16
	fmt.Printf("Введите допустимый возраст для посещения тренировки: ")
	_, err = fmt.Scanf("%d", &acceptableAge)
	if err != nil {
		fmt.Println(err)
		return
	}

	var name string
	fmt.Printf("Введите название тренировки: ")
	_, err = fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println(err)
		return
	}

	hall, err := a.HallService.GetByNumber(hallNum)
	if err != nil {
		fmt.Println(err)
		return
	}

	training := &models.Training{CoachID: coach.ID,
		HallID:             hall.ID,
		DirectionID:        direction.ID,
		Name:               name,
		DateTime:           time.Date(year, month, day, hour, 0, 0, 0, time.UTC),
		PlacesNum:          availablePlacesNum,
		AvailablePlacesNum: availablePlacesNum,
		AcceptableAge:      acceptableAge}

	err = a.TrainingService.Create(training)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Тренировка успешно создана!")
}

func deleteTraining(a *registry.AppServiceFields) {
	err := printTrainingsOnWeek(a)
	if err != nil {
		return
	}

	var tID uint64
	fmt.Printf("Выберете тренировку и введите её id: ")
	_, err = fmt.Scanf("%d", &tID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = a.TrainingService.Delete(tID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Тренировка успешно удалена!")
}

func createSubscriptionType(subscriptionType int) (*models.Subscription, error) {
	var subscription *models.Subscription

	switch subscriptionType {
	case 1:
		subscription = &models.Subscription{TrainingsNum: 12,
			RemainingTrainingsNum: 12,
			Cost:                  5000,
			StartDate:             time.Now(),
			EndDate:               time.Now().AddDate(0, 1, 0)}
	case 2:
		subscription = &models.Subscription{TrainingsNum: 24,
			RemainingTrainingsNum: 24,
			Cost:                  10000,
			StartDate:             time.Now(),
			EndDate:               time.Now().AddDate(0, 2, 0)}
	case 3:
		subscription = &models.Subscription{TrainingsNum: 36,
			RemainingTrainingsNum: 36,
			Cost:                  15000,
			StartDate:             time.Now(),
			EndDate:               time.Now().AddDate(0, 3, 0)}
	default:
		fmt.Println("Типа с таким номером нет!")
		return nil, menuErrors.ErrorMenu
	}

	return subscription, nil
}

const subscription_types_string = `Типы абонементов:
	1 -- срок действия: 1 месяц  (12 занятий)
	2 -- срок действия: 2 месяца (24 занятий)
	3 -- срок действия: 3 месяца (36 занятий)
`

func createSubscription(a *registry.AppServiceFields) {
	var telephone string
	fmt.Printf("Введите телефон клиента: ")
	_, err := fmt.Scanf("%s", &telephone)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := a.ClientService.GetByTelephone(telephone)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", subscription_types_string)
	fmt.Printf("Выберете тип абонимента: ")
	var subscriptionType int
	_, err = fmt.Scanf("%d", &subscriptionType)
	if err != nil {
		fmt.Println(err)
		return
	}

	subscription, err := createSubscriptionType(subscriptionType)
	if err != nil {
		return
	}

	err = a.SubscriptionService.Create(subscription, client.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Абонемент создан успешно!")
}

func printClient(a *registry.AppServiceFields) {

	var telephone string
	fmt.Printf("Введите телефон клиента: ")
	_, err := fmt.Scanf("%s", &telephone)
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := a.ClientService.GetByTelephone(telephone)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nИмя: %s\nТелефон: %s\nПароль: %s\n\n", client.Name, client.Telephone, client.Password)
}

const admin_loop_string = `Меню администратора: 
	0 -- выйти
	1 -- посмотреть расписание на неделю
	2 -- добавить тренировку
	3 -- удалить тренировку
	4 -- добавить тренера
	5 -- добавить направление
	6 -- добавить тренеру направление
	7 -- посмотреть тренеров по направлению
	8 -- добавить зал
	9 -- добавить клиенту абонемент
Выберите действие: `

func adminMenu(a *registry.AppServiceFields) {
	var num int = 1

	for num != 0 {
		fmt.Printf("\n\n%s", admin_loop_string)

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
			printTrainingsOnWeek(a)
		case 2:
			createTraining(a)
		case 3:
			deleteTraining(a)
		case 4:
			createCoach(a)
		case 5:
			createDirection(a)
		case 6:
			addDirectionToCoach(a)
		case 7:
			printCoaches(a)
		case 8:
			createHall(a)
		case 9:
			createSubscription(a)
		case 10:
			printClient(a)
		default:
			fmt.Printf("\nНеверный пункт меню!\n\n")
		}
	}
}
