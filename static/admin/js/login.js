$(function () {
    //验证码
    $('#verify img').each(function () {
        var url = $(this).attr('src')
        $(this).click(function () {
            $(this).attr('src', url + '?randow=' + Math.random())
        })

    })
    //表单提交
    $('form').on('submit', function () {
        var obj = $(".validate"),
            obj_length = obj.length;
        for (var i = 0; i < obj_length; i++) {
            if ($(obj[i]).val() == '') {
                $('html, body').animate({scrollTop: 0}, 'slow');
                layer.tips('必填项不能为空！', $(obj[i]))
                $(obj[i]).focus()
                return false
            }
        }

        var form_action = $(this).attr("action"),
            data = $("form").serialize();

        $.post(form_action, data, function (res) {
            if (res.code == 'ok') {
                layer.msg(res.data, {icon: 6}, function (index) {
                    layer.close(index)
                    window.location.href = res.url
                });
            } else {
                layer.msg(res.data, {icon: 5}, function (index) {
                    layer.close(index)
                    $('#verify img').trigger('click')
                });
            }
        }, 'json')
        return false
    })

});
