package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
	runningCaloriesCoefficient = 1.0  // коэффициент для расчета калорий при беге.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 || height <= 0 {
		return 0, errors.New("вес и рост должны быть положительными значениями")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность тренировки должна быть положительной")
	}

	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := weight * speed * minutes / 60
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 {
		return 0, errors.New("The steps must be greater than 0")
	}
	if weight < 0 {
		return 0, errors.New("The weight must be greater than 0")
	}
	if height < 0 {
		return 0, errors.New("The height must be greater than 0")
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	countCalories := speed * weight * minutes / 60
	return runningCaloriesCoefficient * countCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps < 0 {
		return 0
	}
	dis := Distance(steps, height)
	return dis / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	lenStep := stepLengthCoefficient * height
	return (float64(steps) * lenStep) / mInKm
}
