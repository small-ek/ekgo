import http from './http.js' //上传aws
import {config} from './config'; //配置
//上传脚本
export default {
    /**
     * 创建AWS文件夹
     * @param {Object} file
     */
    createFile(file) {
        var result = {}
        var name = new Date().getTime();
        var FileName = name + '.' + file.file.name.substr(file.file.name.lastIndexOf(".") + 1);

        return new Promise(function (resolve, reject) {
            http.get('/upload/oss?file_name=' + FileName)
                .then(function (response) {
                    result = {
                        FileName: FileName,
                        url: response.data.data.put
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
        var result = false;

        return new Promise(function (resolve, reject) {
            http.put(url, file.file)
                .then(function (response) {
                    result = true
                    resolve(result)
                })
        })
    },
    /**
     * 上传图片
     * @param {Object} file
     */
    async uploadFile(file) {
        const result = await this.createFile(file)
        if (result.url) {
            const Req = new XMLHttpRequest();
            // put 请求上传文件
            Req.open('PUT', result.url, false);

            // 上传
            Req.send(file.file);

            if (Req.status == 200) {
                return config.oss_url + result.FileName;
            }
        }

    }
}
