package config

type Server struct {
	Data Data
	Blob Blob
}

type Data struct {
	Conn string
}

type Blob struct {
	Basepath string
}
