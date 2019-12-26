package exception

type notFoundException struct {

}

func (notFoundException) Error() string  {
	return "Not Found"
}

func NotFoundException() error  {
	return notFoundException{}
}
