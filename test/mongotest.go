package test

import (
	"GINOWEN/global"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 插入用户数据
func InsertUser(client *mongo.Client, database, collection string, user map[string]interface{}) (*mongo.InsertOneResult, error) {
	col := client.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 查询单个用户数据
func FindUser(client *mongo.Client, database, collection string, filter bson.M) (bson.M, error) {
	col := client.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result bson.M
	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 更新用户数据
func UpdateUser(client *mongo.Client, database, collection string, filter, update bson.M) (*mongo.UpdateResult, error) {
	col := client.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := col.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 删除用户数据
func DeleteUser(client *mongo.Client, database, collection string, filter bson.M) (*mongo.DeleteResult, error) {
	col := client.Database(database).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func MongoTest() {

	var err error
	mongoClient := global.OWEN_MONGO
	config := global.OWEN_CONFIG
	if err != nil {
		log.Fatalf("MongoDB 连接失败: %v", err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("MongoDB 断开连接失败: %v", err)
		}
	}()

	// 插入用户
	user := map[string]interface{}{"name": "Alice", "age": 30}
	insertResult, err := InsertUser(mongoClient, config.MongoDB.Database, "users", user)
	if err != nil {
		log.Fatalf("插入用户失败: %v", err)
	}
	fmt.Printf("用户插入成功: %v\n", insertResult.InsertedID)

	// 查询用户
	filter := bson.M{"name": "Alice"}
	foundUser, err := FindUser(mongoClient, config.MongoDB.Database, "users", filter)
	if err != nil {
		log.Fatalf("查询用户失败: %v", err)
	}
	fmt.Printf("查询到的用户: %v\n", foundUser)

	// 更新用户
	update := bson.M{"age": 31}
	updateResult, err := UpdateUser(mongoClient, config.MongoDB.Database, "users", filter, update)
	if err != nil {
		log.Fatalf("更新用户失败: %v", err)
	}
	fmt.Printf("更新用户成功, 修改的文档数量: %d\n", updateResult.ModifiedCount)

	// 删除用户
	deleteResult, err := DeleteUser(mongoClient, config.MongoDB.Database, "users", filter)
	if err != nil {
		log.Fatalf("删除用户失败: %v", err)
	}
	fmt.Printf("删除用户成功, 删除的文档数量: %d\n", deleteResult.DeletedCount)
}
