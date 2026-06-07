package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds DaySteps) Print() {
	ds.Personal.Print()
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")

	if len(slice) != 2 {
		return errors.New("need 2 elements")
	}

	steps, err := strconv.Atoi(slice[0])

	if err != nil {
		return err
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(slice[1])

	if err != nil {
		return err
	}

	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories)

	return info, nil
}
