//只能输入数字
function is_num_int(obj) {   // 值允许输入数字
    obj.value = obj.value.replace(/[^\d]/g, "")     //先把非数字的都替换掉，除了数字
    obj.value = obj.value.replace(/^0{0,}/, "")
}

//只能输入数字和一个小数点
function is_num(obj) {   // 值允许输入一个小数点和数字
    obj.value = obj.value.replace(/[^\d.]/g, "")     //先把非数字的都替换掉，除了数字和.
    obj.value = obj.value.replace(/^\./g, "")         //必须保证第一个为数字而不是.
    obj.value = obj.value.replace(/\.{2,}/g, ".")   //保证只有出现一个.而没有多个.
    obj.value = obj.value.replace(".", "$#$").replace(/\./g, "").replace("$#$", ".")    //保证.只出现一次，而不能出现两次以上
}

(function ($) {
    $.fn.select_radio = function (obj, end_obj) {
        obj.each(function () {
            $(this).click(function () {
                var v = $(this).attr('is_show')
                end_obj.val(v)
                $(this).parent().children('.r_selected').removeClass('r_selected')
                $(this).addClass('r_selected')
            })
        })

    }
})(jQuery)
$(function () {
    //选项显示
    $('.option-show-event').click(function () {
        var option_val = $(this).find("input[name='is_safe_mode']:checked").val()
        console.log('==', option_val)
        if (option_val == 1) {
            $('.option-validate').removeClass('option-hidden')
            $('.option-validate').find('.txt_basic').addClass('validate')
        } else {
            $('.option-validate').addClass('option-hidden')
            $('.option-validate').find('.txt_basic').removeClass('validate')
        }
    })
    //添加授权用户
    $('#grant_user').on('click', function () {
        var aid = $(this).attr('data-aid')
        layer.open({
            type: 1,
            title: '授权用户',
            btn: ['确定', '取消'],
            content: $('#grant_user_content'),
            btn2: function (index, layero) {
                layer.close(index)
            },
            btn1: function (index, layero) {
                var uids = $('#grant_user_content > textarea').val()
                if (!uids.length)
                    return layer.msg('请输入管理员ID')
                $.post(grant_uri, {
                    aid: aid,
                    uids: uids
                }, function (res) {
                    if (res.code == 'success') {
                        layer.open({
                            title: '系统提示',
                            content: res.data,
                            btn: ['确定'],
                            btn1: function (index, layero) {
                                layer.close(index)
                                window.location.reload()
                            },
                        })
                    } else {
                        layer.open({
                            title: '系统提示',
                            content: res.data,
                            btn: ['确定'],
                            btn1: function (index, layero) {
                                layer.close(index)
                            },
                        })
                    }
                })
            }
        })

    })

    //折叠菜单
    $('.is_show').click(function () {
        var status = $(this).find('dd').slideToggle()
    })


    //是否为空
    function check_empty(obj) {
        if (typeof (obj) == "undefined" || obj <= 0) {
            return true
        } else {
            return false
        }
    }


    //删除
    $('.event-delete').click(function () {
        var url = $(this).attr('data-url')

        $.get(url, function (e) {
            if (e.code == 'success') {
                layer.msg(e.data, {time: 2000}, function () {
                    return window.location.reload()
                })
            } else {
                layer.msg(e.data)
            }
        }, 'json')
    })


    //预览菜单折叠
    $(document).on('click', '#preview-section .footer .content>li', function () {
        var index = $(this).index()
        $(this).find('.content-child').slideToggle()
        var obj = $(this).parents('.content').children('li')
        for (var i = 0; i < obj.length; i++) {
            if ($(obj[i]).index() != index) {
                $(obj[i]).find('.content-child').hide()
            }
        }
    })

    //表单提交
    $('#form-submit').click(function () {
        var obj = $('.validate')//必填
        var obj_length = obj.length
        for (var i = 0; i < obj_length; i++) {
            if ($(obj[i]).val() == '') {
                $('html, body').animate({scrollTop: 0}, 'slow')
                if ($(obj[i]).attr('type') == 'hidden') {
                    layer.tips('必填项不能为空', $(obj[i]).parent())
                    return false
                }
                layer.tips('必填项不能为空！', $(obj[i]))
                $(obj[i]).focus()
                return false
            }
        }

        var data = $("form").serialize()
        var url = $("form").attr('action')

        $.post(url, data, function (res) {
            if (res.code == 'ok')
                return layer.msg(res.data, function () {
                    if (res.url)
                        return window.location.href = res.url
                    else
                        return window.location.reload()
                })
            else
                return layer.msg(res.data)
        }, 'json')
    })

    //队列执行记录
    $('.open-detail-log').click(function () {
        var url = $(this).attr('data-uri')

        layer.open({
            type: 2,
            area: ['1000px', '600px'],
            title: '执行记录',
            content: url,
            btn: ['关闭'],
            shadeClose: true,
            btn1: function (index) {
                layer.close(index)
            }
        })

    })
})