function product_query(e) {
  
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'product_query_all') {
    postStr = 'topic=' + encodeURIComponent('product');
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
    postStr = 'topic=' + encodeURIComponent('product');
    postStr = postStr + '&tt=' + encodeURIComponent('some');
  postStr= postStr + '&prdtid='+encodeURIComponent(document.getElementById("prdtid").value.replace(/\s+/g,""));
  postStr= postStr + '&prdtname='+encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g,""));
  postStr= postStr + '&inventor='+encodeURIComponent(document.getElementById("inventor").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      newContent += '<tr class="table_title">';
      newContent += '<td>产品编码</td>';
      newContent += '<td>产品名称</td>';
      newContent += '<td>规格</td>';
      newContent += '<td>生产企业</td>';
      newContent += '<td>单位</td>';
      newContent += '<td>件装量</td>';
      newContent += '<td>产品线</td>';
      newContent += '<td>发票类型</td>';
      newContent += '<td>备注</td>';
      newContent += '<td>操作</td>';
      newContent += '</tr>';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].PrdtId + '">';
        newContent += '<td name="prdtid">' + responseObject[i].PrdtId + '</td>';
        newContent += '<td name="prdtname">' + responseObject[i].PrdtName + '</td>';
        newContent += '<td name="specific">' + responseObject[i].Specific + '</td>';
        newContent += '<td name="inventor">' + responseObject[i].Inventor + '</td>';
        newContent += '<td name="unit">' + responseObject[i].Unit + '</td>';
        newContent += '<td name="boxnumb">' + responseObject[i].BoxNumb + '</td>';
        newContent += '<td name="inline">' + responseObject[i].Inline + '</td>';
        newContent += '<td name="ivtype">' + responseObject[i].IvType + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(`' + responseObject[i].PrdtId + '`)" value="修改" /> <input type="button" class="del" onclick="del(`'+  responseObject[i].PrdtId + '`)" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('product_results').innerHTML = newContent;
    }
  };

  xhr.open("POST", "/selectitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};



var el = document.getElementById('product_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    product_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    product_query(e);
  });
}

var el2 = document.getElementById('product_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    product_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    product_query(e);
  });
}

function upt(prdtid) {
  let domStr='<legend><span>修改记录</span></legend>';
  $('#' + prdtid).children().each(function (index, element) {
    if($(element).attr("name") != undefined){
    domStr += '<label>'+index +'：</label><input id="update' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(`' + prdtid + '`)" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  // window.location.hash = "#change_table";
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);

}

function sv(prdtid) {
  $.post("updateitem",{
  topic:'product',
  prdtid:prdtid,
  prdtname : document.getElementById("updateprdtname").value.replace(/\s+/g,""),
  specific : document.getElementById("updatespecific").value.replace(/\s+/g,""),
  inventor : document.getElementById("updateinventor").value.replace(/\s+/g,""),
  unit : document.getElementById("updateunit").value.replace(/\s+/g,""),
  boxnumb : document.getElementById("updateboxnumb").value.replace(/\s+/g,""),
  inline : document.getElementById("updateinline").value.replace(/\s+/g,""),
  ivtype : document.getElementById("updateivtype").value.replace(/\s+/g,""),
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

function del(prdtid) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("deleteitem",
      {
        topic: "product",
        prdtid: prdtid
      },
      function (data, status) {
        if (status == "success") {
          alert(data);
          $("#" +prdtid).hide();

        } else {
          alert("服务器错误，请联系周京成");
        }
      });
  }
  else {
    return;
  }
}