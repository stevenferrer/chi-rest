package user

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type Storer interface {
	List(filter ...Filter) ([]User, error)
	Get(id int64) (User, error)
	Create(User) error
	Update(User) error
	Delete(User) error
}

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
