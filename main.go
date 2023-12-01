// package main

package main

import(
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/shiori-42/^^^^^^"
)

//材料を格納する構造体
type Ingredient struct{
	Name string `json:"name"`
}

//レシピ全体の情報を格納する構造体
type Recipe struct{
	Title string `json:"title"`
	Servings int `json:"servings"`
	Ingredients []Ingredient `json:"ingredients"`
}

func main(){
	//材料リストのスライスを初期化
	var ingredients []Ingredient
	//CollyのCollectorを初期化
	c:=colly.NewCollector(
	colly.AllowedDomains("cookpad.com", "www.cookpad.com"),
	)

	//有効なUser-Agentヘッダーを設定
	c.UserAgent="Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

	// レシピタイトルとサービングサイズの取得は、実際のHTMLに合わせて適切なセレクタを設定
	c.OnHTML("適切なタイトルのセレクタ", func(e *colly.HTMLElement) {
		// レシピタイトルを取得して Recipe 構造体に格納
	})	

	c.OnHTML("適切なサービングサイズのセレクタ", func(e *colly.HTMLElement) {
		// サービングサイズを取得して Recipe 構造体に格納
	})

	//材料リストの要素を取得
	c.OnHTML("#ingredients_list",func(e *colly.HTMLElement){
		e.ForEach("div.ingredient_row",func(_ int,el*colly.HTMLElement){
			ingredient:=Ingredient{
				Name:el.Text,
			}
			ingredients=append(ingredients,ingredient)
		})
	})
	
	//ウェブサイトを訪問
	c.Visit("https://cookpad.com/recipe/759401")
	
	//材料リストをJSON形式で出力
	jsonFile,jsonErr:=os.Create("ingredients.json")
	if jsonErr!=nil{
		log.Fatalln("Failed to create the output JSONfile:",jsonErr)
	}
	defer jsonFile.Close()
	
	// Recipe構造体を作成してJSONに変換しファイルに保存
	recipe:=Recipe{
		Title: "レシピ名",
		Servings:4,
		Ingredients: ingredients,

	}

	jsonString,_:=json.MarshalIndent(recipe,"","  ")

	jsonFile.Write(jsonString)

	//HTTPサーバーを設定
	http.HandleFunc("/api./recipe",api.RecipeHandler)
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080",nil)

}

//https://cookpad.com/recipe/1751814




c.OnHTML(,func(e *colly.HTMLElement){

})

c.ONHTML(,func(e *colly.HTMLElement){

})


func parseServings(text string)int{

}

func scaleIngredients(ingredients []Ingredient,serbving int,targeSerbings int)[]Ingredient{

}

