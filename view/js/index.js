const name = document.querySelector("#name");
const date = document.querySelector("#date");
const beginTime = document.querySelector("#begin-time");
const endTime = document.querySelector("#end-time");
const reason = document.querySelector("#reason");
const send = document.querySelector("#send");

function setCookie(cname,cvalue,exdays){
    var d = new Date();
    d.setTime(d.getTime()+(exdays*24*60*60*1000));
    var expires = "expires="+d.toGMTString();
    document.cookie = cname+"="+cvalue+"; "+expires;
}

// 获取指定名称的cookie
function getCookie(cname){
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i].trim();
        if (c.indexOf(name)==0) { return c.substring(name.length,c.length); }
    }
    return "";
}

var na = getCookie("name");

//判断cookie是否存在
c_start = document.cookie.indexOf("name=");
if (c_start == -1) {

} else {
    name.value=na;
}

// date.change(function(){
//     date.attr("value",$(this).val()); //赋值
// });


send.onclick = function Send() {
    let sendEmail = new XMLHttpRequest();
    sendEmail.onload = function () {
        var a = sendEmail.responseText;
        var d = JSON.parse(a);
        window.alert(d.message);
        if (d.code === 200) {
            setCookie("name",name.value,300);
        }
    };
    sendEmail.open(
        "get",
        "http://127.0.0.1:8080/email/send?name=" +
        name.value +
        "&time=" +
        date.value+"  "+beginTime.value+"-"+endTime.value+
        "&reason="+
        reason.value,
        true
    );
    sendEmail.send();
};