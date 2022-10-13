function encode(str) {
    let temp = str[0];
    let num = 1;
    let res = "";

    for (let i = 1; i < str.length; i++) {
        if (str[i] === temp) {
            num++;
            continue;
        }
        res += num + temp;
        temp = line[i];
        num = 1;
    }
    res += num + temp;
    return res;
}

let line = "aabbbceddddd";
console.log(encode(line));
