var c = document.cookie.indexOf('lcookie=');
console.log(c)
if(c == -1) {
    window.location.href='resource/login.html';
}

