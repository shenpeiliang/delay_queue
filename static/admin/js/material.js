$(function () {
    $('.open-material-page').click(function () {
        var url = $(this).attr('data-url')
        var type = $(this).attr('data-type')

        //结果处理方式
        var result_handle = $(this).attr('data-result-handle') || 'replace'

        //保存结果的隐藏域
        var target_name = $(this).attr('data-target') || 'msg_value'

        if (type != 'user')
            url += '?type=' + type

        layer.open({
            type: 2,
            area: ['1000px', '600px'],
            title: '选择素材',
            content: url,
            btn: ['确定', '取消'],
            btn1: function (index) {
                var ifram_body = layer.getChildFrame('body', index);
                var item_length = $(ifram_body).find('.checkitem:checked').length
                if (!item_length)
                    return layer.msg('请选择后再操作')

                var res_html = ''

                var res_val = new Array()

                if(result_handle == 'append')
                    res_val = $("input[name='" + target_name + "']").val().split(',')


                $(ifram_body).find('.checkitem:checked').each(function (i) {
                    var msg_id = $(this).val()
                    //不可重复选择素材
                    if(result_handle == 'append'){
                        if(res_val.length && $.inArray(msg_id, res_val) >= 0)
                            return layer.msg('请勿选择重复的素材')
                    }

                    res_val.push(msg_id)

                    res_html += strategies_html[type]($(this), result_handle)

                })

                if(res_val.length > 8) {
                    return layer.msg('图文素材最多为8个')
                }

                $("input[name='" + target_name + "']").val(res_val.join(','))

                //结果可视处理
                if (result_handle == 'append') {
                    var html_txt = res_html
                    $('.material-news-list').append(html_txt)
                } else {
                    if (type == 'user') {
                        $("#user-image").html(res_html)
                    } else {
                        var html_txt = '<ul class="material-base material-news"><li>'
                            + res_html
                            + '</li></ul>'
                        $(".select-material[data-type='" + type + "']").html(html_txt)
                    }
                }

                layer.close(index)
            },
            btn2: function (index) {
                layer.close(index)
            }
        })

    })

    /**
     * 策略 内容获取
     * @type {{user: strategies.'user', 1: strategies.'1'}}
     */
    var strategies_html = {
        'user': function (obj, result_handle) {
            return obj.parents('tr').find('td:eq(4)').html()
        },
        '1': function (obj, result_handle) {
            return '<p>' + obj.parents('tr').find('td:eq(2)').html() + '</p>'
        },
        '2': function (obj, result_handle) {
            return obj.parents('tr').find('td:eq(2)').html()
        },
        '3': function (obj, result_handle) {
            if (result_handle == 'append') {
                return '<li data-id="' + obj.parents('tr').find('td:eq(1)').html() + '"><p>' + obj.parents('tr').find('td:eq(2)').html() + '</p>'
                    + obj.parents('tr').find('td:eq(3)').html() + '<a href="javascript:;" class="delete-handle"></a></li>'
            } else {
                return '<p>' + obj.parents('tr').find('td:eq(2)').html() + '</p>'
                    + obj.parents('tr').find('td:eq(3)').html()
            }

        },
        '4': function (obj, result_handle) {
            return '<p>' + obj.parents('tr').find('td:eq(2)').html() + '</p>'
                + obj.parents('tr').find('td:eq(3)').html()
        },
        '5': function (obj, result_handle) {
            return '<p>' + obj.parents('tr').find('td:eq(2)').html() + '</p>'
                + obj.parents('tr').find('td:eq(3)').html()
        }
    }

    //删除素材
    $('.material-news-list').on('click', '.delete-handle', function () {
        $(this).parent('li').remove()

        var id_del = $(this).parent('li').attr('data-id')

        var ids_now = $("input[name='msg_value']").val().split(',')

        for(var i = 0; i<ids_now.length; i++){
            if(ids_now[i] == id_del)
                ids_now.splice(i, 1)
        }

        $("input[name='msg_value']").val(ids_now.join(','))
    })

    //预览
    $('#handle_preview').click(function(){
        var url = $(this).attr('data-url')

        var user = $("input[name='user']").val()
        if(!user.length)
            return layer.msg('请先自定义接收人')

        var data = $("form").serialize()

        $.post(url, data, function (res) {
           if(res.code)
               layer.msg(res.data)
        }, 'json')
    })
})

