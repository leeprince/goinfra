package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 14:14
 * @Desc:
 */

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const (
	// 注意：27017后面是斜杠后跟着问号！
	mongoUri = "mongodb://mongoadmin:TB5i9K2jD1SAasdr@10.21.32.14:27017/?connect=direct"

	Database   = "ptest-db"
	Collection = "ptest-col"
)

// 结构体必须定义 bson 标签
type dataStruct struct {
	Name         string     `bson:"name" json:"name"`
	Value        float64    `bson:"value" json:"value"`
	CreatedAt    time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `bson:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `bson:"deleted_at" json:"deleted_at"`
	CreatedAtInt int64      `bson:"created_at_int" json:"created_at_int"`
	UpdatedAtInt int64      `bson:"updated_at_int" json:"updated_at_int"`
	DeletedAtInt int64      `bson:"deleted_at_int" json:"deleted_at_int"`
}

var mongoClient *mongo.Client

// initMongoDBClient
func initMongoDBClient() {

	var err error

	// 连接  mongoDB 服务端
	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()*/
	ctx := context.Background()
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	// 使用全局变量，故不直接在 defer 中关闭连接
	/*defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()*/

	if err != nil {
		panic(err)
	}

	// 验证连接
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

}

func main() {
	// 初始化MongoDB客户端
	initMongoDBClient()

	InsertOne()
}

func InsertOne() {
	ctx := context.Background()
	collection := mongoClient.Database(Database).Collection(Collection)

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		panic(err)
	}

	fmt.Printf("InsertOne res:%+v\n", res)

	insertOne1 := dataStruct{
		Name:         "prince",
		Value:        1314.520,
		CreatedAt:    time.Now(),
		UpdatedAt:    nil,
		DeletedAt:    nil,
		CreatedAtInt: time.Now().Unix(),
		UpdatedAtInt: 0,
		DeletedAtInt: 0,
	}
	res, err = collection.InsertOne(ctx, insertOne1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("InsertOne res:%+v\n", res)

}

func Find() {
	ctx := context.Background()
	collection := mongoClient.Database(Database).Collection(Collection)

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....

		fmt.Printf("result: %+v\n", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func FindOne() {
	collection := mongoClient.Database(Database).Collection(Collection)

	var result struct {
		Value float64
	}
	filter := bson.D{{"name", "pi"}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	// Do something with result...

	fmt.Printf("result result: %+v\n", result)

	result1 := dataStruct{}
	// 注意：使用结构体当作过滤条件时，filter 中不会忽略零值,所以只有Name="pi"时不会查询到记录。
	// 解决：可以使用新的结构体当作过滤条件，如下方的`filterStruct := struct {...}`
	/*filterStruct := dataStruct{
		Name:  "pi",
		Value: 3.14159,
	}*/
	// 使用新的结构体当作过滤条件
	filterStruct := struct {
		Name string `bson:"name" json:"name"`
	}{
		Name: "pi",
	}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filterStruct).Decode(&result1)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	// Do something with result...

	fmt.Printf("result result1: %+v\n", result1)
}

func UpdateOne() {
	var coll *mongo.Collection
	coll = mongoClient.Database(Database).Collection(Collection)

	var id primitive.ObjectID
	id, _ = primitive.ObjectIDFromHex("6523a11ff3dd4f5270f5f5ec")

	// find the document for which the _id field matches id and set the email to "newemail@example.com"
	// specify the Upsert option to insert a new document if a document matching the filter isn't found
	//opts := options.Update().SetUpsert(true)

	opts := options.Update().SetUpsert(false)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"value", "1314.53"}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
}
