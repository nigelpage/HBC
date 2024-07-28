package members

import (
	"fmt"
	"regexp"
	"time"
)

// Member types
type memberType byte

const (
	bowlingMember memberType = 'B'
	socialMember  memberType = 'S'
)

// Make the member a bowling member
func (m *Member) MakeBowlingMember() {
	m.Type = bowlingMember
}

// Make the member a social member
func (m *Member) MakeSocialMember() {
	m.Type = socialMember
}

// Test for bowling member
func (m Member) IsBowlingMember() bool {
	return m.Type == bowlingMember
}

// Test for social member
func (m Member) IsSocialMember() bool {
	return m.Type == socialMember
}

// Membership history record
type Membership struct {
	Joined time.Time `json:"joined"`
	Left   time.Time `json:"left"`
}

// Member record
type Member struct {
	Number      int          `json:"number"`
	Firstname   string       `json:"firstname"`
	Lastname    string       `json:"lastname"`
	IsFinancial bool         `json:"isfinancial"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	Type        memberType   `json:"membertype"`
	History     []Membership `json:"history"`
	IsDeceased  bool         `json:"isdeceased"`
}

// Create new bowling member
func NewBowlingMember(firstname, lastname string, email string, phone string) (*Member, error) {
	m, err := newMember(firstname, lastname, email, phone)
	if err != nil {
		return nil, err
	}
	m.MakeBowlingMember()
	return m, nil
}

// Create new social member
func NewSocialMember(firstname, lastname string, email string, phone string) (*Member, error) {
	m, err := newMember(firstname, lastname, email, phone)
	if err != nil {
		return nil, err
	}
	m.MakeSocialMember()
	return m, nil
}

/// -------------------------------------------------------
/// ** Internal only support functions
/// -------------------------------------------------------

// Create new member
func newMember(firstname, lastname string, email string, phone string) (*Member, error) {
	if !isEmailValid(email) {
		return nil, fmt.Errorf("invalid email address")
	}

	if !isPhoneValid(phone) {
		return nil, fmt.Errorf("invalid phone number")
	}

	m := Member{
		Number:      0,
		Firstname:   firstname,
		Lastname:    lastname,
		IsFinancial: false,
		Email:       email,
		Phone:       phone,
		Type:        bowlingMember,
		History:     []Membership{{Joined: time.Now()}},
		IsDeceased:  false,
	}

	return &m, nil
}

// Validate email address
func isEmailValid(email string) bool {
	p := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return p.MatchString(email)
}

// Validate phone number
func isPhoneValid(phone string) bool {
	p := regexp.MustCompile(`^(\+\d{2}[ \-]{0,1}){0,1}(((\({0,1}[ \-]{0,1})0{0,1}\){0,1}[2|3|7|8]{1}\){0,1}[ \-]*(\d{4}[ \-]{0,1}\d{4}))|(1[ \-]{0,1}(300|800|900|902)[ \-]{0,1}((\d{6})|(\d{3}[ \-]{0,1}\d{3})))|(13[ \-]{0,1}([\d \-]{5})|((\({0,1}[ \-]{0,1})0{0,1}\){0,1}4{1}[\d \-]{8,10})))$`)
	return p.MatchString(phone)
}
