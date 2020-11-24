package usecase

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/calendar"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/model"
	"github.com/FirstSS-Sub/k-on-schedule3/server/domain/repository"
	"strings"
)

type GroupUsecase interface {
	Create(name string) error
	FindById(id uint) (*ResponseGroupFindById, error)
	ChangeName(id uint, newName string) (bool, error)
	AddUser(groupId uint, userUid string) error
	Leave(groupId uint, userUid string) error
	Delete(id uint) error
}

type groupUsecase struct {
	GroupRepository repository.GroupRepository
}

type ResponseGroupFindById struct {
	name    string       `json:"name"`
	weekday []weekday    `json:"weekday"`
	users   []*model.User `json:"users"`
}

type weekday struct {
	date  string   `json:"date"`
	times []string `json:"times"`
}

// コンストラクタ
func NewGroupUsecase(groupRepository repository.GroupRepository) GroupUsecase {
	return &groupUsecase{GroupRepository: groupRepository}
}

func (gu *groupUsecase) Create(name string) error {
	return gu.GroupRepository.Insert(name)
}

func (gu *groupUsecase) FindById(id uint) (*ResponseGroupFindById, error) {
	group, err := gu.GroupRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	res := new(ResponseGroupFindById)
	res.name = group.Name
	res.users = group.Users

	weekList, holidayFlagList := calendar.GetCalendar()
	weekdayList := []string{"木", "金", "土", "日", "月", "火", "水"}

	for i := 0; i < 7; i++ {
		wd := new(weekday)

		s := strings.Split(weekList[i], "-")
		// s[0]:2020, s[1]:11, s[2]:24
		wd.date = s[1] + "/" + s[2] + " (" + weekdayList[i] + ")"

		var times []string
		if holidayFlagList[i] == 0 {
			times = append(times, "9:00-10:30")
			times = append(times, "10:30-12:00")
			times = append(times, "12:00-14:00")
			times = append(times, "14:00-16:00")
			times = append(times, "16:00-17:30")
			times = append(times, "17:30-19:00")
			times = append(times, "19:00-20:30")
			times = append(times, "20:30-22:00")
		} else {
			times = append(times, "9:00-11:00")
			times = append(times, "11:00-13:00")
			times = append(times, "13:00-15:00")
			times = append(times, "15:00-17:00")
			times = append(times, "17:00-19:00")
		}
		wd.times = times
		res.weekday = append(res.weekday, *wd)
	}

	return res, nil
}

func (gu *groupUsecase) ChangeName(id uint, newName string) (bool, error) {
	// TODO
	if foundSameName := gu.GroupRepository.SearchSameName(newName); foundSameName == true {
		return true, nil
	}

	group := new(model.Group)
	group.ID = id
	group.Name = newName

	return false, gu.GroupRepository.Update(group)
}

func (gu *groupUsecase) AddUser(groupId uint, userUid string) error {
	group := new(model.Group)
	user := new(model.User)
	group.ID = groupId
	user.UserUID = userUid

	return gu.GroupRepository.AddAssociation(group, user)
}

func (gu *groupUsecase) Leave(groupId uint, userUid string) error {
	group := new(model.Group)
	user := new(model.User)
	group.ID = groupId
	user.UserUID = userUid

	return gu.GroupRepository.DeleteAssociation(group, user)
}

func (gu *groupUsecase) Delete(id uint) error {
	return gu.GroupRepository.Delete(id)
}
