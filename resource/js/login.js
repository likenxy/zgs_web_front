function login() {
    var user_name = document.getElementById('username').value;
    var password= document.getElementById('password').value;
    // TODO预校验
    if (user_name == "" || password == "") {
        return false
    }
    var md5_password = md5(password); 

    jo = new Object()
    jo.name = user_name
    jo.password = md5_password
    s = JSON.stringify(jo)
    console.log(s)
    
    var httpRequest = new XMLHttpRequest();
    httpRequest.open('POST', "/login", true)
    httpRequest.setRequestHeader("Content-type", "application/json");
    httpRequest.send(s)
    httpRequest.onreadystatechange = function () {
        if (httpRequest.status == 200) {
            console.log(httpRequest.responseText)
            window.location.href='index.html';
        } else {
            console.log(httpRequest.status)
            return false;
        }
    }
    return true
}
