function htmlDecode(text) {
    //1.首先动态创建一个容器标签元素，如DIV
    var temp = document.createElement("div");
    //2.然后将要转换的字符串设置为这个元素的innerHTML(ie，火狐，google都支持)
    text = eval("'"+text+"'")
    // temp.innerHTML = text;
    // //3.最后返回这个元素的innerText(ie支持)或者textContent(火狐，google支持)，即得到经过HTML解码的字符串了。
    // var output = temp.innerText || temp.textContent;
    // temp = null;
    return filterXSS(text);
}