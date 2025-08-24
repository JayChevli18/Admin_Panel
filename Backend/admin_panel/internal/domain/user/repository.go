package user

import (
	context "context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context, page, pageSize int) ([]User, int64, error)
	EnsureIndexes(ctx context.Context) error
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
		{Keys: bson.D{{Key: "role", Value: 1}}},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexModel)
	return err
}

func (r *mongoRepository) Create(ctx context.Context, user *User) error {
	res, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}
	return nil
}

func (r *mongoRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	var user User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	return &user, err
}

func (r *mongoRepository) Update(ctx context.Context, user *User) (*User, error) {
	_, err := r.collection.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"role":      user.Role,
		"isActive":  user.IsActive,
	}})
	return user, err
}

func (r *mongoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
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
		SetSort(bson.D{{Key: "_id", Value: -1}}).
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

	count, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
