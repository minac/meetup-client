// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/Guitarbum722/meetup-client/models"

// Clienter is an autogenerated mock type for the Clienter type
type Clienter struct {
	mock.Mock
}

// GroupByID provides a mock function with given fields: groupIDs
func (_m *Clienter) GroupByID(groupIDs ...int) (*models.Groups, error) {
	_va := make([]interface{}, len(groupIDs))
	for _i := range groupIDs {
		_va[_i] = groupIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Groups
	if rf, ok := ret.Get(0).(func(...int) *models.Groups); ok {
		r0 = rf(groupIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Groups)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...int) error); ok {
		r1 = rf(groupIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupByOrganizer provides a mock function with given fields: organizerIDs
func (_m *Clienter) GroupByOrganizer(organizerIDs ...int) (*models.Groups, error) {
	_va := make([]interface{}, len(organizerIDs))
	for _i := range organizerIDs {
		_va[_i] = organizerIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Groups
	if rf, ok := ret.Get(0).(func(...int) *models.Groups); ok {
		r0 = rf(organizerIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Groups)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...int) error); ok {
		r1 = rf(organizerIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupByURLName provides a mock function with given fields: urlNames
func (_m *Clienter) GroupByURLName(urlNames ...string) (*models.Groups, error) {
	_va := make([]interface{}, len(urlNames))
	for _i := range urlNames {
		_va[_i] = urlNames[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Groups
	if rf, ok := ret.Get(0).(func(...string) *models.Groups); ok {
		r0 = rf(urlNames...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Groups)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(urlNames...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Member provides a mock function with given fields: memberID
func (_m *Clienter) Member(memberID int) (*models.Member, error) {
	ret := _m.Called(memberID)

	var r0 *models.Member
	if rf, ok := ret.Get(0).(func(int) *models.Member); ok {
		r0 = rf(memberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(memberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Members provides a mock function with given fields: groupID
func (_m *Clienter) Members(groupID int) (*models.Members, error) {
	ret := _m.Called(groupID)

	var r0 *models.Members
	if rf, ok := ret.Get(0).(func(int) *models.Members); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Members)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
