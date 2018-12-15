function addInvoice() {
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
  postStr= 'topic=' + encodeURIComponent('invoice');
  postStr= postStr + '&ivid='+encodeURIComponent(document.getElementById("ivid").value.replace(/\s+/g,""));
  postStr= postStr + '&ivdate='+encodeURIComponent(document.getElementById("ivdate").value.replace(/\s+/g,""));
  postStr= postStr + '&ccid='+encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g,""));
  postStr= postStr + '&ivsum='+encodeURIComponent(document.getElementById("ivsum").value.replace(/\s+/g,""));
  postStr= postStr + '&postdate='+encodeURIComponent(document.getElementById("postdate").value.replace(/\s+/g,""));
  postStr= postStr + '&expname='+encodeURIComponent(document.getElementById("expname").value.replace(/\s+/g,""));
  postStr= postStr + '&expnumber='+encodeURIComponent(document.getElementById("expnumber").value.replace(/\s+/g,""));
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

$("#ivdate").datepicker();

var el = document.getElementById('invoice_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addInvoice(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addInvoice(e);
  });
}

let allCcId= getAllCcId();
$("#ccid").autocomplete({
  source: allCcId
});