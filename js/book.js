jQuery(document).ready(function ($) {

    // 本めくるアクション追加
    $("#flipbook").turn({
    });


    $("#send").on("click", function () {
        // 練習用データ
        let now = new Date();
        let month = now.getMonth() + 1;
        let nowtime = now.getFullYear() + "/" + month + "/" + now.getDate() + "/" + now.getHours() + ":" + now.getMinutes();
        let nowbooktime = now.getFullYear() + "/" + month + "/" + now.getDate();
        let comment = {
            profileimage: "../image/0.png",
            name: "大谷翔平",
            time: nowtime,
            id: 17123456789,
            text: "google最強！"
        }



        // ajaxメッセージ送受信
        $.ajax({

            url: 'http://localhost:8080/book/ajax',
            type: 'POST',
            data: JSON.stringify({
                "message": $("#message").val()

            }),
            dataType: 'json',
            scriptCharset: 'utf-8',
            responseType: 'json',
            timespan: 1000,
            success: function (data, textStatus) {
                // レスポンス→json化
                var jsondata = JSON.stringify(data);
                var jsdata = JSON.parse(jsondata);
                console.log(jsondata + textStatus);


                // ページに含まれるメッセージの個数調整
                // <* * を調整することで個数変更＆cssで大きさ変更も忘れないこと
                if ($("ul:last li").length < 6) {

                    // 一つのコメント作 成
                    // ユーザープロフィール
                    // list-group-item追加
                    let userlist3 = $("<li></li>");
                    userlist3.addClass("list-group-item");
                    userlist3.appendTo(".list-group-flush:last");


                    // conteiner追加
                    let usercon = $("<div></div>");
                    usercon.addClass("container");
                    usercon.appendTo(".list-group-item:last");

                    // row追加
                    let userinfo = $("<div></div>");
                    userinfo.addClass("row");
                    userinfo.appendTo(".container:last");

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

                    // ユーザー情報・コメント準備
                    let usertext = $("<div></div>");
                    usertext.addClass("media-body");
                    usertext.appendTo(".media:last");
                    let username = $("<div></div>");
                    username.addClass("userdata");
                    username.appendTo(".media-body:last");

                    // ユーザー情報＋コメント表示
                    username.html(comment.name + "\t" + comment.time + "\tID:" + comment.id);
                    let userchat = $("<div></div>");
                    userchat.addClass("usercomment");
                    userchat.appendTo(".media-body:last");
                    userchat.text(jsdata.message);

                    // 書き出しをDOM移動
                    $('#userwrite').insertAfter(".list-group-item:last");
                    console.log("if");

                } else {
                    console.log("ellse");
                    // 次ページ以降を作成
                    // 一つのコメント作 成
                    // ユーザープロフィール
                    // list-group-item追加
                    let userlist2 = $("<li></li>");
                    userlist2.addClass("list-group-item");
                    userlist2.appendTo(".list-group-flush:last");


                    // // conteiner追加
                    let usercon2 = $("<div></div>");
                    usercon2.addClass("container");
                    usercon2.appendTo(".list-group-item:last");

                    // // row追加
                    let userinfo2 = $("<div></div>");
                    userinfo2.addClass("row");
                    userinfo2.appendTo(".container:last");

                    // // col追加
                    let usercol2 = $("<div></div>");
                    usercol2.addClass("col");
                    usercol2.appendTo(".row:last");

                    // // media追加
                    let usercomment2 = $("<div></div>");
                    usercomment2.addClass("media");
                    usercomment2.appendTo(".col:last");

                    // // imgタグ追加・プロフィール画像追加
                    let userprofile2 = $("<img>");
                    userprofile2.addClass("mr-3");
                    userprofile2.appendTo(".media:last");
                    userprofile2.attr('src', comment.profileimage);

                    // // ユーザー情報・コメント準備
                    let usertext2 = $("<div></div>");
                    usertext2.addClass("media-body");
                    usertext2.appendTo(".media:last");
                    let username2 = $("<div></div>");
                    username2.addClass("userdata");
                    username2.appendTo(".media-body:last");

                    // // ユーザー情報＋コメント表示
                    username2.html(comment.name + "\t" + comment.time + "\tID:" + comment.id);
                    let userchat2 = $("<div></div>");
                    userchat2.addClass("usercomment");
                    userchat2.appendTo(".media-body:last");
                    userchat2.text(jsdata.message);





                    // // hard作成
                    // ページ追加
                    var hardnum = $(".hard").length;
                    hardnum = hardnum + 1;
                    element = $("<div />", { "class": `p${hardnum} hard` });
                    $("#flipbook").turn("addPage", element);

                    // // h1作成
                    let userh = $("<h1></h1>");
                    userh.appendTo(".hard:last")
                    userh.text(`${nowbooktime}`);

                    // // ul作成
                    let userul = $("<ul></ul>");
                    userul.addClass("list-group-flush");
                    userul.appendTo(".hard:last");

                    // // 書き出しをDOM移動
                    $('#userwrite').appendTo(".list-group-flush:last");


                    console.log("era-");

                }
            },
            error: function () {
                $('#result').html("data2");
            }


        });



    });

});




