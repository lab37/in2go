function addContract() {
  let t2f = false;
  $("input").each(function () {
    t2f = checkInput(this);
    if (t2f) {
      return false;
    }
  });
  console.log(t2f);
  if (t2f) {
    alert('所有内容必须都填，请检查后重新提交');
    
    return;
  }
  var postStr;
  postStr = 'topic=' + encodeURIComponent('contract');
  postStr = postStr + '&ccid=' + encodeURIComponent(document.getElementById("ccid").value.replace(/\s+/g,""));
  postStr = postStr + '&ccdate=' + encodeURIComponent(document.getElementById("ccdate").value.replace(/\s+/g,""));
  postStr = postStr + '&cctype=' + encodeURIComponent(document.getElementById("cctype").value.replace(/\s+/g,""));
  postStr = postStr + '&cstmname=' + encodeURIComponent(document.getElementById("cstmname").value.replace(/\s+/g,""));
  postStr = postStr + '&prdtname=' + encodeURIComponent(document.getElementById("prdtname").value.replace(/\s+/g,""));
  postStr = postStr + '&specific=' + encodeURIComponent(document.getElementById("specific").value.replace(/\s+/g,""));
  postStr = postStr + '&price=' + encodeURIComponent(document.getElementById("price").value.replace(/\s+/g,""));
  postStr = postStr + '&quantity=' + encodeURIComponent(document.getElementById("quantity").value.replace(/\s+/g,""));
  postStr = postStr + '&remark=' + encodeURIComponent(document.getElementById("remark").value.replace(/\s+/g,""));


  var xhr = new XMLHttpRequest();

  xhr.onload = function () {
    if (xhr.status == 200) {
      $('#rsts').html(xhr.responseText);
    }
  };

  xhr.open("POST", "/insertitem", true);
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  xhr.send(postStr);
};
let cstmSelects = getCustomerName();
$("#cstmname").autocomplete({
  source: cstmSelects
});

let allCcId= getAllCcId();
$("#ccid").autocomplete({
  source: allCcId
});

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

$("#ccdate").datepicker();


var el = document.getElementById('contract_add');
if (el.addEventListener) {
  el.addEventListener('click', function (e) {
    addContract(e);
  }, false);
} else {
  el.attachEvent('onclick', function (e) {
    addContract(e);
  });
}