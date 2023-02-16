layui.extend({
	admin: '{/}/static/js/admin',
});
layui.define(['jquery', 'layer', 'admin', 'table', 'form'], function (exports) {
    var $ = layui.jquery;
    var admin = layui.admin;
    var table = layui.table;
    var form = layui.form;

    window.reloadTable = function() {
        admin.ajax({
            type: "GET",
            url: "/v1/admin/rule",
            data: "",
            success: function (res) {
                console.log(res)
                roleTb(res.data)
            },
        });
    }

    window.delItem = function(obj, id) {
        console.log(obj, id)
        layer.confirm('确认要删除吗？', function(index) {
            admin.ajax({
                type: "DELETE",
                url: "/v1/admin/rule?id=" + id,
                data: "",
                success: function (res) {
                    $(obj).parents("tr").remove();
                    layer.msg('已删除!', {
                        icon: 1,
                        time: 1000
                    });
                },
            });
        });
    }

    function roleTb(data) {
        cols = [[ //表头
            {field: 'ID', title: 'ID', width:80, fixed: 'left'},
            {field: 'Desc', title: '权限名称'},
            {field: 'Rule', title: '路由'},
            {field: 'Act', title: '方法'},
            {field: 'deal', title: '操作', templet: '#dealTpl'}
        ]]

        table.render({
            elem: '#items'
            ,limit: 20
            ,cols: cols
            ,data: data
            ,page: true //开启分页
        });
    }

    //监听提交
    form.on('submit(add)', function (data) {
        console.log(data);
        addUser(data.field)
        return false;
    });

    function addUser(data) {
        admin.ajax({
            type: "POST",
            url: "/v1/admin/rule",
            data: data,
            success: function (res) {
                success(res);
            }
        });
    }

    function success(res) {
        layer.alert("增加成功", { icon: 6 }, function () {
            layer.close(layer.index);
            reloadTable();
            $('form')[0].reset();
        });
    }

    $(function(){
        reloadTable()
    })
});
