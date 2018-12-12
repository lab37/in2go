function inwayInvoiceQuery(e) {

  let tgt = getTarget(e);
  var postStr;
  postStr = 'topic=' + encodeURIComponent('onway_invoice');
  postStr += 'vector=0';
  if (tgt.id == 'inway_invoice_query_all') {
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
    postStr = postStr + '&tt=' + encodeURIComponent('some');
    postStr = postStr + '&ccid=' + encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g, ""));
    postStr = postStr + '&srcname=' + encodeURIComponent(document.getElementById("srcname").value.replace(/\s+/g, ""));
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g, ""));
  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      newContent += '<tr class="table_title">';
      newContent += '<td>售方名称</td>';
      newContent += '<td>购方名称</td>';
      newContent += '<td>合同金额</td>';
      newContent += '<td>已开金额</td>';
      newContent += '<td>差额</td>';
      newContent += '<td>备注</td>';
      // newContent += '<td>操作</td>';
      newContent += '</tr>';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="srcname">' + responseObject[i].SrcName + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="csum">' + responseObject[i].Csum + '</td>';
        newContent += '<td name="cksum">' + responseObject[i].CkSum + '</td>';
        newContent += '<td name="gap">' +(responseObject[i].Csum - responseObject[i].Osum) + '</td>';
        newContent += '<td name="remark">' +responseObject[i].Remark + '</td>';

        // newContent += '<td>' + '<input type="button" class="chg" onclick="upt(`' + responseObject[i].Id + '`)" value="修改" /> <input type="button" class="del" onclick="del(`' + responseObject[i].Id + '`)" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('inway_invoice_results').innerHTML = newContent;
    }
  };

  xhr.open("POST", "http://127.0.0.1:8080/selectitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});
$("#srcname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('inway_invoice_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    inwayInvoiceQuery(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    inwayInvoiceQuery(e);
  });
}

var el2 = document.getElementById('inway_invoice_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    inwayInvoiceQuery(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    inwayInvoiceQuery(e);
  });
}

// function upt(id) {
//   let domStr='<legend><span>修改记录</span></legend>';
//   $('#' + id).children().each(function (index, element) {
//     if($(element).attr("name") != undefined){
//     domStr += '<label>'+index +'：</label><input id="update' + $(element).attr("name") + '" value="' + $(element).text() + '">'
//     }
//   });
//   domStr += '<input type="button" class="sv" onclick="sv(`' + id + '`)" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
//   console.log(domStr);
//   $("#change_table").html(domStr);
//   $("#change_table").show();
//   window.location.hash = "#change_table";
// }

// function sv(id) {
//   $.post("updateitem",{
//   topic:'inway_invoice',
//   id:id,
//   cstmname : document.getElementById("updatecstmname").value.replace(/\s+/g,""),
//   prdtname : document.getElementById("updateprdtname").value.replace(/\s+/g,""),
//   specific : document.getElementById("updatespecific").value.replace(/\s+/g,""),
//   quantity : document.getElementById("updatequantity").value.replace(/\s+/g,""),
//   pnumber : document.getElementById("updatepnumber").value.replace(/\s+/g,""),
//   remark : document.getElementById("updateremark").value.replace(/\s+/g,"")
//   },
//   function(data,status){
//     alert(data);
//     if(status=="success"){
//       $("#change_table").hide();
//     }
//   });
// }

$("#change_table").hide();
// function hd(){
//   $("#change_table").hide();
//   }

// function del(id) {
//   var cfm = confirm("确认要删这条记录吗？");
//   if (cfm == true) {
//     $.post("deleteitem",
//       {
//         topic: "inway_invoice",
//         id: id
//       },
//       function (data, status) {
//         if (status == "success") {
//           alert(data);
//           $("#" +id).hide();
//         } else {
//           alert("服务器错误，请联系周京成");
//         }
//       });
//   }
//   else {
//     return;
//   }
// }