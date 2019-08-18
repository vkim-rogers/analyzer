package connectors

var Sqs sqs

type sqs struct {

}

func InitSqs() *sqs{

	Sqs = sqs{}

	return &Sqs
}