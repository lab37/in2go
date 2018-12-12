function tax(low, hight, lowP, diffP, allP) {
    var diffPrice = 0.0;
    var profit = 0.0;
    if (allP == "") {
        diffPrice = low * lowP + (hight - low) * diffP;
    } else if (diffP == "") {
        diffPrice = low * lowP + hight * allP;
    } else {
        diffPrice = (hight - low) * diffP + hight * allP;
    }
    profit = hight - diffPrice - low;

    return [diffPrice, profit];

}

$("#do").click(function () {
    let L = $("#lowP").val();
    let H = $("#highP").val();
    let A = $("#allP").val();
    if ($("#lowPrice").val() == "" || $("#highPrice").val() == "") {
        alert('必须输入低价和高价');
        return;
    } else if ((L + A + H) == (L > A ? (L > H ? L : H) : (A > H ? A : H))) {
        alert('必须输入两个扣率');
        return;
    } else if (L != "" && H != "" && A != "") {
        alert('只能输入两个扣率');
        return;
    } else {
       let rsts=tax($("#lowPrice").val(), $("#highPrice").val(), L / 100, H / 100, A / 100);

        $("#diffPrice").text(rsts[0].toFixed(4));
        $("#profic").text(rsts[1].toFixed(4));
    }
});