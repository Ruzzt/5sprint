package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ruzzt/5sprint/internal/personaldata"
	"github.com/Ruzzt/5sprint/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных: ожидается формат 'шаги,продолжительность'")
	}

	// Парсим количество шагов
	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга количества шагов: %w", err)
	}
	if ds.Steps < 0 {
		return fmt.Errorf("количество шагов не может быть отрицательным")
	}

	// Парсим продолжительность
	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка парсинга продолжительности: %w", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Вычисляем количество сожженных калорий
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	// Формируем строку с информацией
	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}

func (ds DaySteps) Print() {
	fmt.Printf("Имя: %s\n", ds.Name)
	fmt.Printf("Вес: %.2f кг.\n", ds.Weight)
	fmt.Printf("Рост: %.2f м.\n\n", ds.Height)
}
