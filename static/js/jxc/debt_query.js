function debt_query(e) {

  let tgt = getTarget(e);
  var postStr;
  postStr = 'topic=' + encodeURIComponent('debt');
  if (tgt.id == 'debt_query_all') {

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
    postStr = postStr + '&srcname=' + encodeURIComponent(document.getElementById("srcname").value.replace(/\s+/g, ""));
    postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g, ""));
    postStr = postStr + '&remark=' + encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g, ""));
  }



  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].Id + '">';
        newContent += '<td name="srcname">' + responseObject[i].SrcName + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="dbtsum">' + responseObject[i].DbtSum + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(' +"'"+ responseObject[i].Id +"'"+ ')" value="修改" /> <input type="button" class="del" onclick="del(' +"'"+ responseObject[i].Id +"'"+ ')" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('rows').innerHTML = newContent;
      $('#debt_results').show();
      $('#export').show()
    }
  };

  xhr.open("POST", "/selectitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};
$('#debt_results').hide();
$('#export').hide();
$('#export').click(function(){
  table2xlsx('xlsx','debt_results');
});
$("#pmdate").datepicker();

var el = document.getElementById('debt_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    debt_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    debt_query(e);
  });
}

var el2 = document.getElementById('debt_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    debt_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    debt_query(e);
  });
}

function upt(id) {
  let domStr = '<legend><span>修改记录</span></legend>';
  $('#' + id).children().each(function (index, element) {
    if ($(element).attr("name") != undefined) {
      domStr += '<label>' + index + '：</label><input id="update' + $(element).attr("name") + '" value="' + $(element).text() + '">'
    }
  });
  domStr += '<input type="button" class="sv" onclick="sv(' +"'"+ id +"'"+ ')" value="保存" /><input type="button" class="sv" onclick="hd()" value="取消" />';
  $("#change_table").html(domStr);
  $("#change_table").show();
  $("html,body").animate({scrollTop:$("#change_table").offset().top},1000);
}

function sv(id) {
  $.post("updateitem", {
    topic: 'debt',
    id: id,
    srcname: document.getElementById("updatesrcname").value.replace(/\s+/g, ""),
    cstmname: document.getElementById("updatecstmname").value.replace(/\s+/g, ""),
    dbtsum: document.getElementById("updatedbtsum").value.replace(/\s+/g, ""),
    remark: document.getElementById("updateremark").value.replace(/\s+/g, "")
  },
    function (data, status) {
      alert(data);
      if (status == "success") {
        $("#change_table").hide();
      }
    });
}

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});
$("#srcname").autocomplete({
  source: cstmSelects
});

$("#change_table").hide();
function hd() {
  $("#change_table").hide();
}

function del(id) {
  var cfm = confirm("确认要删这条记录吗？");
  if (cfm == true) {
    $.post("deleteitem",
      {
        topic: "debt",
        id: id
      },
      function (data, status) {
        if (status == "success") {
          alert(data);
          $("#" + id).hide();
        } else {
          alert("服务器错误，请联系周京成");
        }
      });
  }
  else {
    return;
  }
}