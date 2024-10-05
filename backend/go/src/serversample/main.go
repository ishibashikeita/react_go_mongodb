package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBに接続してデータをフェッチし、ログ出力する関数
func fetchAndLogFromMongoDB() {
	// 環境変数からmongoの接続情報を取得
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		log.Fatal("MONGO_URIが設定されていません！！！")
	}

	// 接続オプションを設定
	clientOptions := options.Client().ApplyURI(mongoUri)

	// MongoDBに接続
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// データベースとコレクションを指定
	db := client.Database("mydb")
	collection := db.Collection("users")

	// コレクションからデータを取得
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// 取得したデータをログに出力
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		log.Println(result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}

// HTTPのルートハンドラ
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("ルートにアクセスされました")
	fetchAndLogFromMongoDB()
	fmt.Fprintf(w, "ログ出力完了")
}

func main() {
	// ルートハンドラを設定
	http.HandleFunc("/", handler)

	// HTTPサーバーをポート8080で起動
	log.Println("サーバーがポート9090で起動しました...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
