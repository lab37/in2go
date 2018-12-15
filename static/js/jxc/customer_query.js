function customer_query(e) {
  
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'customer_query_all') {
    postStr = 'topic=' + encodeURIComponent('customer');
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
    postStr = 'topic=' + encodeURIComponent('customer');
    postStr = postStr + '&tt=' + encodeURIComponent('some');
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&cstmtype='+encodeURIComponent(document.getElementById("cstmtype").value.replace(/\s+/g,""));
  postStr= postStr + '&gname='+encodeURIComponent(document.getElementById("gname").value.replace(/\s+/g,""));
  postStr= postStr + '&ivname='+encodeURIComponent(document.getElementById("ivname").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {
      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      newContent += '<tr class="table_title">';
      newContent += '<td>客户编码</td>';
      newContent += '<td>客户名称</td>';
      newContent += '<td>客户类型</td>';
      newContent += '<td>收货地址</td>';
      newContent += '<td>收货人</td>';
      newContent += '<td>收货电话</td>';
      newContent += '<td>收票地址</td>';
      newContent += '<td>收票人</td>';
      newContent += '<td>收票电话</td>';
      newContent += '<td>备注</td>';
      newContent += '<td>操作</td>';
      newContent += '</tr>';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].CstmId + '">';
        newContent += '<td name="cstmid">' + responseObject[i].CstmId + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="cstmtype">' + responseObject[i].CstmType + '</td>';
        newContent += '<td name="gaddr">' + responseObject[i].Gaddr + '</td>';
        newContent += '<td name="gname">' + responseObject[i].Gname + '</td>';
        newContent += '<td name="gphone">' + responseObject[i].Gphone + '</td>';
        newContent += '<td name="ivaddr">' + responseObject[i].IvAddr + '</td>';
        newContent += '<td name="ivname">' + responseObject[i].IvName + '</td>';
        newContent += '<td name="ivphone">' + responseObject[i].IvPhone + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' +"'"+ responseObject[i].CstmId +"'"+ ')" value="修改" /> <input type="button" class="del" onclick="del(' +"'"+ responseObject[i].CstmId +"'"+ ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('customer_results').innerHTML = newContent;
      $('#export').show()
    }
  };
  xhr.open("POST", "/selectitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','customer_results');
});
let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('customer_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    customer_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    customer_query(e);
  });
}

var el2 = document.getElementById('customer_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    customer_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    customer_query(e);
  });
}

function upt(cstmid) {
  let domStr='<legend><span>修改记录</span></legend>';
  $('#' + cstmid).children().each(function (index, element) {
    if($(element).attr("name") != undefined){
    domStr += '<label>'+index +'：</label><input id="update' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(' +"'"+ cstmid +"'"+ ')" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(cstmid) {
  $.post("updateitem",{
  topic:'customer',
  cstmid:cstmid,
  cstmname : document.getElementById("updatecstmname").value.replace(/\s+/g,""),
  cstmtype : document.getElementById("updatecstmtype").value.replace(/\s+/g,""),
  gaddr : document.getElementById("updategaddr").value.replace(/\s+/g,""),
  gname : document.getElementById("updategname").value.replace(/\s+/g,""),
  gphone : document.getElementById("updategphone").value.replace(/\s+/g,""),
  ivaddr : document.getElementById("updateivaddr").value.replace(/\s+/g,""),
  ivname : document.getElementById("updateivname").value.replace(/\s+/g,""),
  ivphone : document.getElementById("updateivphone").value.replace(/\s+/g,""),
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

function del(cstmid) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("deleteitem",
      {
        topic: "customer",
        cstmid: cstmid
      },
      function (data, status) {
        if (status == "success") {
          alert(data);
          $("#" +cstmid).hide();
        } else {
          alert("服务器错误，请联系周京成");
        }
      });
  }
  else {
    return;
  }
}