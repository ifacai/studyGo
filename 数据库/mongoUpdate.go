package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	var update bson.D
	fmt.Println(update)
	update = append(update, bson.E{Key: "ccK", Value: "ccV"})
	fmt.Println(update)
	update = append(update, bson.E{Key: "ddK", Value: "ddV"})
	fmt.Println(update)

}

func updateAny(id string, updateParam bson.D) {
	ID, _ := primitive.ObjectIDFromHex(id)
	//update := bson.D{{"$set", bson.D{
	//	{"picUrl", picLocalUrl},
	//}}}
	_, err := MongoDBClient.Collection("news").UpdateOne(context.TODO(), bson.D{{"_id", ID}}, updateParam)
	if err != nil {
		fmt.Println(msg("更新 失败", "rbg"))
		panic(err)
	}
	fmt.Println(msg("更新 成功", "g"))
}
