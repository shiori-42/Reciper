
document.getElementById('generate-list').addEventListener('click', function() {
    const recipeLink = document.getElementById('recipe-link').value;
    const scale  = document.getElementById('scale-select').value;

// 仮の材料リスト
const testIngredients = ["ほうれん草 1/2束", "ベーコン 3枚", "たまご 2個"];	
displayIngredients(testIngredients, scale);
});

// 材料リストを表示する関数
// 材料リストを表示する関数
function displayIngredients(ingredients, scale) {
    const listDisplay = document.getElementById('list-display');
    listDisplay.innerHTML = ''; // リストをクリア
    ingredients.forEach(ingredient => {
        const listItem = document.createElement('li');

        // チェックボックスを作成 (変更: チェックボックスの追加)
        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.className = 'ingredient-checkbox'; // classを追加してCSSスタイルを適用
        checkbox.id = ingredient.replace(/[^a-zA-Z0-9]/g, "-"); // idにはスペースや特殊文字が含まれないようにする

        // ラベルを作成 (変更: チェックボックスと関連付けるためのラベルの追加)
        const label = document.createElement('label');
        label.htmlFor = checkbox.id; // チェックボックスと関連付け
        label.appendChild(document.createTextNode(ingredient));

        // リストアイテムにチェックボックスとラベルを追加 (変更: チェックボックスとラベルをリストアイテムに追加)
        listItem.appendChild(checkbox);
        listItem.appendChild(label);

        listDisplay.appendChild(listItem); // 最終的にリストにアイテムを追加
    });
}

document.getElementById('print-list').addEventListener('click', function() {
    window.print();
});

