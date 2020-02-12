var createPane = document.getElementById("createArticlePane")
var showCreatePaneBtn = document.getElementById("showCreatePaneButton")

// 顯示id為'createArticlePane'的新建文章區塊(如果已經顯示，則隱藏。)
function showCreatePane() {
  showCreatePaneBtn.hidden = !showCreatePaneBtn.hidden;
  createPane.hidden = !createPane.hidden
}

// 顯示id為'edit<panID>'的編輯區塊(如果已經顯示，則隱藏。)
function showEditPane(paneID) {
  var pane = document.getElementById('edit'+paneID)
  var showEditButton = document.getElementById("showEditButton"+paneID)
  if (pane.hidden == true) {
      showEditButton.innerText = "取消編輯"
  } else {
      showEditButton.innerText = "編輯"
  }
  pane.hidden = !pane.hidden
}

// 直接發送一個刪除文章的POST請求，「刪除對應articleID的文章」
function deleteArticle(articleID) {
  var form = document.createElement('form');
  form.setAttribute('method', 'post');
  form.setAttribute('action', '/delete?articleID='+articleID);
  form.style.display = 'hidden';
  document.body.appendChild(form)
  form.submit();
}

// 顯示'刪除確認的dialog'，「避免」誤點擊刪除按鈕導致「直接刪除」。
function showDeleteDialog(articleID) {
  var dialog = document.getElementById('deleteDialog'+articleID)
  dialog.showModal();
  dialog.addEventListener('click', function (event) {
    var rect = dialog.getBoundingClientRect();
    var isInDialog=(rect.top <= event.clientY && event.clientY <= rect.top + rect.height
      && rect.left <= event.clientX && event.clientX <= rect.left + rect.width);
    if (!isInDialog) {
        dialog.close();
    }
});

}