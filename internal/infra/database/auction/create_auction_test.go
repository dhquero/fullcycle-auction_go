package auction_test

import (
	"context"
	"os"
	"testing"
	"time"

	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, doc)
	return nil, args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return nil, args.Error(1)
}

func (m *MockCollection) Find(context.Context, interface{}, ...*options.FindOptions) (*mongo.Cursor, error) {
	return nil, nil
}

func (m *MockCollection) FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult {
	return nil
}

func TestCreateAuction_GoroutineUpdatesStatus(t *testing.T) {
	os.Setenv("AUCTION_INTERVAL", "1s")

	mockCollection := new(MockCollection)
	repo := &auction.AuctionRepository{Collection: mockCollection}

	ctx := context.Background()
	auctionEntity := &auction_entity.Auction{
		Id:          "myId",
		ProductName: "Mint candy",
		Category:    "Sweet",
		Description: "The best candy in the world",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	mockCollection.On("InsertOne", ctx, mock.Anything).Return(nil, nil)

	mockCollection.On("UpdateOne",
		ctx,
		bson.M{"_id": auctionEntity.Id},
		bson.M{"$set": bson.M{"status": auction_entity.Completed}},
	).Return(nil, nil).Once()

	err := repo.CreateAuction(ctx, auctionEntity)

	assert.Nil(t, err)

	time.Sleep(2 * time.Second)

	mockCollection.AssertCalled(t, "InsertOne", ctx, mock.Anything)
	mockCollection.AssertCalled(t, "UpdateOne",
		ctx,
		bson.M{"_id": auctionEntity.Id},
		bson.M{"$set": bson.M{"status": auction_entity.Completed}},
	)
	mockCollection.AssertExpectations(t)
}
