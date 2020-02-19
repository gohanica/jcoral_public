jQuery(document).ready(function ($) {



    $("#send").on("click", function () {
        // 練習用データ
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

        // ajaxメッセージ送受信
        $.ajax({

            url: 'http://localhost:8080/pra',
            type: 'POST',
            data: JSON.stringify({
                "message": $("#message").val()

            }),
            dataType: 'json',
            scriptCharset: 'utf-8',
            responseType: 'json',
            timespan: 1000,
            success: function (data, textStatus) {
                var jsondata = JSON.stringify(data);
                var jsdata = JSON.parse(jsondata);
                console.log(jsondata + textStatus);
                // $('#result').append(jsondata + "<br>");
                // $('#result').append(textStatus + "<br>");

                // $('#result').append(jsdata.message + "<br>");

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
                    'width': '60',
                    'height': '60',
                }
                userprofile.css(profilestyle);
                // ユーザー情報・コメント準備
                let usertext = $("<div></div>");
                usertext.addClass("media-body");
                usertext.appendTo(".media:last");
                let username = $("<h5></h5>");
                username.appendTo(".media-body:last");

                // ユーザー情報＋コメント表示
                username.html(comment.name + "\t" + comment.time + "\tID:" + comment.id);
                let userchat = $("<h4></h4>");
                userchat.appendTo(".media-body:last");
                userchat.text(jsdata.message);

            },
            error: function () {
                $('#result').html("data2");
            }


        });


    });
});




