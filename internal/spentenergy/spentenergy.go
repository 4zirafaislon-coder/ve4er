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
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("All parameters must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("All parameters must be positive")
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	averagespeed := distance / duration.Hours()
	return averagespeed
}

func Distance(steps int, height float64) float64 {
	lenghtstep := (height * stepLengthCoefficient) * float64(steps)
	distance := lenghtstep / mInKm
	return distance
}
