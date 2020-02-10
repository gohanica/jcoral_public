var now = new Date();
var month = now.getMonth() + 1;
var nowtime = now.getFullYear() + "/" + month + "/" + now.getDate() + "/" + now.getHours() + ":" + now.getMinutes();
var comment = {
    profileimage: "../image/1.png",
    name: "大谷翔平",
    time: nowtime,
    id: 17123456789,
    text: "google最強！"
}

// 要素へデータ入れる
// var userimage = document.getElementById('image');
// var username = document.getElementById('name');
// var usertime = document.getElementById('time');
// var userid = document.getElementById('id');
// var usertext = document.getElementById('text');
// var userinfo = document.getElementById('userinfo');
// userimage.src = comment.profileimage
// username.innerHTML = comment.name;
// usertime.textContent = comment.time;
// userid.textContent = comment.id;
// usertext.textContent = comment.text;



// 送信系HTML取得
var input = document.getElementById('input');
var username = document.getElementById('username');

// websocket通信開始
var socket = new WebSocket("ws://" + window.location.host + "/chat");

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


var userdata = [];
// websocket サーバからメッセージ受信
socket.onmessage = function (e) {
    var js = JSON.parse(e.data);

    // テンプレートデータ取得整形
    userdata = { Message: js.message, Username: js.username }

    //    スレッドDOM追加
    // bodyにdivタグ追加
    var newdiv = document.createElement('div');
    var textdiv = document.createElement('div');
    document.body.appendChild(newdiv);
    document.body.appendChild(textdiv);
    // spanタグ用意
    // var newimagesss = ducument.createElement('img');
    var newimage = document.createElement('img');
    var newname = document.createElement('span');
    var newtime = document.createElement('span');
    var newid = document.createElement('span');

    // divタグの中にspanタグ追加
    newdiv.appendChild(newimage);
    newdiv.appendChild(newname);
    newdiv.appendChild(newtime);
    newdiv.appendChild(newid);


    // 追加するもの準備
    newimage.src = comment.profileimage;
    newimage.setAttribute("width", "40");
    newimage.setAttribute("height", "40");
    newimage.setAttribute("align", "middle");
    var divname = document.createTextNode(userdata.Username);
    var divtime = document.createTextNode(comment.time + " ID:");
    var divid = document.createTextNode(comment.id);
    textdiv.innerHTML = userdata.Message;

    // spanタグの中に情報追加
    newname.appendChild(divname);
    newtime.appendChild(divtime);
    newid.appendChild(divid);
};






