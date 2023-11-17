package errorstable

type ErrorCode struct{
	PAGE_NOT_FOUND  int 
}

var ErrorTable map[string]string = map[string]string{
	"0":"Success",
	"401":"Auth error",
	"404":"Not Found",
}