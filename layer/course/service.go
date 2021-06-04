package course

import (
	"errors"
	"fmt"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"time"
)

type Service interface {
	GetAllCourse() ([]entity.Courses, error)
	SaveNewCourse(course entity.CoursesInput) (entity.Courses, error)
	UpdateCourseByID(ID string, dataInput entity.UpdateCoursesInput) (entity.Courses, error)
	DeleteCourseByID(ID string) (interface{}, error)
	FilterCoursesByCategory(ID string) ([]entity.Courses, error)
	FilterCoursesByMastery(ID string) ([]entity.Courses, error)
	FilterCoursesBySubcription(ID string) ([]entity.Courses, error)
	FilterCoursesByCategoryAndMasteryAndSubcription(CID string, MID string, SID string) ([]entity.Courses, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllCourse() ([]entity.Courses, error) {
	courses, err := s.repository.GetAll()

	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (s *service) SaveNewCourse(course entity.CoursesInput) (entity.Courses, error) {

	var newCourse = entity.Courses{
		Name:        course.Name,
		Description: course.Description,
		VideoCourse: course.VideoCourse,
		CategoryID:  course.CategoryID,
		MasteryID:   course.MasteryID,
		SubID:       course.SubID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdCourse, err := s.repository.Create(newCourse)

	if err != nil {
		return createdCourse, err
	}

	return createdCourse, nil
}

func (s *service) UpdateCourseByID(ID string, dataInput entity.UpdateCoursesInput) (entity.Courses, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return entity.Courses{}, err
	}

	course, err := s.repository.FindByID(ID)

	if err != nil {
		return entity.Courses{}, err
	}

	if course.ID == 0 {
		newError := fmt.Sprintf("course id %s not found", ID)
		return entity.Courses{}, errors.New(newError)
	}

	if dataInput.Name != "" || len(dataInput.Name) != 0 {
		dataUpdate["name"] = dataInput.Name
	}
	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}
	if dataInput.VideoCourse != "" || len(dataInput.VideoCourse) != 0 {
		dataUpdate["video_course"] = dataInput.VideoCourse
	}
	if dataInput.CategoryID > 0 {
		dataUpdate["category_id"] = dataInput.CategoryID
	}
	if dataInput.MasteryID > 0 {
		dataUpdate["mastery_id"] = dataInput.MasteryID
	}
	if dataInput.SubID > 0 {
		dataUpdate["sub_type"] = dataInput.SubID
	}

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	courseUpdated, err := s.repository.UpdateByID(ID, dataUpdate)

	if err != nil {
		return courseUpdated, err
	}

	return courseUpdated, nil
}

func (s *service) DeleteCourseByID(ID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", ID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(ID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	if err != nil {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", ID)

	formatDelete := msg

	return formatDelete, nil

}

func (s *service) FilterCoursesByCategory(ID string) ([]entity.Courses, error) {
	courses, err := s.repository.GetByCategoryID(ID)

	if err != nil {
		return courses, err
	}

	if len(courses) == 0 {
		newError := fmt.Sprintf("course with category id %s not found", ID)
		return courses, errors.New(newError)
	}

	return courses, nil
}

func (s *service) FilterCoursesByMastery(ID string) ([]entity.Courses, error) {
	courses, err := s.repository.GetByMasteryID(ID)

	if err != nil {
		return courses, err
	}

	if len(courses) == 0 {
		newError := fmt.Sprintf("course with mastery id %s not found", ID)
		return courses, errors.New(newError)
	}

	return courses, nil
}

func (s *service) FilterCoursesBySubcription(ID string) ([]entity.Courses, error) {
	courses, err := s.repository.GetBySubcription(ID)

	if err != nil {
		return courses, err
	}

	if len(courses) == 0 {
		newError := fmt.Sprintf("course with subcription id %s not found", ID)
		return courses, errors.New(newError)
	}

	return courses, nil
}

func (s *service) FilterCoursesByCategoryAndMasteryAndSubcription(CID string, MID string, SID string) ([]entity.Courses, error) {
	courses, err := s.repository.GetByCategoryAndMasteryAndSubcription(CID, MID, SID)

	if err != nil {
		return courses, err
	}

	if len(courses) == 0 {
		newError := fmt.Sprintf("course with category id %s and mastery id %s and subcription id %s not found", CID, MID, SID)
		return courses, errors.New(newError)
	}

	return courses, nil
}
