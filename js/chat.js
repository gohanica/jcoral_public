        //ここで上記のインプットをここでのインプットに定義する
        var input = document.getElementById('input');
        //ここで上記のアウトプットをここでのアウトプットに定義する
        var output = document.getElementById('output');
        // ウェブソケットの呼び出し、通信の確立 ws://+URL コンストラクタです
        let path= location.pathname;
        
        var ws = new WebSocket("ws://localhost:8080"+path+"/room");

        //これは通信確立時におこるイベントの関数、これを呼び出して接続の送信をする。
        ws.onopen = function() {
           output.innerHTML += "Connection OK\n";
        };
        //ここでサーバーからメッセージを受け取る,eにはいってるよ,outputにも入ってるはず
        ws.onmessage = function(e) {
           var msg = JSON.parse(e.data)
           output.innerHTML += "<image src="+msg.AvatarURL+">"+msg.Name+msg.Message+msg.When+"\n"

        };
        //ここでサーバーにメッセージを送信,JSON.stringify() メソッドは、ある JavaScript のオブジェクトや値を JSON 文字列に変換します,それでテキスト入れ込む。
        function send() {
            ws.send(JSON.stringify(
                {
                    message: input.value
                }
            ));
            //ここでクオーテーションマークつける
            input.value = "";
        };