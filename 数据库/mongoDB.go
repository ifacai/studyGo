package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
	"time"
)

var MongoDBClient = mongoDB()

// ConnectToDB pool 连接池模式
func mongoDB() *mongo.Database {
	//user := ""
	//password := ""
	host := "127.0.0.1"
	port := "27017"
	dbName := "news"
	timeOut := 5
	maxNum := 50
	minNum := 15
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?w=majority", user, password, host, port, dbName)
	uri := fmt.Sprintf("mongodb://%s:%s/%s", host, port, dbName)
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut))
	defer cancel()
	// 通过传进来的uri连接相关的配置
	opt := options.Client().ApplyURI(uri)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	opt.SetMaxPoolSize(uint64(maxNum))
	//最小连接数
	opt.SetMinPoolSize(uint64(minNum))
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

//	type indexSortListStruct struct {
//		Title string
//		Time  string
//		Id    string `bson:"_id"`
//	}
//
//	func indexSortList(cateUrl string) []indexSortListStruct {
//		var (
//			//list    = make([]map[string]string, 10)
//			results []indexSortListStruct
//		)
//		findOptions := options.Find()
//		findOptions.SetSort(map[string]int{"time": -1})
//		findOptions.SetSkip(0)   // skip whatever you want, like `offset` clause in mysql
//		findOptions.SetLimit(10) // like `limit` clause in mysql
//		cur, err := MongoDBClient.Collection("news").Find(context.TODO(), bson.D{{"cateUrl", cateUrl}}, findOptions)
//		if err != nil {
//			panic(err)
//		}
//		if err = cur.All(context.TODO(), &results); err != nil {
//			panic(err)
//		}
//		return results
//		//for k, v := range results {
//		//	list[k] = make(map[string]string, 10)
//		//	list[k]["title"] = v.Title
//		//	list[k]["time"] = v.Time
//		//}
//		//return list
//	}

func indexSortList(cateUrl string) []map[string]string {
	type indexSortListStruct struct {
		Title string
		Time  string
		Id    string `bson:"_id"`
	}
	var (
		list    = make([]map[string]string, 20)
		results []indexSortListStruct
	)
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"time": -1})
	findOptions.SetSkip(0)   // skip whatever you want, like `offset` clause in mysql
	findOptions.SetLimit(20) // like `limit` clause in mysql
	cur, err := MongoDBClient.Collection("news").Find(context.TODO(), bson.D{{"cateUrl", cateUrl}}, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for k, v := range results {
		list[k] = make(map[string]string, 6)
		newsTime, err := strconv.Atoi(v.Time)
		if err != nil {
			newsTime = 1666666666
		}
		list[k]["y"] = time.Unix(int64(newsTime), 0).Format("2006")
		list[k]["m"] = time.Unix(int64(newsTime), 0).Format("01")
		list[k]["d"] = time.Unix(int64(newsTime), 0).Format("02")
		list[k]["time"] = time.Unix(int64(newsTime), 0).Format("2006-01-02 15:04:05")
		list[k]["title"] = v.Title
		list[k]["id"] = v.Id
	}
	return list
}

func findInfoById(id string) map[string]string {
	type news struct {
		Title       string
		Cover       string
		Time        string
		CateName    string
		CateUrl     string
		SubCateName string
		SubCateUrl  string
		Tags        string
		Source      string
		SourceUrl   string
		Description string
		Content     string
	}
	var (
		r    news
		data = make(map[string]string, 12)
	)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	err = MongoDBClient.Collection("news").FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&r)
	if err != nil {
		panic(err)
	}
	data["title"] = r.Title
	data["cover"] = r.Cover
	data["time"] = r.Time
	data["cateName"] = r.CateName
	data["cateUrl"] = r.CateUrl
	data["subCateName"] = r.SubCateName
	data["subCateUrl"] = r.SubCateUrl
	data["tags"] = r.Tags
	data["source"] = r.Source
	data["sourceUrl"] = r.SourceUrl
	data["description"] = r.Description
	data["content"] = r.Content
	return data
}

func allCate() (map[string]string, map[string]map[string]map[string]map[string]string) {
	type cateStruct struct {
		Parent string
		Name   string
		Url    string
	}
	var (
		res    []cateStruct
		parent = make(map[string]string)
		child  = make(map[string]map[string]map[string]map[string]string)
	)
	cur, err := MongoDBClient.Collection("cate").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		if v.Parent == "top" {
			parent[v.Url] = v.Name
		} else {
			if child[v.Parent] == nil {
				child[v.Parent] = make(map[string]map[string]map[string]string)
			}
			if child[v.Parent]["list"] == nil {
				child[v.Parent]["list"] = make(map[string]map[string]string)
			}
			if child[v.Parent]["list"][v.Url] == nil {
				child[v.Parent]["list"][v.Url] = make(map[string]string, 2)
			}
			child[v.Parent]["list"][v.Url]["name"] = v.Name
			child[v.Parent]["list"][v.Url]["url"] = v.Url
		}
	}
	return parent, child
}

