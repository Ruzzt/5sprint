package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	datastrings := strings.Split(datastring, ",")
	if len(datastrings) != 3 {
		return fmt.Errorf("the length is less than 3")
	}
	t.Steps, err = strconv.Atoi(datastrings[0])
	if err != nil {
		return fmt.Errorf("invalid steps count: %w", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	t.TrainingType = datastrings[1]

	t.Duration, err = time.ParseDuration(datastrings[2])
	if err != nil {
		return fmt.Errorf("time parsing error: %w", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	distance := spentenergy.Distance(t.Steps, t.Height)

	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}

func (t Training) Print() {
	fmt.Printf("Имя: %s\n", t.Name)
	fmt.Printf("Вес: %.2f кг.\n", t.Weight)
	fmt.Printf("Рост: %.2f м.\n\n", t.Height)
}
