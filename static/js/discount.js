function hetong_addItem() {
	// Create a new element and store it in a variable.
	var newTr = document.createElement('tr');
	var position = document.getElementById('hetong');
	position.appendChild(newTr);
	var newTd = document.createElement('td');
	var lastTr = position.lastChild;
	for (i=0;i<8;i++) {
	  newTd = document.createElement('td')
	  newInput = document.createElement('input')
	  newTd.appendChild(newInput);
	  if (i>5) {
		newInput.disabled = "true";
	  }
	  lastTr.appendChild(newTd);
	}
	
};



function hetong_delItem() {

// Find the element which contains the element to be removed.
var containerEl = document.getElementById('hetong');
var removeEl = containerEl.lastChild;
if (containerEl.children.length>1) {
    containerEl.removeChild(removeEl);
}
}

function patch_addItem() {
	// Create a new element and store it in a variable.
	var newTr = document.createElement('tr');
	var position = document.getElementById('patch');
	position.appendChild(newTr);
	var newTd = document.createElement('td');
	var lastTr = position.lastChild;
	for (i=0;i<8;i++) {
	  newTd = document.createElement('td')
	  newInput = document.createElement('input')
	  if (i>5) {
		newInput.disabled = "true";
	  }
	  newTd.appendChild(newInput);
	  lastTr.appendChild(newTd);
	}
	
};



function patch_delItem() {

// Find the element which contains the element to be removed.
var containerEl = document.getElementById('patch');
var removeEl = containerEl.lastChild;
if (containerEl.children.length>1) {
    containerEl.removeChild(removeEl);
}
};

function mFor(a,b,d,fs,bz) {
	var hetongItem = document.querySelectorAll('#hetong tr');
    var patchItem = document.querySelectorAll('#patch tr');
	var totalA=1, totalB=1;
	var w, t, k, m, v;
	var z=1;
	var lastSum=0;
	var thisSum = 0;
    var n = new Array();
	var tempArry = new Array();
	var tempSum1, tempSum2=10000000;
	for (t=0;t<=d;t++) {
	  totalA = totalA*(b[t]-a[t]+1);
      n[t]=a[t];
      tempArry[t]=a[t];	  
	}

    for (t=d+1;t<a.length;t++) {
	  totalB = totalB*(b[t]-a[t]+1);
      n[t]=a[t];
      tempArry[t]=a[t];
	}
	for (w=0;w<totalA;w++) {        
		for (t=d;t>=0;t--){
		    if(n[t]>b[t]) {
			  n[t-1]++;
			  n[t]=a[t];
			}
	    }
		for (v=d+1;v<totalB;v++) {
			for (t=a.length-1;t>d;t--){
		      if(n[t]>b[t]) {
			    n[t-1]++;
			    n[t]=a[t];
			  }
	        }
			
			for (k=0;k<=d;k++) {
			  lastSum = lastSum + n[k]*bz[k];
			}
			for (k=d+1;k<(d+a.length-1)/2+1;k++) {
			  lastSum=lastSum + n[k]*n[k+(a.length-d-1)/2]*bz[k];		
			}
			tempSum1 = Math.abs(lastSum - fs);
			if ((lastSum - fs)>tempSum2) {
				for (t=d+1;t<a.length;t++) {	       
                  n[t]=a[t];
	            }
				
				break;
			}
			if (tempSum2 > tempSum1) {
				tempSum2 = tempSum1;
				thisSum = lastSum;
				for (m=0;m<a.length;m++){
				  tempArry[m]=n[m];
				}
			}
			n[a.length-1]++;
			lastSum = 0;
	    }
		n[d]++;
	}
	
	alert("计算完成");
	
	result2.value = thisSum;
	for (m=0;m<=d;m++) {
	  hetongItem[z].children[7].firstChild.value=tempArry[m];
	
	  z++;
	}
	z=1;
	for (m=0;m<patchItem.length-1;m++) {
	  patchItem[z].children[7].firstChild.value=tempArry[d+1+m];
	  z++;
	}
	z=1;
	for (m=(d+a.length-1)/2+1;m<a.length;m++) {
	  
	  patchItem[z].children[6].firstChild.value=tempArry[m];
	  z++;
	}
	
};
	

function addItem() {
  var i=0, n=0, g=0, h=0;
  var maxArr = new Array();
  var minArr = new Array();
  var perQuanArr = new Array();
  var firstSum=0;
  var hetongItem = document.querySelectorAll('#hetong tr');
  var patchItem = document.querySelectorAll('#patch tr');
  
  
  h =  hetongItem.length-2;
  for (i=1;i<hetongItem.length;i++) {
    firstSum = firstSum + hetongItem[i].children[1].firstChild.value*hetongItem[i].children[2].firstChild.value*hetongItem[i].children[3].firstChild.value;
	minArr.push(Math.floor(hetongItem[i].children[4].firstChild.value));
	maxArr.push(Math.floor(hetongItem[i].children[5].firstChild.value));
	perQuanArr.push((hetongItem[i].children[2].firstChild.value*hetongItem[i].children[1].firstChild.value));
  }
  for (n=1;n<patchItem.length;n++) {
	minArr.push(Math.floor(patchItem[n].children[4].firstChild.value));
	maxArr.push(Math.floor(patchItem[n].children[5].firstChild.value));
	perQuanArr.push(Math.floor(patchItem[n].children[3].firstChild.value));
  }
  
  for (g=1;g<patchItem.length;g++) {
	minArr.push(Math.floor(patchItem[g].children[1].firstChild.value));
	maxArr.push(Math.floor(patchItem[g].children[2].firstChild.value));
  }
  
  result1.value = firstSum;
  mFor(minArr,maxArr,Math.floor(h),Math.floor(firstSum),perQuanArr);
  
};





var hetong_add = document.getElementById('hetong_addItem');
hetong_add.addEventListener('click', hetong_addItem, false);

var hetong_del = document.getElementById('hetong_delItem');
hetong_del.addEventListener('click', hetong_delItem, false);

var patch_add = document.getElementById('patch_addItem');
patch_add.addEventListener('click', patch_addItem, false);

var patch_del = document.getElementById('patch_delItem');
patch_del.addEventListener('click', patch_delItem, false);

var lastAddItem = document.getElementById('addItem');
lastAddItem.addEventListener('click', addItem, false);

var result1 = document.getElementById('firstsum');
var result2 = document.getElementById('lastsum'); 