func cateName(cateParent, cateUrl string) string {
	type cate struct {
		Name string
	}
	var res cate
	err := MongoDBClient.Collection("cate").FindOne(context.TODO(), bson.D{
		{"parent", cateParent},
		{"url", cateUrl},
	}).Decode(&res)
	if err != nil {
		panic(err)
	}
	return res.Name
}

func cateListNews(cateType, cateUrl string) [50]map[string]string {
	type cateListNewsStruct struct {
		Title string
		Cover string
		Time  string
		//CateName    string
		CateUrl     string
		SubCateName string
		SubCateUrl  string
		//Tags        string
		//Source      string
		//SourceUrl   string
		Description string
		Id          string `bson:"_id"`
	}
	var (
		res  []cateListNewsStruct
		list [50]map[string]string
	)
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"time": -1})
	findOptions.SetSkip(0)   // skip whatever you want, like `offset` clause in mysql
	findOptions.SetLimit(50) // like `limit` clause in mysql
	cur, err := MongoDBClient.Collection("news").Find(context.TODO(), bson.D{{cateType, cateUrl}}, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		list[k] = make(map[string]string, 11)
		newsTime, err := strconv.Atoi(v.Time)
		if err != nil {
			newsTime = 1666666666
		}
		list[k]["title"] = v.Title
		list[k]["cover"] = v.Cover
		list[k]["y"] = time.Unix(int64(newsTime), 0).Format("2006")
		list[k]["m"] = time.Unix(int64(newsTime), 0).Format("01")
		list[k]["d"] = time.Unix(int64(newsTime), 0).Format("02")
		list[k]["time"] = time.Unix(int64(newsTime), 0).Format("2006-01-02 15:04:05")
		//list[k]["cateName"] = v.CateName
		list[k]["cateUrl"] = v.CateUrl
		list[k]["subCateName"] = v.SubCateName
		list[k]["subCateUrl"] = v.SubCateUrl
		//list[k]["tags"] = v.Tags
		//list[k]["source"] = v.Source
		//list[k]["sourceUrl"] = v.SourceUrl
		list[k]["description"] = v.Description
		list[k]["id"] = v.Id
	}
	return list
}

func newsIndexList() [100]map[string]string {
	type cateListNewsStruct struct {
		Title       string
		Cover       string
		Time        string
		CateUrl     string
		SubCateName string
		SubCateUrl  string
		Description string
		Id          string `bson:"_id"`
	}
	var (
		res  []cateListNewsStruct
		list [100]map[string]string
	)
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"time": -1})
	findOptions.SetSkip(0)
	findOptions.SetLimit(100)
	cur, err := MongoDBClient.Collection("news").Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		list[k] = make(map[string]string, 11)
		newsTime, err := strconv.Atoi(v.Time)
		if err != nil {
			newsTime = 1666666666
		}
		list[k]["title"] = v.Title
		list[k]["cover"] = v.Cover
		list[k]["y"] = time.Unix(int64(newsTime), 0).Format("2006")
		list[k]["m"] = time.Unix(int64(newsTime), 0).Format("01")
		list[k]["d"] = time.Unix(int64(newsTime), 0).Format("02")
		list[k]["time"] = time.Unix(int64(newsTime), 0).Format("2006-01-02 15:04:05")
		list[k]["cateUrl"] = v.CateUrl
		list[k]["subCateName"] = v.SubCateName
		list[k]["subCateUrl"] = v.SubCateUrl
		list[k]["description"] = v.Description
		list[k]["id"] = v.Id
	}
	return list
}

func newsInfoList(cateType, cateUrl string) [23]map[string]string {
	type cateListNewsStruct struct {
		Title string
		Cover string
		Time  string
		Id    string `bson:"_id"`
	}
	var (
		res  []cateListNewsStruct
		list [23]map[string]string
	)
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"time": -1})
	findOptions.SetSkip(0)   // skip whatever you want, like `offset` clause in mysql
	findOptions.SetLimit(23) // like `limit` clause in mysql
	cur, err := MongoDBClient.Collection("news").Find(context.TODO(), bson.D{{cateType, cateUrl}}, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		list[k] = make(map[string]string, 7)
		newsTime, err := strconv.Atoi(v.Time)
		if err != nil {
			newsTime = 1666666666
		}
		list[k]["title"] = v.Title
		list[k]["cover"] = v.Cover
		list[k]["y"] = time.Unix(int64(newsTime), 0).Format("2006")
		list[k]["m"] = time.Unix(int64(newsTime), 0).Format("01")
		list[k]["d"] = time.Unix(int64(newsTime), 0).Format("02")
		list[k]["time"] = time.Unix(int64(newsTime), 0).Format("2006-01-02 15:04:05")
		list[k]["id"] = v.Id
	}
	return list
}

