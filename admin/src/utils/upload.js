import http from '../api/http.js' //上传aws
import config from "../config/index.js"
import {Sha256} from "./secret";
//上传脚本
export default {
    /**
     * 创建AWS文件夹
     * @param {Object} file
     */
    createFile(file, FileBase64, dir = "") {
        var result = {}
        var name = Sha256(FileBase64)
        var FileName = dir + name + "." + file.name.substr(file.name.lastIndexOf(".") + 1);

        return new Promise(function (resolve) {
            http.get('/upload/oss?file_name=' + FileName)
                .then(function (response) {
                    result = {
                        FileName: FileName,
                        url: response.data.put
                    }
                    resolve(result)
                })
        })
    },
    /**
     * 修改图片的文件
     * @param {Object} data
     */
    updateFile(url, file) {
        return new Promise(function (resolve) {
            http.put(url, file.file)
                .then(function (response) {
                    resolve(response)
                })
        })
    },
    /**
     * 文件转base64
     * @param {Object} file
     */
    getFileToBase64(file) {
        let reader = new FileReader();
        reader.readAsDataURL(file);

        let result = new Promise((resolve, reject) => {
            reader.onload = function () {
                resolve(reader.result)
            };
            reader.onerror = function (error) {
                reject(error);
            };
            reader.onloadend = function () {
                resolve(file.name + new Date().getTime());
            };
        });
        return result
    },
    /**
     * 上传图片
     * @param {Object} file
     */
    async uploadFile(file, dir) {
        const fileBase64 = await this.getFileToBase64(file.file);

        const result = await this.createFile(file.file, fileBase64, dir)
        if (result.url) {
            const Req = new XMLHttpRequest();
            // put 请求上传文件
            Req.open('PUT', result.url, false);

            // 上传
            Req.send(file.file);

            if (Req.status == 200) {
                return config.oss + result.FileName;
            }
        }

    }
}
