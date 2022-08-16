function num2HexArray(num) {
    if (num === void 0) {
        return [];
    }
    num = parseInt(num);
    if (num === 0) {
        return [0];
    }
    let str = num.toString(16);
    str.length % 2 && (str = '0' + str);
    const array = [];
    for (let i = 0, len = str.length; i < len; i += 2) {
        array.push(`0x${str.substr(i, 2)}`);
    }
    return array;
}


function isInRange(ip, segment) {
      ipBlocks = ip.split(".");
      binaryIp = (Number(ipBlocks[0]) << 24) | (Number(ipBlocks[1]) << 16) | (Number(ipBlocks[2]) << 8) | (Number(ipBlocks[3]))
      type = Number(segment.split("/")[1])
      mask = 0xFFFFFFFF << (32 - type)
      segmentIpBlocks = segment.split("/")[0].split(".")
      binarySegmentIp = (Number(segmentIpBlocks[0]) << 24) | (Number(segmentIpBlocks[1]) << 16) | (Number(segmentIpBlocks[2]) << 8) | (Number(segmentIpBlocks[3]))
      
      console.log("ipBlocks=", ipBlocks);
      console.log("mask=%s",num2HexArray(mask));
      console.log("binaryIp=%s",binaryIp.toString(16));
      console.log("binarySegmentIp=%s",binarySegmentIp.toString(16));
      return (binaryIp & mask) == (binarySegmentIp & mask);
}

console.log(999, isInRange('192.168.158.49', '192.168.158.1/24'));
