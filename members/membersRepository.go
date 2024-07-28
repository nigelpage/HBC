package members

type MembersRepository interface {
	AddMember(*Member) error
	UpdateMember(Member) error
	FindMemberByName(lastname string) ([]Member, error)
	GetAllCurrentMembers() ([]Member, error)
	GetAllLapsedMembers() ([]Member, error)
	GetAllDeceasedMembers() ([]Member, error)
}