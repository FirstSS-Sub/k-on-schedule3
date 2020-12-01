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
	Uid  string `json:"uid"`
	Week []day  `json:"week"`
}

type ResponseUserGetSchedule struct {
	Week []day `json:"week"`
}

type day struct {
	Date      string      `json:"date"`
	Timetable []timetable `json:"timetable"`
}

type timetable struct {
	Flag  bool   `json:"flag"`
	Times string `json:"times"`
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
				tt.Flag = false
				d.Timetable = append(d.Timetable, *tt)
			} else {
				tt.Flag = true
				d.Timetable = append(d.Timetable, *tt)
			}
		}
		schedule.Week = append(schedule.Week, *d)

		s := strings.Split(weekList[i], "-")
		// s[0]:2020, s[1]:11, s[2]:24
		schedule.Week[i].Date = s[1] + "/" + s[2] + " (" + weekdayList[i] + ")"

		if holidayFlagList[i] == 0 {
			schedule.Week[i].Timetable[0].Times = "9:00-10:30"
			schedule.Week[i].Timetable[1].Times = "10:30-12:00"
			schedule.Week[i].Timetable[2].Times = "12:00-14:00"
			schedule.Week[i].Timetable[3].Times = "14:00-16:00"
			schedule.Week[i].Timetable[4].Times = "16:00-17:30"
			schedule.Week[i].Timetable[5].Times = "17:30-19:00"
			schedule.Week[i].Timetable[6].Times = "19:00-20:30"
			schedule.Week[i].Timetable[7].Times = "20:30-22:00"
		} else {
			schedule.Week[i].Timetable[0].Times = "9:00-11:00"
			schedule.Week[i].Timetable[1].Times = "11:00-13:00"
			schedule.Week[i].Timetable[2].Times = "13:00-15:00"
			schedule.Week[i].Timetable[3].Times = "15:00-17:00"
			schedule.Week[i].Timetable[4].Times = "17:00-19:00"
		}
	}

	return schedule, nil
}

func (uu *userUsecase) UpdateSchedule(schedule *RequestUserUpdateSchedule) error {
	// TODO
	user := new(model.User)
	user.UserUID = schedule.Uid

	for _, weekday := range schedule.Week {
		f := ""
		for _, tt := range weekday.Timetable {
			if tt.Flag == false {
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
