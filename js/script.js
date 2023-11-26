
document.getElementById('generate-list').addEventListener('click', function() {
    const recipeLink = document.getElementById('recipe-link').value;
    const servings = document.getElementById('serving-size').value;

    // ここにOpenAI APIを使用してレシピの解析と材料リストを生成するコードを記述
    // 例:
    // fetch('APIのURL', {
    //     method: 'POST',
    //     headers: {'Content-Type': 'application/json'},
    //     body: JSON.stringify({url: recipeLink, servings: servings})
    // })
    // .then(response => response.json())
    // .then(data => displayIngredients(data.ingredients))
    // .catch(error => console.error('Error:', error));

    // 以下は仮のレスポンスを表示するためのテストコードです。
    const testIngredients = ["材料1", "材料2", "材料3"]; // 仮の材料リスト
    displayIngredients(testIngredients);
});

function displayIngredients(ingredients) {
    const listDisplay = document.getElementById('list-display');
    listDisplay.innerHTML = ''; // リストをクリア
    ingredients.forEach(ingredient => {
        const listItem = document.createElement('li');
        listItem.textContent = ingredient;
        listDisplay.appendChild(listItem);
    });
}

document.getElementById('print-list').addEventListener('click', function() {
    window.print();
});

