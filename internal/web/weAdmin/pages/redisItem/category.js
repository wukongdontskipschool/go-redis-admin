var layer
var $
var jQuery
var admin
var treeGird
var table
var form

function del(nodeId) {
	alert(nodeId)
}


layui.extend({
	admin: '{/}/static/js/admin',
	jsonview: '{/}/static/js/extends/jquery.jsonview',
	treeGird: '{/}/lib/layui/lay/treeGird' // {/}的意思即代表采用自有路径，即不跟随 base 路径
});
layui.use(['treeGird', 'jquery', 'admin', 'layer', 'table', 'form', 'jsonview'], function() {
	layer = layui.layer,
	$ = jQuery = layui.jquery,
	admin = layui.admin,
	treeGird = layui.treeGird;
	table = layui.table;
	form = layui.form;

	console.log(window.location)
	console.log(layui.router(window.location.href))
	var url = new URL(window.location.href);
	console.log(url.searchParams.get("type"));
	
	var rdId = 3
	$.ajax({
		type: "GET",
		url: "/v1/redisItem/keys/" + rdId,
		data: "",
		success: function(res) {
			console.log(res)
			success(res);
		},
		error: function(e, msg, codeMsg) {
			console.log(e, msg, codeMsg)
		}
	});

	function success(res) {
		$('#keyList').html('')
		for (var i in res['l']) {
			let li = '<li class="" onclick="getVal(this, ' + rdId + ');" title="' + res['l'][i] + '" style="word-wrap:break-word">' + res['l'][i] + '</li>'
			$('#keyList').append(li)
		}
	}

	//触发行单击事件
	table.on('row(items)', function(obj){
		// console.log(obj.tr) //得到当前行元素对象
		console.log(obj.data) //得到当前行数据
		//obj.del(); //删除当前行
		//obj.update(fields) //修改当前行数据
		$('#content').text(obj.data.val)
		$('#subKey').val(obj.data.subKey)
	});

	form.on('select(showText)', function(data){
		// console.log(data.elem); //得到select原始DOM对象
		console.log(data.value); //得到被选中的值
		// console.log(data.othis); //得到美化后的DOM对象

		if (data.value == 'php-unserialize') {
			let com = $('#content').text()
			try {
				com = PHPUnserialize.unserialize(com);
				console.log(com)
				$('#content').text(JSON.stringify(com))
			} catch (error) {
				console.log(error)
			}
		}

		if (data.value == 'text') {
			let com = $('#content').text()
			com = {a: 'a', b: 'b'}
			try {
				com = PHPSerialize.serialize(com);
				console.log(com)
				// $('#content').text(JSON.stringify(com))
			} catch (error) {
				console.log(error)
			}
		}

		if (data.value == 'json') {
			let com = $('#content').text()
			$("#json").JSONView(JSON.parse(com));
			// a = PHPSerialize.serialize(JSON.parse(com))
		}
	}); 
});

function resetFrom()
{
	$('#seyType').text('')
	$('#items + div').hide()
	$('#subKey').val('')
	$('#content').text('')
}

function getVal(elem, rdId)
{
	var keyy = $(elem).text()
	resetFrom();
	$('#key').val(keyy)

	$.ajax({
		type: "GET",
		url: "/v1/redisItem/getVal/" + rdId + "/" + keyy,
		data: "",
		success: function(res) {
			console.log(res)
			printVal(res)
		},
		error: function(e, msg, codeMsg) {
			console.log(e, msg, codeMsg)
		}
	});
}

function printVal(res)
{
	$('.printVal').html('')
	$('#seyType').text(res.key_type)
	
	if (res.key_type == 'string') {
		$('#content').text(res.data)
	} else if (res.key_type == 'list') {
		let cols = [[ //表头
		  {field: 'id', title: 'index', width:80, fixed: 'left'}
		  ,{field: 'val', title: 'val'}
		]]

		let data = []
		for (var k in res.data) {
			id ++
			let tmep = {'id': k, 'val': res.data[k]}
			data[data.length] = tmep
		}
		printItem(data, cols)
	} else if (res.key_type == 'set') {
		let cols = [[ //表头
		  {field: 'id', title: '序号', width:80, fixed: 'left'}
		  ,{field: 'val', title: 'val'}
		]]

		let data = []
		let id = 0
		for (var k in res.data) {
			id ++
			let tmep = {'id': id, 'val': res.data[k]}
			data[data.length] = tmep
		}
		printItem(data, cols)
	} else if (res.key_type == 'zset') {
		let cols = [[ //表头
		  {field: 'id', title: 'index', width:80, fixed: 'left'}
		  ,{field: 'subKey', title: 'member'}
		  ,{field: 'val', title: 'score'}
		]]

		let data = []
		for (var k in res.data) {
			let tmep = {'id': k, 'subKey': res.data[k]['Member'], 'val': res.data[k]['Score']}
			data[data.length] = tmep
		}
		printItem(data, cols)
	} else if (res.key_type == 'hash') {
		let data = []
		let id = 0
		for (var k in res.data) {
			id ++
			let tmep = {'id': id, 'subKey': k, 'val': res.data[k]}
			data[data.length] = tmep
		}
		printItem(data)
	}
}

function printItem(data, cols)
{
	cols = cols || [[ //表头
		{field: 'id', title: '序号', width:80, fixed: 'left'},
		{field: 'subKey', title: 'key'},
		{field: 'val', title: 'val'}
	]]

	$('#items + div').show()

	table.render({
		elem: '#items'
		,height: 193
		,limit: data.length
		,cols: cols
		,data: data
		// ,url: '../../demo/table/user/-page=1&limit=30.js' //数据接口
		// ,page: true //开启分页
		// ,data: [{'id': 0, 'subkey': 'hiha', 'val': 'hahha'}]
	});
}