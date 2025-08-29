package user

import (
	context "context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	EnsureIndexes(ctx context.Context) error
	Create(ctx context.Context, user *User) error
	GetByUserId(ctx context.Context, userId int64) (*User, error)
	UpdateByUserId(ctx context.Context, user *User) error
	DeleteByUserId(ctx context.Context, userId int64) error
	List(ctx context.Context, page, pageSize int) ([]User, int64, error)
}

type mongoRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &mongoRepository{collection: db.Collection(CollectionName())}
}

func (r *mongoRepository) EnsureIndexes(ctx context.Context) error {
	indexModel := []mongo.IndexModel{
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "userId", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "role", Value: 1}}},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexModel)
	return err
}

func (r *mongoRepository) Create(ctx context.Context, user *User) error {
	_, err := r.collection.InsertOne(ctx, user)
	// if err != nil {
	// 	return err
	// }
	// if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
	// 	user.ID = oid
	// }
	return err
}

func (r *mongoRepository) GetByUserId(ctx context.Context, userId int64) (*User, error) {
	var user User
	if err := r.collection.FindOne(ctx, bson.M{"userId": userId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mongoRepository) UpdateByUserId(ctx context.Context, user *User) error {
	_, err := r.collection.UpdateOne(ctx,
		bson.M{"userId": user.UserID},
		bson.M{"$set": bson.M{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
			"role":      user.Role,
			"isActive":  user.IsActive,
		}})
	return err
}

func (r *mongoRepository) DeleteByUserId(ctx context.Context, userId int64) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"userId": userId})
	return err
}

func (r *mongoRepository) List(ctx context.Context, page, pageSize int) ([]User, int64, error) {

	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	findOptions := options.Find().
		SetSort(bson.D{{Key: "userId", Value: -1}}).
		SetLimit(int64(pageSize)).
		SetSkip(int64((page - 1) * pageSize))

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
