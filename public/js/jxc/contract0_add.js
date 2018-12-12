function addContract0() {
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
  postStr= 'topic=' + encodeURIComponent('contract0');
  postStr= postStr + '&ccid='+encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g,""));
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&vector='+encodeURIComponent(document.getElementById("vector").value.replace(/\s+/g,""));
  
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));


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

$("#ivdata").datepicker();
let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});
var el = document.getElementById('contract0_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addContract0(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addContract0(e);
  });
}