function addPayment() {
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
  postStr= 'topic=' + encodeURIComponent('payment');
  postStr= postStr + '&ccid='+encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g,""));
  postStr= postStr + '&pmdate='+encodeURIComponent(document.getElementById("pmdate").value.replace(/\s+/g,""));
  postStr= postStr + '&pmsum='+encodeURIComponent(document.getElementById("pmsum").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));


  var xhr = new XMLHttpRequest();
  
  xhr.onload = function () {
    if (xhr.status == 200) {
      $('#rsts').html(xhr.responseText);
    }
  };
  
  xhr.open("POST", "http://127.0.0.1:8080/insertitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};

$("#pmdate").datepicker();

var el = document.getElementById('payment_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addPayment(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addPayment(e);
  });
}