package main

import(
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
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
	
	//レシピタイトルを取得
	c.OnHTML("selecta",func(e *colly.HTMLElement){

	})
	
	//サービングサイズを取得
	c.OnHTML("class="servings"",func(e *colly.HTMLElement){

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
	
	//Recipe構造体を作成
	// recipe:=Recipe{
	// 	Title:="レシピ名"

	// }
	jsonString,_:=json.MarshalIndent(ingredients,"","")

	jsonFile.Write(jsonString)

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

