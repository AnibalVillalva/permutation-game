package repositories

import (
	"context"
	"errors"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"permutation-game/api/cmd/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 10 * time.Second
)

type Repo struct {
	user     string
	password string
	//Service      repositories.GetOperator
}

func New() *Repo {
	u := os.Getenv("USER")
	p := os.Getenv("PASSWORD")

	log.Debugf("[Repository] Read env.")

	return &Repo{user: u, password: p}
}

func (r Repo) Get(ctx context.Context) (context.Context, error) {
	var result2 ResponseDto

	log.Debugf("[Repository] Starting. %v", ctx)

	value := ctx.Value(entities.CtxKey)
	if value == nil {
		return ctx, errors.New(" Context is not initialized")
	}

	m, ok := value.(entities.Context)
	if !ok {
		return ctx, errors.New("invalid context")
	}

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + r.user + ":" + r.password + "@ensarnia.0l6wi.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("[Repository] Error: %v", err)
	}

	coll := client.Database("permutation").Collection("result")

	err = coll.FindOne(context.TODO(), bson.M{"number": m.Number()}).Decode(&result2)
	if err != nil {
		log.Fatalf("[Repository] Error: %v", err)
		panic(err)
	}

	log.Debugf("[Repository] result: name: %s, number:%d, result: %v", result2.Name, result2.Number, result2.Result)

	c := entities.Build().SetNumber(result2.Number).SetResult(result2.Result)
	rctx := context.WithValue(ctx, entities.CtxKey, c)
	log.Debug(rctx)

	return rctx, nil
}
