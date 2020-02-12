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



    // 実験
    let comments = document.getElementsByClassName("comment");



    let js = JSON.parse(e.data);

    // テンプレートデータ取得整形
    userdata = { Message: js.message, Username: js.username }

    // スレッドDOM追加
    // bodyにdivタグ追加
    // ユーザー情報
    let userinfo = $("<div></div>");
    userinfo.appendTo("body");

    let image = $("<img>");
    image.appendTo(userinfo);
    image.attr('src', comment.profileimage);
    let style = {
        'width': '40',
        'height': '40',
        'align': 'middle'
    }
    image.css(style);

    let name = $("<span></span>");
    name.appendTo(userinfo);
    name.text(userdata.Username);
    name.addClass("name");
    let time = $("<span></span>");
    time.appendTo(userinfo);
    time.text(comment.time);
    time.addClass("time");
    let id = $("<span></span>");
    id.appendTo(userinfo);
    id.text(comment.id);
    id.addClass("id");

    let usercomment = $("<div></div>");
    usercomment.appendTo("body");
    usercomment.text(userdata.Message);
    usercomment.addClass("comment");
}

