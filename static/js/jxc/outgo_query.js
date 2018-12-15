function outgo_query(e) {
  
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'outgo_query_all') {
    postStr = 'topic=' + encodeURIComponent('outgo');
    postStr = postStr + '&tt=' + encodeURIComponent('all');
  } else {
    let t2f = true;
  $(".input").each(function() {
    t2f = checkInput(this);
    if (!t2f) {
      return false;
    }
  });
  if (t2f) {
    alert('请至少填写一项，查询全部请直接点击查询全部按钮');
		return;
  }
    postStr = 'topic=' + encodeURIComponent('outgo');
    postStr = postStr + '&tt=' + encodeURIComponent('some');
    postStr = postStr + '&ccid=' + encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g,""));
    postStr = postStr + '&srcname=' + encodeURIComponent(document.getElementById("srcname").value.replace(/\s+/g,""));
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
    postStr = postStr + '&ogdate=' + encodeURIComponent(document.getElementById("ogdate").value.replace(/\s+/g,""));
    postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g,""));
    postStr = postStr + '&specific=' + encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g,""));
    postStr = postStr + '&pnumber=' + encodeURIComponent(document.getElementById("pnumber").value.replace(/\s+/g,""));
    postStr = postStr + '&remark=' + encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      newContent += '<tr class="table_title">';
      newContent += '<td>合同编号</td>';
      newContent += '<td>仓库名称</td>';
      newContent += '<td>客户名称</td>';
      newContent += '<td>出库日期</td>';
      newContent += '<td>产品名称</td>';
      newContent += '<td>规格</td>';
      newContent += '<td>价格</td>';
      newContent += '<td>数量</td>';
      newContent += '<td>批号</td>';
      newContent += '<td>物流名称</td>';
      newContent += '<td>物流单号</td>';
      newContent += '<td>备注</td>';
      newContent += '<td>操作</td>';
      newContent += '</tr>';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="ccid">' + responseObject[i].CcId + '</td>';
        newContent += '<td name="srcname">' + responseObject[i].SrcName + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="ogdate">' + responseObject[i].OgDate + '</td>';
        newContent += '<td name="prdtname">' + responseObject[i].PrdtName + '</td>';
        newContent += '<td name="specific">' + responseObject[i].Specific + '</td>';
        newContent += '<td name="price">' + responseObject[i].Price + '</td>';
        newContent += '<td name="quantity">' + responseObject[i].Quantity + '</td>';
        newContent += '<td name="pnumber">' + responseObject[i].Pnumber + '</td>';
        newContent += '<td name="expname">' + responseObject[i].ExpName + '</td>';
        newContent += '<td name="expnumber">' + responseObject[i].ExpNumber + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' +"'"+ responseObject[i].Id +"'"+ ')" value="修改" /> <input type="button" class="del" onclick="del(' +"'"+ responseObject[i].Id +"'"+ ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('outgo_results').innerHTML = newContent;
      $('#export').show();
    }
  };

  xhr.open("POST", "/selectitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','outgo_results');
});
$("#ogdate").datepicker();

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('outgo_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    outgo_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    outgo_query(e);
  });
}

let allCcId= getAllCcId();
$("#ccid").autocomplete({
  source: allCcId
});

var el2 = document.getElementById('outgo_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    outgo_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    outgo_query(e);
  });
}

function upt(id) {
  let domStr='<legend><span>修改记录</span></legend>';
  $('#' + id).children().each(function (index, element) {
    if($(element).attr("name") != undefined){
    domStr += '<label>'+index +'：</label><input id="update' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(' +"'"+ id +"'"+ ')" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(id) {
  $.post("updateitem",{
  topic:'outgo',
  id:id,
  ccid : document.getElementById("updateccid").value.replace(/\s+/g,""),
  ogdate : document.getElementById("updateogdate").value.replace(/\s+/g,""),
  prdtname : document.getElementById("updateprdtname").value.replace(/\s+/g,""),
  specific : document.getElementById("updatespecific").value.replace(/\s+/g,""),
  quantity : document.getElementById("updatequantity").value.replace(/\s+/g,""),
  pnumber : document.getElementById("updatepnumber").value.replace(/\s+/g,""),
  expname : document.getElementById("updateexpname").value.replace(/\s+/g,""),
  expnumber : document.getElementById("updateexpnumber").value.replace(/\s+/g,""),
  remark : document.getElementById("updateremark").value.replace(/\s+/g,"")

  },
  function(data,status){
    alert(data);
    if(status=="success"){
      $("#change_table").hide();
    }
  });
}

$("#change_table").hide();
function hd(){
  $("#change_table").hide();
  }

function del(id) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("deleteitem",
      {
        topic: "outgo",
        id: id
      },
      function (data, status) {
        if (status == "success") {
          alert(data);
          $("#" +id).hide();
        } else {
          alert("服务器错误，请联系周京成");
        }
      });
  }
  else {
    return;
  }
}