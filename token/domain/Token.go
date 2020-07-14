package domain

//Object base af all.
type Token struct {
	Token        string
	Exp          int64
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
