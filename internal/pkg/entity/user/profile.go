package user

// Profile declare user information
type Profile struct {
	ID        int64  `json:"id,omitempty" db:"id"`
	Email     string `json:"email,omitempty" db:"email"`
	Password  string `json:"-" db:"password"`
	Token     string `json:"token,omitempty" db:"token"`
	Name      string `json:"name,omitempty" db:"name"`
	CreatedAt int64  `json:"created_at,omitempty" db:"created_at"`
}
