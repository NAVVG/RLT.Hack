package entity

type Achievement struct {
	Id         int
	Header     string
	Descripton string
	Icon       string
}

type AchievementProgress struct {
	AchievementId int
	UserId        int
}
