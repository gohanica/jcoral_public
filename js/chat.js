


// 送信系HTML取得
let input = document.getElementById('input');
let username = document.getElementById('username');

// websocket通信開始
let socket = new WebSocket("ws://" + window.location.host + "/chat");

// websocketクライアントからサーバーへ送信
function send() {
    socket.send(JSON.stringify(
        {
            username: username.value,
            message: input.value,
        }
    ));
    input.value = ""
    username.value = "Otani"
};


let userdata = [];
// websocket サーバからメッセージ受信
socket.onmessage = function (e) {
    let now = new Date();
    let month = now.getMonth() + 1;
    let nowtime = now.getFullYear() + "/" + month + "/" + now.getDate() + "/" + now.getHours() + ":" + now.getMinutes();
    let comment = {
        profileimage: "../image/0.png",
        name: "大谷翔平",
        time: nowtime,
        id: 17123456789,
        text: "google最強！"
    }
    // 実験
    let comments = document.getElementsByClassName("comment");
    let js = JSON.parse(e.data);

    // テンプレートデータ取得整形
    userdata = { Message: js.message, Username: js.username }

    // 一つのコメント作成
    // ユーザープロフィール
    // div追加
    let userinfo = $("<div></div>");
    userinfo.addClass("row");
    userinfo.appendTo(".container");

    // col追加
    let usercol = $("<div></div>");
    usercol.addClass("col");
    usercol.appendTo(".row:last");

    // media追加
    let usercomment = $("<div></div>");
    usercomment.addClass("media");
    usercomment.appendTo(".col:last");

    // imgタグ追加・プロフィール画像追加
    let userprofile = $("<img>");
    userprofile.addClass("mr-3");
    userprofile.appendTo(".media:last");
    userprofile.attr('src', comment.profileimage);
    let profilestyle = {
        'width': '40',
        'height': '40',
    }
    userprofile.css(profilestyle);
    // ユーザー情報・コメント
    let usertext = $("<div></div>");
    usertext.addClass("media-body");
    usertext.appendTo(".media:last");
    let username = $("<h5></h5>");
    username.appendTo(".media-body:last");
    username.html(userdata.Username + "\t" + comment.time + "\tID:" + comment.id);
    let userchat = $("<h4></h4>");
    userchat.appendTo(".media-body:last");
    userchat.text(userdata.Message);

}
// コメント検索
let searchs = $('#search');
function search() {
    let commentcolor = $('h4');
    let commentname = $('.h5');

    commentcolor.css("color", "black");
    commentname.css("color", "black");


    $(`h4:contains(${searchs.val()})`).css("color", "blue");
}



