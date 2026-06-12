package trainings

import (
	"errors"
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

func (t Training) Print() {
	t.Personal.Print()
}

func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")

	if len(slice) != 3 {
		return errors.New("need 3 elements")
	}

	steps, err := strconv.Atoi(slice[0])

	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("need steps > 0")
	}

	t.Steps = steps

	t.TrainingType = slice[1]

	duration, err := time.ParseDuration(slice[2])

	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("need duration > 0")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error

	distance := spentenergy.Distance(t.Steps, t.Height)

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories,
	)

	return info, nil
}
