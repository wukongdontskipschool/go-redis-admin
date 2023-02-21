layui.extend({
	admin: '{/}/static/js/admin',
});
layui.define(['jquery', 'layer', 'admin', 'table'], function (exports) {
	var $ = layui.jquery;
	var admin = layui.admin;
	var table = layui.table;

	window.reloadTable = function () {
		admin.ajax({
			type: "GET",
			url: "/v1/redisList/item",
			data: "",
			success: function (res) {
				console.log(res)
				roleTb(res.data)
			},
		});
	}

	window.delItem = function (obj, id) {
		console.log(obj, id)
		layer.confirm('确认要删除吗？', function (index) {
			admin.ajax({
				type: "DELETE",
				url: "/v1/redisList/item?id=" + id,
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
			{ field: 'ID', title: 'ID', width: 80, fixed: 'left' },
			{ field: 'Desc', title: '名称' },
			{ field: 'Host', title: '地址' },
			{ field: 'Port', title: '端口' },
			{ field: 'Tname', title: '分类' },
			{ field: 'deal', title: '操作', templet: '#dealTpl' }
		]]

		table.render({
			elem: '#items'
			, limit: data.length
			, cols: cols
			, data: data
			, page: true //开启分页
		});
	}

	$(function () {
		reloadTable()
	})
});
