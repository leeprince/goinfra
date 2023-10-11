package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/9 14:14
 * @Desc:
 */

import (
	"context"
	"fmt"
	"github.com/leeprince/goinfra/utils/dumputil"
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
	CreatedAt    time.Time  `bson:"created_at" json:"created_at"` // mongoDB 存储time.Time是跟时区有关的。即存储是会转为默认时区，读取时转为当前时区
	UpdatedAt    *time.Time `bson:"updated_at" json:"updated_at"` // mongoDB 存储time.Time是跟时区有关的。即存储是会转为默认时区，读取时转为当前时区
	DeletedAt    *time.Time `bson:"deleted_at" json:"deleted_at"` // mongoDB 存储time.Time是跟时区有关的。即存储是会转为默认时区，读取时转为当前时区
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
	fmt.Println("res.InsertedID:", res.InsertedID)

	ctime := time.Now()
	ctimeU := ctime.Unix()

	name := "prince"
	value := 1314.520
	insertOne1 := dataStruct{
		Name:         name,
		Value:        value,
		CreatedAt:    ctime,
		UpdatedAt:    nil,
		DeletedAt:    nil,
		CreatedAtInt: ctimeU,
		UpdatedAtInt: 0,
		DeletedAtInt: 0,
	}
	res, err = collection.InsertOne(ctx, insertOne1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("InsertOne res:%+v\n", res)

}

// 批量查询
func Find() {
	ctx := context.Background()
	collection := mongoClient.Database(Database).Collection(Collection)

	/*查询到bson.D*/
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

		fmt.Printf("cur.Next result: %+v\n", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------------------")
	/*查询到结构体*/
	cur, err = collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result []dataStruct
	err = cur.All(ctx, &result)
	if err != nil {
		log.Fatal(err)
	}
	// Do something with result...

	dumputil.Println("result result Println: ", result)
}

// 批量查询
func FindFilter() {
	ctx := context.Background()
	collection := mongoClient.Database(Database).Collection(Collection)

	/*通过bson.D作为查询条件，并查询到结构体*/
	filter := bson.D{{"name", "pi"}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var result1 []dataStruct
	err = cur.All(ctx, &result1)
	if err != nil {
		log.Fatal(err)
	}
	// Do something with result...
	dumputil.Println("result result1 Println: ", result1)

	fmt.Println("-----------------------------------------")
	/*通过结构体作为查询条件，并查询到结构体*/
	// 注意：使用结构体当作过滤条件时，filter 中不会忽略零值和nil。下面是解决办法
	// 解决方法1：可以使用新的结构体当作过滤条件
	// 解决方法2：在结构体的字段 bson 标签中添加忽略空或零值
	/*filterStruct := dataStruct{
		Name:  "pi",
		Value: 3.14159,
	}*/
	// 解决方法1：可以使用新的结构体当作过滤条件
	/*filterStruct := struct {
		Name  string `bson:"name" json:"name"`
	}{
		Name: "pi",
	}*/
	// 解决方法2：在结构体的字段 bson 标签中添加忽略空或零值
	filterStruct := struct {
		Name  string `bson:"name" json:"name"`
		Value string `bson:"value,omitempty" json:"value"`
	}{
		Name: "pi",
	}
	cur, err = collection.Find(ctx, filterStruct)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result2 []dataStruct
	err = cur.All(ctx, &result2)
	if err != nil {
		log.Fatal(err)
	}
	// Do something with result...
	dumputil.Println("result result2 Println: ", result2)
}

// 批量查询:排序，升序=1；降序=-1
func FindAndSort() {
	ctx := context.Background()
	collection := mongoClient.Database(Database).Collection(Collection)

	/*查询到结构体*/
	opt := options.Find().SetSort(bson.D{{"created_at_int", -1}})
	cur, err := collection.Find(ctx, bson.D{}, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result1 []dataStruct
	err = cur.All(ctx, &result1)
	if err != nil {
		log.Fatal(err)
	}
	// Do something with result...

	dumputil.Println("result result created_at_in=-1", result1)

	fmt.Println("-----------------------------------------")

	/*查询到结构体*/
	opt = options.Find().SetSort(bson.D{{"created_at_int", 1}})
	cur, err = collection.Find(ctx, bson.D{}, opt)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result2 []dataStruct
	err = cur.All(ctx, &result2)
	if err != nil {
		log.Fatal(err)
	}
	// Do something with result...

	dumputil.Println("result result created_at_in=1", result2)
}

func Count() {
	var coll *mongo.Collection
	coll = mongoClient.Database(Database).Collection(Collection)

	filter := bson.D{{"name", "prince"}}
	count, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	dumputil.Println("count", count)
}

func Page() {
	var coll *mongo.Collection
	coll = mongoClient.Database(Database).Collection(Collection)

	filter := bson.D{}
	opts := options.Find().SetLimit(2).SetSkip(1)
	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		panic(err)
	}
	var result2 []dataStruct
	err = cursor.All(context.Background(), &result2)
	if err != nil {
		log.Fatal(err)
	}
	// Do something with result...

	dumputil.Println("result result created_at_in=1", result2)
}

func FindOne() {
	collection := mongoClient.Database(Database).Collection(Collection)

	/*通过bson.D作为查询条件，并查询到结构体*/
	var result dataStruct
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

	dumputil.Println("result result", result)

	fmt.Println("-----------------------------------------")

	/*通过结构体作为查询条件，并查询到结构体*/
	result1 := dataStruct{}
	filter = bson.D{{"name", "pi"}}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result1)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	// Do something with result...

	dumputil.Println("result result1", result1)
}

func UpdateOne() {
	var coll *mongo.Collection
	coll = mongoClient.Database(Database).Collection(Collection)

	var id primitive.ObjectID
	id, _ = primitive.ObjectIDFromHex("6523a11ff3dd4f5270f5f5ec")

	// 不存在时是否查询
	// find the document for which the _id field matches id and set the email to "newemail@example.com"
	// specify the Upsert option to insert a new document if a document matching the filter isn't found
	//opts := options.Update().SetUpsert(true)
	opts := options.Update().SetUpsert(false)

	// 查询条件和设置的值可以参考插入和查询时使用的结构体方式
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"value", "1314.53"}}}}

	result, err := coll.UpdateOne(context.Background(), filter, update, opts)
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
