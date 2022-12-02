package alert

type Options struct {
	// Username is the username to use to authenticate to the SMTP server.
	username string
	// Password is the password to use to authenticate to the SMTP server.
	password string
}

type Option func(*Options)

func From(from string) Option {
	return func(opts *Options) {
		opts.username = from
	}
}

func Password(password string) Option {
	return func(opts *Options) {
		opts.password = password
	}
}
