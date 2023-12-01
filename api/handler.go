//handler.go
package api

import(
	"net/http"
	"encoding/json"
)

//APIから受け取るリクエストの形式
type RecipeRequest struct{
	URL string `json:"url"`
	Servings int `json:"servings"`
}

func RecipeHandler(w http.ResponseWriter,r *http.Request)
	if r.Method!="POST"{
		http.Error(w,"Invalid request method",http.StatusMethodNotAllowed)
		return

	var req RecipeRequest
	err:=json.NewDecoder(r.Body).Decode(&req)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	//材料リストの計算
	ingredients:=calculateIngredients(req.URL,req.Servings)
	json.NewEncoder(w).Encode(ingredients)

}

	func caluclateIngredients(url string,servings int)[]strint{
		return[]string("a","b","cew")
}


	
	