package global

var from string

func FROM() string {
	return from
}

func SetFrom(f string) {
	from = f
}
