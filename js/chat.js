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
    $("<div></div>", {
        "class": "userinfo",
    }).appendTo("body");
    $("<img>", {
        src: comment.profileimage,
        width: 40,
        height: 40,
    }).appendTo(".userinfo");

    $("<span></span>", {
        "class": "name",
    }).appendTo(".userinfo");
    $("<span></span>", {
        "class": "time",
    }).appendTo(".userinfo");
    $("<span></span>", {
        "class": "id",
    }).appendTo(".userinfo");
    $(".name").text(userdata.Username);
    $(".time").text(comment.time);
    $(".id").text(comment.id);

    // コメント
    $("<div></div>", {
        "class": "comment",
    }).appendTo("body");
    $('.comment').text(userdata.Message + "<br>");


    // let newdiv = document.createElement('div');
    // let textdiv = document.createElement('div');
    // let buttondiv = document.createElement('div');
    // document.body.appendChild(newdiv);
    // document.body.appendChild(textdiv);
    // document.body.appendChild(buttondiv);

    // // span・img・inputタグ用意
    // let newimage = document.createElement('img');
    // let newname = document.createElement('span');
    // let newtime = document.createElement('span');
    // let newid = document.createElement('span');
    // let newgood = document.createElement('input');
    // let newbad = document.createElement('input');

    // // divタグの中にspanタグinputタグ追加
    // newdiv.appendChild(newimage);
    // newdiv.appendChild(newname);
    // newdiv.appendChild(newtime);
    // newdiv.appendChild(newid);
    // buttondiv.appendChild(newgood);
    // buttondiv.appendChild(newbad);

    // // ユーザー情報
    // newname.setAttribute("class", "newname");
    // newtime.setAttribute("class", "newtime");
    // newid.setAttribute("class", "newid");


    // // 追加するもの準備
    // //プロフィール画像 
    // newimage.src = comment.profileimage;
    // newimage.setAttribute("width", "40");
    // newimage.setAttribute("height", "40");
    // newimage.setAttribute("align", "middle");
    // // ユーザー情報・コメント
    // let children = document.body.childNodes;
    // let divname = document.createTextNode(userdata.Username);
    // let divtime = document.createTextNode(comment.time);
    // let divid = document.createTextNode(comment.id);
    // textdiv.innerHTML = userdata.Message + ":commentclass:" + comments.length + " Bodychild:" + children.length;
    // textdiv.setAttribute("class", "comment");
    // // いいねだめだね機能
    // newgood.setAttribute("type", "button");
    // newbad.setAttribute("type", "button");
    // newgood.setAttribute("onclick", "goodcount()");
    // newbad.setAttribute("onclick", "badcount()");
    // newgood.setAttribute("id", "good");
    // newbad.setAttribute("id", "bad");

    // // spanタグの中に情報追加
    // newname.appendChild(divname);
    // newtime.appendChild(divtime);
    // newid.appendChild(divid);




    // ↑onmesssageここまで
};



// イイね機能
// let goodcounts = 0;
// let badcounts = 0;
// function goodcount() {
//     goodcounts++;
//     document.getElementById("good").value = goodcounts;

// }
// function badcount() {
//     badcounts++;
//     document.getElementById("bad").value = badcounts;
// }

