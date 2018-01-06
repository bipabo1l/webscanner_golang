// layui下拉框选择,支持键盘
var listenSelect = {
    style: "layui-this",
    up: function (event, thisinput) {
        var keyCode = event.keyCode;
        var dl = $(thisinput).parent().next(); //找到渲染后的dl
        var curDd = (dl).find('.layui-this');
        if (keyCode == 40) { //按下下键
            $(thisinput).parents('.layui-form-select').addClass('layui-form-selected');
            curDd = $(dl).find('.layui-this').nextAll(':not(.layui-hide)');
            if (curDd.length == 0) {
                curDd = $(dl).find('dd:first');
            } else {
                curDd = curDd[0];
            }
        } else if (keyCode == 38) {
            $(thisinput).parents('.layui-form-select').addClass('layui-form-selected');
            curDd = $(dl).find('.layui-this').prevAll(':not(.layui-hide)');
            if (curDd.length == 0) {
                curDd = $(dl).find('dd:last');
            } else {
                curDd = curDd[0];
            }
        }


        dl.find("dd").removeClass(this.style); //移除样式
        $(curDd).addClass(this.style);

        var dd = $(dl).find('.layui-this');


        // 计算高度--start
        try {
            dd.offset().top - dl.offset().top + dl.scrollTop();
            dl.scrollTop(
                dd.offset().top - dl.offset().top + dl.scrollTop() - 100
            );
        } catch (err) {
            //console.log(err.stack);
        }
        // 计算高度--end
        if (keyCode == 13) {
            $(dd).click();
            $(thisinput).focus(); // 再次得到焦点
            $(thisinput).attr("onkeydown", "listenSelect.up(event,this)")
        }
        return false;
    }
}