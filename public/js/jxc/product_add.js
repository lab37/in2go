function addProduct() {
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
  postStr= 'topic=' + encodeURIComponent('product');
  postStr= postStr + '&prdtid='+encodeURIComponent(document.getElementById("prdtid").value.replace(/\s+/g,""));
  postStr= postStr + '&prdtname='+encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g,""));
  postStr= postStr + '&specific='+encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g,""));
  postStr= postStr + '&inventor='+encodeURIComponent(document.getElementById("inventor").value.replace(/\s+/g,""));
  postStr= postStr + '&unit='+encodeURIComponent(document.getElementById("unit").value.replace(/\s+/g,""));
  postStr= postStr + '&boxnumb='+encodeURIComponent(document.getElementById("boxnumb").value.replace(/\s+/g,""));
  postStr= postStr + '&inline='+encodeURIComponent(document.getElementById("inline").value.replace(/\s+/g,""));
  postStr= postStr + '&ivtype='+encodeURIComponent(document.getElementById("ivtype").value.replace(/\s+/g,""));
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



var el = document.getElementById('product_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addProduct(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addProduct(e);
  });
}