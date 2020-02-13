
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

    // 一つのコメント作成
    // ユーザープロフィール
    let userinfo = $("<div></div>");
    userinfo.appendTo("body");
    let image = $("<img>");
    image.appendTo(userinfo);
    image.attr('src', comment.profileimage);
    image.attr('align', 'middle');
    let style = {
        'width': '40',
        'height': '40',

    }
    image.css(style);

    // ユーザー情報
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

    // コメント
    let usercomment = $("<div></div>");
    usercomment.appendTo("body");
    usercomment.text(userdata.Message);
    usercomment.addClass("comment");

    let elements = $("div").toArray();
    console.log(elements);

}
// コメント検索
let searchs = $('#search');
function search() {
    let commentcolor = $('.comment');
    let commentname = $('.name');
    let commenttime = $('.time');
    let commentid = $('.id');
    commentcolor.css("color", "red");
    commentname.css("color", "black");
    commenttime.css("color", "black");
    commentid.css("color", "black");
    $(`div:contains(${searchs.val()})`).css("color", "blue");
}
// 検索窓から離したとき色付け
// $("#search").on("change", function () {
//     let searchs = $('#search');
//     let commentcolor = $('.comment');
//     let commentname = $('.name');
//     let commenttime = $('.time');
//     let commentid = $('.id');
//     commentcolor.css("color", "red");
//     commentname.css("color", "black");
//     commenttime.css("color", "black");
//     commentid.css("color", "black");
//     $(`div:contains(${searchs.val()})`).css("color", "blue");
// })



