package repository

import (
	"testing"
)

// Реализация интерфейса MemberRepository

type ArrayMemberRepository struct {
	members []Member
}

func (m *ArrayMemberRepository) GetAll() []Member {
	return m.members
}

func (m *ArrayMemberRepository) Save(mb Member) {
	m.members = append(m.members, mb)
}

func (m *ArrayMemberRepository) FindByEmail(email string) (Member, bool) {
	for _, mb := range m.members {
		if mb.Email == email {
			return mb, true
		}
	}
	return Member{}, false
}

// Пример использования ArrayMemberRepository в другом "классе"

type RegisterMember struct {
	Email, Password string
}

type RegisterMemberHandler struct {
	members MemberRepository
}

func NewRegisterMemberHandler(members MemberRepository) RegisterMemberHandler {
	return RegisterMemberHandler{members: members}
}

func (h *RegisterMemberHandler) Handle(command RegisterMember) {
	member := Member{Email: command.Email, Password: command.Password}
	h.members.Save(member)
}

func TestRepository(t *testing.T) {
	members := ArrayMemberRepository{}

	registerMembers := NewRegisterMemberHandler(&members)
	registerMembers.Handle(RegisterMember{Email: "email1", Password: "12345678"})
	registerMembers.Handle(RegisterMember{Email: "email2", Password: "1234"})
	registerMembers.Handle(RegisterMember{Email: "email3", Password: "123456"})

	if _, find := members.FindByEmail("email1"); !find {
		t.Error("email1 не найден")
	}

	if _, find := members.FindByEmail("email1"); !find {
		t.Error("email3 не найден")
	}

	if _, find := members.FindByEmail("email1"); !find {
		t.Error("email2 не найден")
	}

	if _, find := members.FindByEmail("email4"); find {
		t.Error("email4 найден")
	}
}
