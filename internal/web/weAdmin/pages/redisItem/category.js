var layer
var $
var jQuery
var admin
var treeGird
var table
var form
var rdType
var rdId
var rdDb = '0'
var selectValue; // 当前选择内容

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

	// redis type
	rdType = (new URL(window.location.href)).searchParams.getAll('typeId');
	// redis id
	rdId = form.val("rdId");
	
	function getRedisList() {
		admin.ajax({
			type: "GET",
			url: "/v1/redisItem/redisList/" + rdType,
			data: "",
			success: function(res) {
				$('#rdId').html('');
				let li = '<option value="">请选择</option>';
				$('#rdId').append(li);
				for (var i in res['data']) {
					li = '<option value="' + res['data'][i]['ID'] + '">' + res['data'][i]['Desc'] + '</option>';
					$('#rdId').append(li);
				}

				form.render('select', 'rdId');
			}
		});
	}

	function getKeys() {
		$('#keyList').html('');
		resetFrom();
		admin.ajax({
			type: "GET",
			url: "/v1/redisType/" + rdType + "/redisItem/" + rdId + "/keys?db=" + rdDb,
			data: "",
			success: function(res) {
				for (var i in res['l']) {
					let li = '<li class="" onclick="getVal(this, ' + rdId + ');" title="' + res['l'][i] + '" style="word-wrap:break-word">' + res['l'][i] + '</li>'
					$('#keyList').append(li)
				}
			}
		});
	}

	//触发行单击事件
	table.on('row(items)', function(obj){
		// console.log(obj.tr) //得到当前行元素对象
		// console.log(obj.data) //得到当前行数据
		// obj.del(); //删除当前行
		// obj.update(fields) //修改当前行数据
		selectValue = obj.data.val;
		$('#content').text(obj.data.val)
		$('#subKey').val(obj.data.subKey)
	});

	// 切换显示格式
	form.on('select(showText)', function(data){
		// console.log(data.elem); //得到select原始DOM对象
		// console.log(data.value); //得到被选中的值
		// console.log(data.othis); //得到美化后的DOM对象

		if (data.value == 'php-unserialize') {
			let com = selectValue;
			try {
				com = PHPUnserialize.unserialize(com);
				// PHPSerialize.serialize(com);
				$('#content').text(JSON.stringify(com))
			} catch (error) {
				console.log(error)
			}
		}

		if (data.value == 'text') {
			$('#content').text(selectValue);
		}

		if (data.value == 'json') {
			let com = selectValue;
			$("#json").JSONView(JSON.parse(com));
			// a = PHPSerialize.serialize(JSON.parse(com))
		}
	}); 

	// 切换redis
	form.on('select(rdIdSelect)', function(data) {
		rdId = data.value
		getKeys()
	})

	// 切换库
	form.on('select(dbSelect)', function(data) {
		rdDb = data.value
		getKeys()
	})

	getRedisList();
});

function resetFrom()
{
	$('#valInfo')[0].reset();
	$('#seyType').text(''); // key 类型
	$('#items + div').html(''); // 子项列表
	$('#items + div').hide();
	$('#subKey').val(''); // 子key
	$('#subKey').hide();
	$('#content').text('');
	selectValue = ""; // 选中内容置空
}

function getVal(elem, rdId)
{
	var keyy = $(elem).text()
	resetFrom();

	$('#key').val(keyy)

	admin.ajax({
		type: "GET",
		url: "/v1/redisType/" + rdType + "/redisItem/" + rdId + "/val?key=" + keyy + "&db=" + rdDb,
		data: "",
		success: function(res) {
			printVal(res)
		}
	});
}

function printVal(res)
{
	$('#seyType').text(res.keyType)

	if (res.keyType == 'string') {
		selectValue = res.data;
		$('#content').text(res.data);
	} else if (res.keyType == 'list') {
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
	} else if (res.keyType == 'set') {
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
	} else if (res.keyType == 'zset') {
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
		$('#subKey').show();
		printItem(data, cols)
	} else if (res.keyType == 'hash') {
		let data = []
		let id = 0
		for (var k in res.data) {
			id ++
			let tmep = {'id': id, 'subKey': k, 'val': res.data[k]}
			data[data.length] = tmep
		}
		$('#subKey').show();
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

	// 子项列表框显示
	$('#items + div').show(); 

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