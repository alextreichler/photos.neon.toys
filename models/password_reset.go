package models

type PasswordReset struct {
	ID     int
	UserID int
	Token string
	ExpiresAt time.time
}

const (
	DefaultResetDuration = 1 *time.Hour
)

type PasswordReset Service struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each password reset token. If the value is not set or is less than
	// the MinBytesPerToken const it will be ignored and use MinBytesPerToken
	BytesPerToken int
	// Duration is the amount of time that a PasswordReset is valid for.
	// Defaults to Default ResetDuration
	Duration time.Duration
}

func (service *PasswordResetService) Create(email string)(*PasswordReset, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Create")
}

func (service *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: Implement PasswordResetService.Consume")
}