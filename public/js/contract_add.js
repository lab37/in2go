function getTarget(e) {
  if (!e) {
    e = window.event;
  }
  return e.target || e.srcElement;
}

function itemDone() {

  var postStr;
   postStr= 'topic=' + encodeURIComponent('contract');
  postStr= postStr + '&ccid='+encodeURIComponent(document.getElementById("ccid").value);
  postStr= postStr + '&cctype='+encodeURIComponent(document.getElementById("cctype").value);
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value);
  postStr= postStr + '&prdtname='+encodeURIComponent(document.getElementById("prdtname").value);
  postStr= postStr + '&specific='+encodeURIComponent(document.getElementById("specific").value);
  postStr= postStr + '&price='+encodeURIComponent(document.getElementById("price").value);
  postStr= postStr + '&quantity='+encodeURIComponent(document.getElementById("quantity").value);
  postStr= postStr + '&remark='+encodeURIComponent(document.getElementById("remark").value);


  var xhr = new XMLHttpRequest();
  
  xhr.onload = function () {
    if (xhr.status == 200) {
      responseObject = JSON.parse(xhr.responseText);
      document.getElementById('results').innerHTML = newContent;
    }
  };
  
  xhr.open("POST", "http://127.0.0.1:8080/insertcon", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
  console.log(postStr);
};



var el = document.getElementById('add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    itemDone(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    itemDone(e);
  });
}