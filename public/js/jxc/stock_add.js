function addStock() {
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
  postStr= 'topic=' + encodeURIComponent('stock');
  postStr= postStr + '&cstmname='+encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr= postStr + '&prdtname='+encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g,""));
  postStr= postStr + '&specific='+encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g,""));
  postStr= postStr + '&quantity='+encodeURIComponent(document.getElementById("quantity").value.replace(/\s+/g,""));
  postStr= postStr + '&pnumber='+encodeURIComponent(document.getElementById("pnumber").value.replace(/\s+/g,""));
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

let prdtObjects = getProductNS();
let prdtNames = new Array();
for (var i = 0; i < prdtObjects.length; i++) {
  prdtNames.push(prdtObjects[i].PrdtName);
}
$("#prdtname").autocomplete({
  source: prdtNames
});

$("#specific").focus(function () {
  let prdtSpecifics = new Array();
  
  for (var i = 0; i < prdtObjects.length; i++) {
    if (prdtObjects[i].PrdtName == $("#prdtname").val()) {
      prdtSpecifics.push(prdtObjects[i].Specific);
    }
  }
  $("#specific").autocomplete({
    source: prdtSpecifics
  });
});

$("#icdata").datepicker();

let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

var el = document.getElementById('stock_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addStock(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addStock(e);
  });
}