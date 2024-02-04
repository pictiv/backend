package model

type Role string

const (
	ADMIN     Role = "ADMIN"
	MODERATOR Role = "MODERATOR"
	USER      Role = "USER"
)

var (
	Roles = map[Role]int{
		ADMIN:     0,
		MODERATOR: 1,
		USER:      2,
	}
)
