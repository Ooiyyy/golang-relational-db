package model

type ActivityLog struct {
	ID        int    `json:"id"`
	IP        string `json:"ip"`
	Aktivitas string `json:"aktivitas"`
	CreatedAt string `json:"created_at"`
}
