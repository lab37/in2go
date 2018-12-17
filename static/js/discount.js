function hetong_addItem() {
	// Create a new element and store it in a variable.
	var newTr = document.createElement('tr');
	var position = document.getElementById('hetong');
	position.appendChild(newTr);
	var newTd = document.createElement('td');
	var lastTr = position.lastChild;
	for (let i = 0; i < 6; i++) {
		newTd = document.createElement('td')
		newInput = document.createElement('input')
		newTd.appendChild(newInput);
		if (i > 5) {
			newInput.disabled = "true";
		}
		lastTr.appendChild(newTd);
	}

};



function hetong_delItem() {

	// Find the element which contains the element to be removed.
	var containerEl = document.getElementById('hetong');
	var removeEl = containerEl.lastChild;
	if (containerEl.children.length > 1) {
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
	for (let i = 0; i < 6; i++) {
		newTd = document.createElement('td')
		newInput = document.createElement('input')
		if (i > 5) {
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
	if (containerEl.children.length > 1) {
		containerEl.removeChild(removeEl);
	}
};

function hook(){
	progressLabel.text("计算中。。。。。");
	progressbar.progressbar("value", false);//清除进度条value的值，使其显示为条形
	
}

function doIt() {
	let prdtName = new Array();
	let contracts = new Array();
	let needSum = 0;
	var hetongItem = document.querySelectorAll('#hetong tr');
	let patchs = new Array;
	var patchItem = document.querySelectorAll('#patch tr');
	let tmpContractArr = new Array();
	let tmpPatchQuantityArr = new Array();
	let tmpPatchPriceArr = new Array();
	let resultAll = new Array();

	console.log("合同品种个数:" + (hetongItem.length - 1));

	let hetongCount = hetongItem.length - 1;
	for (let i = 0; i < hetongCount; i++) {
		let contract = new Array();

		needSum = needSum + Math.floor(hetongItem[i + 1].children[1].firstChild.value) * Math.floor(hetongItem[i + 1].children[2].firstChild.value) * Math.floor(hetongItem[i + 1].children[3].firstChild.value);

		contract.push(Math.floor(hetongItem[i + 1].children[1].firstChild.value));
		contract.push(Math.floor(hetongItem[i + 1].children[2].firstChild.value));
		contract.push(Math.floor(hetongItem[i + 1].children[4].firstChild.value));
		contract.push(Math.floor(hetongItem[i + 1].children[5].firstChild.value));
		tmpContractArr.push(Math.floor(hetongItem[i + 1].children[4].firstChild.value));
		prdtName.push(hetongItem[i + 1].children[0].firstChild.value + '单价');

		contracts.push(contract);

	}
	let patchCount = patchItem.length - 1;
	console.log("赠品品种个数:" + (patchItem.length - 1));
	for (let j = 0; j < patchCount; j++) {
		let patch = new Array();

		patch.push(Math.floor(patchItem[j + 1].children[1].firstChild.value));
		patch.push(Math.floor(patchItem[j + 1].children[2].firstChild.value));
		patch.push(Math.floor(patchItem[j + 1].children[3].firstChild.value));
		patch.push(Math.floor(patchItem[j + 1].children[4].firstChild.value));
		patch.push(Math.floor(patchItem[j + 1].children[5].firstChild.value));
		tmpPatchQuantityArr.push(Math.floor(patchItem[j + 1].children[1].firstChild.value));
		tmpPatchPriceArr.push(Math.floor(patchItem[j + 1].children[4].firstChild.value));
		prdtName.push(patchItem[j + 1].children[0].firstChild.value + '件数');
		patchs.push(patch);
	}

	for (let j = 0; j < patchCount; j++) {
		prdtName.push(patchItem[j + 1].children[0].firstChild.value + '单价');
	}


	// console.log(needSum);
	// console.log(contracts);
	// console.log(patchs);
	// console.log(tmpContractArr);
	// console.log(tmpPatchPriceArr);
	// console.log(tmpPatchQuantityArr);

	let lastSum = 0;
	while (tmpContractArr[0] <= contracts[0][3]) {

		let tmpSum = 0;
		for (let i = 0; i < tmpContractArr.length; i++) {//计算此时的合同总额
			tmpSum = tmpSum + contracts[i][0] * contracts[i][1] * tmpContractArr[i];
		}
		for (let j = 0; j < tmpPatchQuantityArr.length; j++) {//再加上此时的赠品总额
			tmpSum = tmpSum + patchs[j][2] * tmpPatchQuantityArr[j] * tmpPatchPriceArr[j];
		}

		if (Math.abs(tmpSum - needSum) < 300) {//如果结果符合要求,记录当前结果状态
			let t1 = tmpContractArr.concat(tmpPatchQuantityArr);
			let t2 = t1.concat(tmpPatchPriceArr);
			t2.push(tmpSum);
			resultAll.push(t2);
		}

		tmpPatchPriceArr[tmpPatchPriceArr.length - 1]++;
		for (let k = tmpPatchPriceArr.length - 1; k > 0; k--) {//处理赠品价格数组的进位
			if (tmpPatchPriceArr[k] == patchs[k][4]) {
				tmpPatchPriceArr[k - 1]++;
				tmpPatchPriceArr[k] = patchs[k][3];
			}
		}
		if (tmpPatchPriceArr[0] == patchs[0][4]) {//赠品价格数组进位到头后触发赠品数量数组增加
			tmpPatchQuantityArr[tmpPatchQuantityArr.length - 1]++;
			tmpPatchPriceArr[0] = patchs[0][3];
			lastSum = 0;//标记下次lastSum不计入判断
		} else {
			//赠品价格没有进位到头，但是若差值越远了的话后面的也不用测了因为是单调的，直接进位。
			if ((lastSum != 0) && (Math.abs(lastSum - needSum) < Math.abs(tmpSum - needSum)) && Math.abs(tmpSum - needSum) > 300) {
				tmpPatchQuantityArr[tmpPatchQuantityArr.length - 1]++;
				tmpPatchPriceArr[0] = patchs[0][3];
				lastSum = 0;//标记下次lastSum不计入判断
			} else {
				lastSum = tmpSum;//标记下次lastSum计入判断
			}
		}
		for (let k = tmpPatchQuantityArr.length - 1; k > 0; k--) {//处理赠品数量数组进位
			if (tmpPatchQuantityArr[k] == patchs[k][1]) {
				tmpPatchQuantityArr[k - 1]++;
				tmpPatchQuantityArr[k] = patchs[k][0];
			}
		}

		if (tmpPatchQuantityArr[0] == patchs[0][1]) {//赠品数量数组进位到头后触发合同价格数组增加
			tmpContractArr[tmpContractArr.length - 1]++;
			tmpPatchQuantityArr[0] = patchs[0][0];
		}
		for (let k = tmpContractArr.length - 1; k > 0; k--) {//处理合同价格数组进位
			if (tmpContractArr[k] == contracts[k][3]) {
				tmpContractArr[k - 1]++;
				tmpContractArr[k] = contracts[k][2];
			}
		}
		// let pct = Math.round((tmpContractArr[0] - contracts[0][2]) / (contracts[0][3] - contracts[0][2]) * 100);
		// progressbar.progressbar("value", pct);//设置value的值

	}
	
	

	var newContent = '<h3>可行方案：</h3><p>原合同总额：' + needSum + '</p><table>';
	newContent += '<tr class="table_title">';
	for (let i = 0; i < prdtName.length; i++) {
		newContent += '<td>' + prdtName[i] + '</td>';
	}
	newContent += '<td>总额</td>';
	newContent += '<td>差额</td>';
	newContent += '</tr>';
	for (let i = 0; i < resultAll.length; i++) {
		newContent += '<tr>';
		for (let j = 0; j < resultAll[i].length; j++) {
			newContent += '<td>' + resultAll[i][j] + '</td>';
		}

		newContent += '<td>' + (needSum - resultAll[i][resultAll[0].length - 1]) + '</td>';
		newContent += '</tr>';

	}
	newContent += '</table>';
	// Update the page with the new content
	document.getElementById('results').innerHTML = newContent;
	$("#results").show();
	progressbar.progressbar("value", 100);
	// progressLabel.text("计算完成！");
	// alert("计算完成");
}
// mFor(minArr, maxArr, Math.floor(h), Math.floor(firstSum), perQuanArr);


$('#hetong_addItem').on('click', hetong_addItem);
document.getElementById('hetong_delItem').addEventListener('click', hetong_delItem, false);

$('#patch_addItem').on('click', patch_addItem);
$('#patch_delItem').on('click', patch_delItem);
$('#doit').on('click', hook);


var result1 = document.getElementById('firstsum');
var result2 = document.getElementById('lastsum');
$("#results").hide();

var progressbar = $("#progressbar");
var progressLabel = $("#progress-label");

progressbar.progressbar({
	// value: 0,//进度条的初始值，value就是进度的值
	change: function () {//进度改变的触发函数
		if(progressbar.progressbar("value")!=100){
		setTimeout( doIt, 1000 );
	}
	},
	complete: function () {//时度完成后的触发函数
		progressLabel.text("计算完成！");
		alert("计算完成");
	}
});
// function progress() {//通过此函数不断的改变进度条的值
// 	let pct = Math.round((tmpContractArr[0] - contracts[0][2]) / (contracts[0][3] - contracts[0][2]) * 100);
// 	progressbar.progressbar("value", pct);//设置value的值
// 	setTimeout( progress, 3000 );

// }
