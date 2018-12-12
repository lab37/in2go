function queryContract(e) {

  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'contract_query_all') {
    postStr = 'topic=' + encodeURIComponent('contract');
    postStr = postStr + '&tt=' + encodeURIComponent('all');
  } else {
    let t2f = true;
    $(".input").each(function () {
      t2f = checkInput(this);
      if (!t2f) {
        return false;
      }
    });
    if (t2f) {
      alert('请至少填写一项，查询全部请直接点击查询全部按钮');
      return;
    }
    postStr = 'topic=' + encodeURIComponent('contract');
    postStr = postStr + '&tt=' + encodeURIComponent('some');
    postStr = postStr + '&ccid=' + encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g, ""));
    postStr = postStr + '&srcname=' + encodeURIComponent(document.getElementById("srcname").value.replace(/\s+/g, ""));
    postStr = postStr + '&vector=' + encodeURIComponent(document.getElementById("vector").value.replace(/\s+/g, ""));
    postStr = postStr + '&cctype=' + encodeURIComponent(document.getElementById("cctype").value.replace(/\s+/g, ""));
    postStr = postStr + '&ccdate=' + encodeURIComponent(document.getElementById("ccdate").value.replace(/\s+/g, ""));
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g, ""));
    postStr = postStr + '&specific=' + encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g, ""));
    postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g, ""));
    postStr = postStr + '&remark=' + encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g, ""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      newContent += '<tr class="table_title">';
      newContent += '<td>合同编号</td>';
      newContent += '<td>购进售出</td>';
      newContent += '<td>售出客户</td>';
      newContent += '<td>签定日期</td>';
      newContent += '<td>类型</td>';
      newContent += '<td>客户名称</td>';
      newContent += '<td>产品名称</td>';
      newContent += '<td>规格</td>';
      newContent += '<td>价格</td>';
      newContent += '<td>数量</td>';
      newContent += '<td>备注</td>';
      newContent += '<td>操作</td>';
      newContent += '</tr>';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        // let domStr = '<legend><span>修改记录</span></legend>';
        // domStr += '<label>合同编号：</label><input type="text" id="updateccid"  value="' +  responseObject[i].CcId +'/>'
        // domStr += '<label>购进售出：</label><input type="text" id="updatevector"  value="' +  responseObject[i].Vector +'/>'
        // domStr += '<label>签定日期：</label><input type="text" id="updateccdate"  value="' +  responseObject[i].ccdate +'/>'
        // domStr += '<label>合同类型：</label><input type="text" id="updatecctype"  value="' +  responseObject[i].CcType +'/>'
        // domStr += '<label>客户编码：</label><input type="text" id="updatecstmid"  value="' +  responseObject[i].CstmId +'/>'
        // domStr += '<label>产品编码：</label><input type="text" id="updateprdtid"  value="' +  responseObject[i].PrdtId +'/>'
        // domStr += '<label>销售价格：</label><input type="text" id="updateprice"  value="' +  responseObject[i].Price +'/>'
        // domStr += '<label>销售数量：</label><input type="text" id="updatequantity"  value="' + responseObject[i].Quantity  +'/>'
        // domStr += '<label>备注信息：</label><input type="text" id="updateremark"  value="' +  responseObject[i].Remark +'/>'
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="ccid">' + responseObject[i].CcId + '</td>';
        newContent += '<td name="victor">' + (responseObject[i].Vector==0?"购进":"售出") + '</td>';
        newContent += '<td name="srcname">' + responseObject[i].SrcName + '</td>';
        newContent += '<td name="ccdate">' + responseObject[i].ccdate + '</td>';
        newContent += '<td name="cctype">' + responseObject[i].CcType + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="prdtname">' + responseObject[i].PrdtName + '</td>';
        newContent += '<td name="specific">' + responseObject[i].Specific + '</td>';
        newContent += '<td name="price">' + responseObject[i].Price + '</td>';
        newContent += '<td name="quantity">' + responseObject[i].Quantity + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(`' + responseObject[i].Id + '`)" value="修改" /> <input type="button" class="del" onclick="del(`' + responseObject[i].Id + '`)" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('contract_results').innerHTML = newContent;
    }
  };

  xhr.open("POST", "/selectitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  console.log(postStr);
  xhr.send(postStr);
};

$("#ccdate").datepicker();

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});
$("#srcname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('contract_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    queryContract(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    queryContract(e);
  });
}

var el2 = document.getElementById('contract_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    queryContract(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    queryContract(e);
  });
}


function upt(id) {
  let domStr='<legend><span>修改记录</span></legend>';
  $('#' + id).children().each(function (index, element) {
    if($(element).attr("name") != undefined){
    domStr += '<label>'+index +'：</label><input id="update' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(`' + id + '`)" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(id) {
  $.post("updateitem", {
    topic: 'contract',
    id: id,
    ccid: document.getElementById("updateccid").value.replace(/\s+/g, ""),
    ccdate: document.getElementById("updateccdate").value.replace(/\s+/g, ""),
    cctype: document.getElementById("updatecctype").value.replace(/\s+/g, ""),
    cstmname: document.getElementById("updatecstmname").value.replace(/\s+/g, ""),
    prdtname: document.getElementById("updateprdtname").value.replace(/\s+/g, ""),
    specific: document.getElementById("updatespecific").value.replace(/\s+/g, ""),
    price: document.getElementById("updateprice").value.replace(/\s+/g, ""),
    quantity: document.getElementById("updatequantity").value.replace(/\s+/g, ""),
    remark: document.getElementById("updateremark").value.replace(/\s+/g, "")
  },
    function (data, status) {
      alert(data);
      if (status == "success") {
        $("#change_table").hide();
      }
    });
}

$("#change_table").hide();

function hd() {
  $("#change_table").hide();
}

function del(id) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("deleteitem",
      {
        topic: "contract",
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