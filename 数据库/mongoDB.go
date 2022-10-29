package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// ConnectToDB pool 连接池模式
func mongoDB() *mongo.Database {
	//user := ""
	//password := ""
	host := "127.0.0.1"
	port := "27017"
	dbName := "news"
	timeOut := 3
	maxNum := 50
	//minNum := 5
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?w=majority", user, password, host, port, dbName)
	uri := fmt.Sprintf("mongodb://%s:%s/%s", host, port, dbName)
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut))
	defer cancel()
	// 抄的
	// 通过传进来的uri连接相关的配置
	opt := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	opt.SetMaxPoolSize(uint64(maxNum))
	//最小连接数
	//opt.SetMinPoolSize(uint64(minNum))
	//want, err := readpref.New(readpref.SecondaryMode) //表示只使用辅助节点
	//if err != nil {
	//	fmt.Println(err)
	//}
	//second
	//wc := writeconcern.New(writeconcern.WMajority())
	//readconcern.Majority()
	//opt := options.Client().ApplyURI(uri)
	//opt.SetLocalThreshold(3 * time.Second)     //只使用与mongo操作耗时小于3秒的
	opt.SetMaxConnIdleTime(5 * time.Second) //指定连接可以保持空闲的最大毫秒数
	//opt.SetMaxPoolSize(200)                    //使用最大的连接数
	//opt.SetReadPreference(want) //表示只使用辅助节点
	//opt.SetReadConcern(readconcern.Majority()) //指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
	//opt.SetWriteConcern(wc)                    //请求确认写操作传播到大多数mongod实例
	//if client, err = mongo.Connect(getContext(), opt); err != nil {
	//	fmt.Println(err)
	//}
	// 发起链接
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println("ConnectToDB", err)
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		fmt.Println("ConnectToDB", err)
	}
	// 返回 client.Collection
	return client.Database(dbName)
}

func cateUrl(parent, name string) string {
	type urlResultStruct struct {
		Url string
	}
	var (
		urlResult      urlResultStruct
		urlExistResult urlResultStruct
		cateUrl        string
		addUrl         string
	)
	filter := bson.M{
		"parent": parent,
		"name":   name,
	}

	err := mongoDB().Collection("cate").FindOne(context.TODO(), filter).Decode(&urlResult)
	if err != nil {
		//没有数据 新增
		firstPinyin := pinyin(name)
		firstPinyin = unABCNumber(firstPinyin)
		//check Url exist
		checkUrlExist := mongoDB().Collection("cate").FindOne(context.TODO(), bson.M{"url": firstPinyin}).Decode(&urlExistResult)
		if checkUrlExist != nil {
			//没有URL 直接用
			addUrl = firstPinyin
		} else {
			topPinyin := pinyin(parent)
			addUrl = topPinyin + firstPinyin
		}
		_, err := mongoDB().Collection("cate").InsertOne(context.TODO(), bson.D{
			{"parent", parent},
			{"name", name},
			{"url", addUrl},
		})
		if err != nil {
			panic(err)
		}
		//fmt.Println("id:", insertCateUrl.InsertedID)
		cateUrl = addUrl
	} else {
		//存在数据 直接返回
		cateUrl = urlResult.Url
	}
	return cateUrl
}

type source struct {
	Source      string
	SourceId    string
	Parent      string
	Logo        string
	Description string
	Url         string
}

func sourceUrl(sourceName, sourceId, sourceParent, sourceLogo, sourceDescription string) string {
	var (
		result         source
		urlExistResult source
		addUrl         string
		sourceUrl      string
	)
	err := mongoDB().Collection("source").FindOne(context.TODO(), bson.M{"source": sourceName}).Decode(&result)
	if err != nil {
		//没有数据 新增
		fmt.Println(msg("no data", "gbg"))
		firstPinyin := pinyin(sourceName)
		firstPinyin = unABCNumber(firstPinyin)
		//check Url exist
		checkUrlExist := mongoDB().Collection("source").FindOne(context.TODO(), bson.M{"url": firstPinyin}).Decode(&urlExistResult)
		if checkUrlExist != nil {
			//没有URL 直接用
			addUrl = firstPinyin
		} else {
			addUrl = firstPinyin + sourceId
		}
		insertCateUrl, err := mongoDB().Collection("source").InsertOne(context.TODO(), bson.M{
			"status":      "1",
			"source":      sourceName,
			"sourceId":    sourceId,
			"parent":      sourceParent,
			"logo":        sourceLogo,
			"description": sourceDescription,
			"url":         addUrl,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", insertCateUrl.InsertedID)
		sourceUrl = addUrl
	} else {
		//存在数据 直接返回
		sourceUrl = result.Url
	}
	return sourceUrl
}
