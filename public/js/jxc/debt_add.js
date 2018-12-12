function addDebt() {
  let t2f = false;
  $("input").each(function() {
    t2f = checkInput(this);
    if (t2f) {
      return false;
    }
  });
  if (t2f) {
    alert('所有内容必须都填，请检查后重新提交');
		return;
  }
  var postStr;
  postStr= 'topic=' + encodeURIComponent('debt');
  postStr= postStr + '&srcname='+encodeURIComponent(document.getElementById("srcname").value.replace(/\s+/g,""));
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&dbtsum='+encodeURIComponent(document.getElementById("dbtsum").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));

  let cstmSelects = getCustomerName();
  $("#cstmname").autocomplete({
    source: cstmSelects
  });
  $("#srcname").autocomplete({
    source: cstmSelects
  });
  var xhr = new XMLHttpRequest();
  
  xhr.onload = function () {
    if (xhr.status == 200) {
      $('#rsts').html(xhr.responseText);
    }
  };
  
  xhr.open("POST", "/insertitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};

$("#pmdate").datepicker();

var el = document.getElementById('debt_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addDebt(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addDebt(e);
  });
}