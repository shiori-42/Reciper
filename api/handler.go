package api

import (
	"encoding/json"
	"net/http"
	"github.com/gocolly/colly/v2"
	"log"
	
)

// APIから受け取るリクエストの形式
type RecipeRequest struct {
	URL      string `json:"url"`
	Servings int    `json:"servings"`
}

// RecipeHandler はレシピの材料リストを取得し、スケーリングするAPIエンドポイント
func RecipeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req RecipeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 材料リストの計算
	ingredients, err := calculateIngredients(req.URL, req.Servings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// エンコードしてレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}

// calculateIngredients は与えられたURLとサービングに基づいて材料リストを計算します。
func calculateIngredients(url string, servings int) ([]string, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("cookpad.com", "www.cookpad.com"),
	)

	var ingredients []string
	var scrapeErr error // スクレイピング中に発生したエラーを保存する変数

	c.OnHTML("#ingredients_list", func(e *colly.HTMLElement) {
		e.ForEach("div.ingredient_row", func(_ int, el *colly.HTMLElement) {
			ingredient := el.Text
			// ここで材料をスケーリングするロジックを実装する場合
			// 例: scaledIngredient := scaleIngredient(ingredient, servings)
			// ingredients = append(ingredients, scaledIngredient)
			ingredients = append(ingredients, ingredient)
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error in scraping:", err)
		scrapeErr = err
	})

	// ウェブサイトを訪問
	c.Visit(url)

	if scrapeErr != nil {
		return nil, scrapeErr
	}

	// 実際のスケーリングロジックが必要です
	// 例えば、サービング数に応じて量を計算します
	// scaledIngredients := scaleIngredients(ingredients, servings)

	return ingredients, nil
}

// scaleIngredient は、与えられた材料の量を指定されたサービング数に応じてスケーリングします。
// 実際には、材料の量を解析し、新しいサービング数に基づいて量を計算するロジックが必要です。
func scaleIngredient(ingredient string, servings int) string {
	// 材料の量を解析し、サービング数に基づいてスケールするロジックを実装する
	return ingredient
}

// 例: scaleIngredients は、与えられた材料リストの量を指定されたサービング数に応じてスケーリングします。
func scaleIngredients(ingredients []string, servings int) []string {
	// 材料リストの各材料に scaleIngredient 関数を適用する
	for i, ingredient := range ingredients {
		 ingredients[i] = scaleIngredient(ingredient, servings)
	}
	return ingredients
}
