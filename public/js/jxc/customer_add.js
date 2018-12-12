function addCustomer() {
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
  postStr= 'topic=' + encodeURIComponent('customer');
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&cstmtype='+encodeURIComponent(document.getElementById("cstmtype").value.replace(/\s+/g,""));
  postStr= postStr + '&gaddr='+encodeURIComponent(document.getElementById("gaddr").value.replace(/\s+/g,""));
  postStr= postStr + '&gname='+encodeURIComponent(document.getElementById("gname").value.replace(/\s+/g,""));
  postStr= postStr + '&gphone='+encodeURIComponent(document.getElementById("gphone").value.replace(/\s+/g,""));
  postStr= postStr + '&ivaddr='+encodeURIComponent(document.getElementById("ivaddr").value.replace(/\s+/g,""));
  postStr= postStr + '&ivname='+encodeURIComponent(document.getElementById("ivname").value.replace(/\s+/g,""));
  postStr= postStr + '&ivphone='+encodeURIComponent(document.getElementById("ivphone").value.replace(/\s+/g,""));
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));


  var xhr = new XMLHttpRequest();
  
  xhr.onload = function () {
    if (xhr.status == 200) {
      $('#rsts').html(xhr.responseText);
    }
  };
  console.log(postStr);
  xhr.open("POST", "http://127.0.0.1:8080/insertitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
};



var el = document.getElementById('customer_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addCustomer(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addCustomer(e);
  });
}