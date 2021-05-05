import CryptoJS from 'crypto-js/crypto-js'

const key = "142857ABCDEF";
const iv = "ABCDEF142857";

/**
 * AES加密
 * @param word
 * @returns {string}
 * @constructor
 */
export function AesEncrypt(text) {
    return CryptoJS.AES.encrypt(text, CryptoJS.enc.Utf8.parse(key), {
        iv: CryptoJS.enc.Utf8.parse(iv),
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
    }).toString()
}

/**
 * AES解密
 * @param word
 * @returns {string}
 * @constructor
 */
export function AesDecrypt(text) {
    let decrypted = CryptoJS.AES.decrypt(text, CryptoJS.enc.Utf8.parse(key), {
        iv: CryptoJS.enc.Utf8.parse(iv),
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
    })
    return decrypted.toString(CryptoJS.enc.Utf8)
}

/**
 * 哈希编码
 * @param str
 * @returns {*}
 * @constructor
 */
export function Sha256(str) {
    return CryptoJS.SHA256(str).toString()
}

/**
 * base64编码
 * @param str
 * @returns {*}
 * @constructor
 */
export function Base64Encode(str) {
    return CryptoJS.enc.Base64.stringify(CryptoJS.enc.Utf8.parse(str))
}

/**
 * base64解码
 * @param str
 * @returns {*}
 * @constructor
 */
export function Base64Decode(str) {
    var parsedWordArray = CryptoJS.enc.Base64.parse(str);
    return parsedWordArray.toString(CryptoJS.enc.Utf8);
}

//加密编码算法
export default {
    AesDecrypt,
    AesEncrypt,
    Sha256,
    Base64Encode,
    Base64Decode
}