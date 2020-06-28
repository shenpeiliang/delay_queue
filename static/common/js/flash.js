function chkFlash() {
    var isIE = (navigator.appVersion.indexOf("MSIE") >= 0);
    var hasFlash = true;
    if (isIE) {
        try {
            var objFlash = new ActiveXObject("ShockwaveFlash.ShockwaveFlash");
        } catch (e) {
            hasFlash = false;
        }
    } else {
        if (!navigator.plugins["Shockwave Flash"]) {
            hasFlash = false;
        }
    }
    return hasFlash;
}


//如果没有安装或开启，提示需要安装点击链接自动开启插件（已安装）
function MagFlash() {
    if (!chkFlash()) {
        var str = '<div class="ui-popmsg"><ul class="clearfix"><li>检测到您尚未安装flash player插件，无法使用图片上传功能，请您到 <em class="font-ql">http://get.adobe.com/cn/flashplayer/?fpchrome</em> 下载安装！</li><li class="ui-popmsg-txt-c"><a href="http://get.adobe.com/cn/flashplayer/?fpchrome" target="_blank" class="flash-btn">点击安装</a>或<a href="https://qzonestyle.gtimg.cn/qzone/photo/v7/js/module/flashDetector/flash_tutorial.pdf" target="_blank" class="flash-btn">查看开启方法</a></li></ul></div>';
        layer.open({
            title: '温馨提示',
            content: str
        });
        return false;
    }
    return true;
}


//最后加载时检测提示：
MagFlash();