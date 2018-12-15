function outwayInvoiceQuery(e) {

  let tgt = getTarget(e);
  var postStr;
  postStr = 'topic=' + encodeURIComponent('onway_invoice');
  postStr += '&vector=1';
  if (tgt.id == 'outway_invoice_query_all') {
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
      newContent += '<td>合同号码</td>';
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
        newContent += '<td name="ccid">' + responseObject[i].CcId + '</td>';
        newContent += '<td name="srcname">' + responseObject[i].SrcName + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="csum">' + responseObject[i].Csum + '</td>';
        newContent += '<td name="cksum">' + responseObject[i].CkSum + '</td>';
        newContent += '<td name="gap">' +(responseObject[i].Csum - responseObject[i].CkSum) + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';

        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('outway_invoice_results').innerHTML = newContent;
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
  table2xlsx('xlsx','outway_invoice_results');
});
let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});
$("#srcname").autocomplete({
  source: cstmSelects
});

let allCcId= getAllCcId();
$("#ccid").autocomplete({
  source: allCcId
});

var el = document.getElementById('outway_invoice_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    outwayInvoiceQuery(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    outwayInvoiceQuery(e);
  });
}

var el2 = document.getElementById('outway_invoice_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    outwayInvoiceQuery(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    outwayInvoiceQuery(e);
  });
}
$("#change_table").hide();