func tagListNews(tag string) [50]map[string]string {
	type cateListNewsStruct struct {
		Title       string
		Cover       string
		Time        string
		CateUrl     string
		SubCateName string
		SubCateUrl  string
		Description string
		Id          string `bson:"_id"`
	}
	var (
		res  []cateListNewsStruct
		list [50]map[string]string
	)
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"time": -1})
	findOptions.SetSkip(0)   // skip whatever you want, like `offset` clause in mysql
	findOptions.SetLimit(50) // like `limit` clause in mysql
	filter := bson.D{}
	filter = append(filter, bson.E{
		Key: "tags",
		//i 表示不区分大小写
		Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + tag + ".*", Options: "i"}},
	})
	cur, err := MongoDBClient.Collection("news").Find(context.TODO(), filter, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		list[k] = make(map[string]string, 11)
		newsTime, err := strconv.Atoi(v.Time)
		if err != nil {
			newsTime = 1666666666
		}
		list[k]["title"] = v.Title
		list[k]["cover"] = v.Cover
		list[k]["y"] = time.Unix(int64(newsTime), 0).Format("2006")
		list[k]["m"] = time.Unix(int64(newsTime), 0).Format("01")
		list[k]["d"] = time.Unix(int64(newsTime), 0).Format("02")
		list[k]["time"] = time.Unix(int64(newsTime), 0).Format("2006-01-02 15:04:05")
		list[k]["cateUrl"] = v.CateUrl
		list[k]["subCateName"] = v.SubCateName
		list[k]["subCateUrl"] = v.SubCateUrl
		list[k]["description"] = v.Description
		list[k]["id"] = v.Id
	}
	return list
}

func sourceInfo(sourceUrl string) map[string]string {
	type sourceStruct struct {
		Source      string
		Parent      string
		Logo        string
		Description string
	}
	var (
		r    sourceStruct
		data = make(map[string]string, 5)
	)
	err := MongoDBClient.Collection("source").FindOne(context.TODO(), bson.M{"url": sourceUrl}).Decode(&r)
	if err != nil {
		panic(err)
	}
	data["source"] = r.Source
	data["parent"] = r.Parent
	data["logo"] = r.Logo
	data["description"] = r.Description
	return data
}

func sourceNewsList(sourceUrl string) [50]map[string]string {
	type cateListNewsStruct struct {
		Title       string
		Cover       string
		Time        string
		CateUrl     string
		SubCateName string
		SubCateUrl  string
		Description string
		Id          string `bson:"_id"`
	}
	var (
		res  []cateListNewsStruct
		list [50]map[string]string
	)
	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"time": -1})
	findOptions.SetSkip(0)
	findOptions.SetLimit(50)
	cur, err := MongoDBClient.Collection("news").Find(context.TODO(), bson.D{{"sourceUrl", sourceUrl}}, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		list[k] = make(map[string]string, 11)
		newsTime, err := strconv.Atoi(v.Time)
		if err != nil {
			newsTime = 1666666666
		}
		list[k]["title"] = v.Title
		list[k]["cover"] = v.Cover
		list[k]["y"] = time.Unix(int64(newsTime), 0).Format("2006")
		list[k]["m"] = time.Unix(int64(newsTime), 0).Format("01")
		list[k]["d"] = time.Unix(int64(newsTime), 0).Format("02")
		list[k]["time"] = time.Unix(int64(newsTime), 0).Format("2006-01-02 15:04:05")
		list[k]["cateUrl"] = v.CateUrl
		list[k]["subCateName"] = v.SubCateName
		list[k]["subCateUrl"] = v.SubCateUrl
		list[k]["description"] = v.Description
		list[k]["id"] = v.Id
	}
	return list
}

func sitemapCount() int64 {
	findOptions := options.Find()
	count, err := MongoDBClient.Collection("news").CountDocuments(context.TODO(), findOptions)
	if err != nil {
		panic(err)
	}
	return count
}

func sitemapList(page string) [5000]map[string]string {
	type sitemapStruct struct {
		Time string
		Id   string `bson:"_id"`
	}
	var (
		setSkip int
		res     []sitemapStruct
		list    [5000]map[string]string
	)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		panic(err)
	}
	if pageInt > 0 {
		setSkip = (pageInt - 1) * 5000
	} else {
		setSkip = 0
	}
	findOptions := options.Find()
	findOptions.SetSkip(int64(setSkip))
	findOptions.SetLimit(5000)
	cur, err := MongoDBClient.Collection("news").Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		if v.Id == "" {
			break
		}
		list[k] = make(map[string]string, 5)
		newsTime, err := strconv.Atoi(v.Time)
		if err != nil {
			newsTime = 1666666666
		}
		list[k]["y"] = time.Unix(int64(newsTime), 0).Format("2006")
		list[k]["m"] = time.Unix(int64(newsTime), 0).Format("01")
		list[k]["d"] = time.Unix(int64(newsTime), 0).Format("02")
		list[k]["time"] = time.Unix(int64(newsTime), 0).Format("2006-01-02T15:04:05+00:00")
		list[k]["id"] = v.Id
	}
	return list
}

func sitemapCateList() map[int]map[string]string {
	type cateStruct struct {
		Parent string
		Url    string
	}
	var (
		res  []cateStruct
		list = make(map[int]map[string]string)
	)
	cur, err := MongoDBClient.Collection("cate").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cur.All(context.TODO(), &res); err != nil {
		panic(err)
	}
	for k, v := range res {
		list[k] = make(map[string]string, 2)
		list[k]["parent"] = v.Parent
		list[k]["url"] = v.Url
	}
	return list
}
