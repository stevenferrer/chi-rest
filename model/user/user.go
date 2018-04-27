package user

type User struct {
	ID       int64  `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// Storer shall be implemented by database stores
type Storer interface {
	List(filter ...Filter) ([]User, error)
	GetByID(id int64) (User, error)
	GetByEmail(email string) (User, error)

	// in create, update, and delete methods
	// should we return the user id?
	Create(User) error
	UpdateByID(id int64, u User) error
	UpdateByEmail(email string, u User) error
	Delete(User) error
}

// Filter is used for filtering results
type Filter func(*FilterConfig) error

type FilterConfig struct {
	//User // easy method, inherit all User fields
	ID int64
}

func IDFilter(id int64) Filter {
	return func(fc *FilterConfig) error {
		fc.ID = id
		return nil
	}
}
