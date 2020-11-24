package usecase

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/calendar"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
	"strings"
)

type UserUsecase interface {
	Create(name string) error
	FindByUid(uid string) (*model.User, error)
	GetSchedule(uid string) (*ResponseUserGetSchedule, error)
	UpdateSchedule(schedule *RequestUserUpdateSchedule) error
	ChangeName(uid string, newName string) (bool, error)
	Delete(uid string) error
}

type userUsecase struct {
	UserRepository repository.UserRepository
}

type RequestUserUpdateSchedule struct {
	uid  string `json:"uid"`
	week []day  `json:"week"`
}

type ResponseUserGetSchedule struct {
	week []day `json:"week"`
}

type day struct {
	date      string      `json:"date"`
	timetable []timetable `json:"timetable"`
}

type timetable struct {
	flag  bool   `json:"flag"`
	times string `json:"times"`
}

// コンストラクタ
func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepository: userRepository}
}

func (uu *userUsecase) Create(name string) error {
	return uu.UserRepository.Insert(name)
}

func (uu *userUsecase) FindByUid(uid string) (*model.User, error) {
	return uu.UserRepository.FindByUid(uid)
}

func (uu *userUsecase) GetSchedule(uid string) (*ResponseUserGetSchedule, error) {
	user, err := uu.UserRepository.FindByUid(uid)
	if err != nil {
		return nil, err
	}

	weekList, holidayFlagList := calendar.GetCalendar()
	weekdayList := []string{"木", "金", "土", "日", "月", "火", "水"}

	schedule := new(ResponseUserGetSchedule)
	for i, weekday := range user.Schedule {
		d := new(day)
		for ii := 0; ii < len(weekday.Flags); ii++ {
			tt := new(timetable)
			if weekday.Flags[ii:ii+1] == "0" {
				tt.flag = false
				d.timetable = append(d.timetable, *tt)
			} else {
				tt.flag = true
				d.timetable = append(d.timetable, *tt)
			}
		}
		schedule.week = append(schedule.week, *d)

		s := strings.Split(weekList[i], "-")
		// s[0]:2020, s[1]:11, s[2]:24
		schedule.week[i].date = s[1] + "/" + s[2] + " (" + weekdayList[i] + ")"

		if holidayFlagList[i] == 0 {
			schedule.week[i].timetable[0].times = "9:00-10:30"
			schedule.week[i].timetable[1].times = "10:30-12:00"
			schedule.week[i].timetable[2].times = "12:00-14:00"
			schedule.week[i].timetable[3].times = "14:00-16:00"
			schedule.week[i].timetable[4].times = "16:00-17:30"
			schedule.week[i].timetable[5].times = "17:30-19:00"
			schedule.week[i].timetable[6].times = "19:00-20:30"
			schedule.week[i].timetable[7].times = "20:30-22:00"
		} else {
			schedule.week[i].timetable[0].times = "9:00-11:00"
			schedule.week[i].timetable[1].times = "11:00-13:00"
			schedule.week[i].timetable[2].times = "13:00-15:00"
			schedule.week[i].timetable[3].times = "15:00-17:00"
			schedule.week[i].timetable[4].times = "17:00-19:00"
		}
	}

	return schedule, nil
}

func (uu *userUsecase) UpdateSchedule(schedule *RequestUserUpdateSchedule) error {
	// TODO
	user := new(model.User)
	user.UserUID = schedule.uid

	for _, weekday := range schedule.week {
		f := ""
		for _, tt := range weekday.timetable {
			if tt.flag == false {
				f += "0"
			} else {
				f += "1"
			}
		}
		s := new(model.Schedule)
		s.Flags = f
		user.Schedule = append(user.Schedule, *s)
	}

	return uu.UserRepository.Update(user)
}

func (uu *userUsecase) ChangeName(uid string, newName string) (bool, error) {
	// TODO
	if foundSameName := uu.UserRepository.SearchSameName(newName); foundSameName == true {
		return true, nil
	}

	user := new(model.User)
	user.UserUID = uid
	user.Name = newName

	return false, uu.UserRepository.Update(user)
}

func (uu *userUsecase) Delete(uid string) error {
	return uu.UserRepository.Delete(uid)
}
