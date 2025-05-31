package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ruzzt/5sprint/internal/personaldata"
	"github.com/Ruzzt/5sprint/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	datastrings := strings.Split(datastring, ",")
	if len(datastrings) != 3 {
		return fmt.Errorf("the length is less than 3")
	}
	t.Steps, err = strconv.Atoi(datastrings[0])
	if err != nil {
		return fmt.Errorf("invalid steps count: %w", err)
	}
	t.TrainingType = datastrings[1]
	if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {
		return fmt.Errorf("unacceptable type of training: %s", t.TrainingType)
	}

	t.Duration, err = time.ParseDuration(datastrings[2])
	if err != nil {
		return fmt.Errorf("time parsing error: %w", err)
	}

	return nil

}

func (t Training) ActionInfo() (string, error) {
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", fmt.Errorf("неизвестный тип тренировки")
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
			"Сожгли калорий: %.2f",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}
