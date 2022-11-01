package repositories

type ResponseDto struct {
	Name   string    `bson:"name"`
	Number int64     `bson:"number"`
	Result [][]int64 `bson:"result"`
}
