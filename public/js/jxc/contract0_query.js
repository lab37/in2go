function contract0_query(e) {
 
  let tgt = getTarget(e);
  var postStr;
  if (tgt.id == 'contract0_query_all') {
    postStr = 'topic=' + encodeURIComponent('contract0');
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
    postStr = 'topic=' + encodeURIComponent('contract0');
    postStr = postStr + '&tt=' + encodeURIComponent('some');
    postStr= postStr + '&ccid='+encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g,""));
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&vector='+encodeURIComponent(document.getElementById("vector").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));

  }

  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {

      responseObject = JSON.parse(xhr.responseText);
      var newContent = '';
      newContent += '<tr class="table_title">';
      newContent += '<td>合同号码</td>';
      newContent += '<td>所属企业</td>';
      newContent += '<td>购进售出</td>';
      newContent += '<td>备注</td>';
      newContent += '<td>操作</td>';
      newContent += '</tr>';
      for (var i = 0; i < responseObject.length; i++) {    // Loop through object
        newContent += '<tr id="' + responseObject[i].CcId + '">';
        newContent += '<td name="ccid">' + responseObject[i].CcId + '</td>';
        newContent += '<td name="cstmname">' + responseObject[i].CstmName + '</td>';
        newContent += '<td name="vector">' + (responseObject[i].Vector==0?"购进":"售出") + '</td>';
        newContent += '<td name="remark">' + responseObject[i].Remark + '</td>';
        newContent += '<td>' + '<input type="button" class="chg" onclick="upt(`' + responseObject[i].CcId + '`)" value="修改" /> <input type="button" class="del" onclick="del(`' + responseObject[i].CcId + '`)" value="删除" />' + '</td>';
        newContent += '</tr>';
      }
      // Update the page with the new content
      document.getElementById('contract0_results').innerHTML = newContent;
    }
  };

  xhr.open("POST", "http://127.0.0.1:8080/selectitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};

$("#ivdata").datepicker();

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('contract0_query_some');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    contract0_query(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    contract0_query(e);
  });
}

var el2 = document.getElementById('contract0_query_all');
if (el2.addEventListener) {
  el2.addEventListener('click', function (e) {
    contract0_query(e);
  }, false);
} else {
  el2.attachEvent('onclick', function (e) {
    contract0_query(e);
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
  $.post("updateitem",{
  topic:'contract0',
  id:id,
  ccid : document.getElementById("updateccid").value.replace(/\s+/g,""),
  inventor : document.getElementById("updatecstmname").value.replace(/\s+/g,""),
  vector : document.getElementById("updatevector").value.replace(/\s+/g,""),
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
        topic: "contract0",
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