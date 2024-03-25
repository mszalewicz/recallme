package server

type Config struct {
	Database Database
}

type Database struct {
	Url string
}
