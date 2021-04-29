package repository

type Member struct {
	Email, Password string
}

type MemberRepository interface {
	GetAll() []Member
	Save(m Member)
	FindByEmail(email string) (Member, bool)
}
