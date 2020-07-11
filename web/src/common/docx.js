import 'docxtemplater/build/docxtemplater.js'
import 'pizzip/dist/pizzip.js'
import 'pizzip/dist/pizzip-utils.js'
import {saveAs} from 'file-saver'

/**
 4. 导出docx
 5. @param { String } tempDocxPath 模板文件路径
 6. @param { Object } data 文件中传入的数据
 7. @param { String } fileName 导出文件名称
 */
export const exportDocx = (tempDocxPath, data, fileName) => {
    // 读取并获得模板文件的二进制内容
    PizZipUtils.getBinaryContent(tempDocxPath, (error, content) => {

        if (error) {
            throw error;
        }

        let zip = new PizZip(content);
        var doc = new window.docxtemplater().loadZip(zip)
        doc.setData(data);
        try {
            // render the document (replace all occurences of {first_name} by John, {last_name} by Doe, ...)
            doc.render();
        } catch (error) {
            let e = {
                message: error.message,
                name: error.name,
                stack: error.stack,
                properties: error.properties,
            };
            // The error thrown here contains additional information when logged with JSON.stringify (it contains a property object).
            throw error;
        }
        let out = doc.getZip().generate({
            type: "blob",
            mimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
        }); //Output the document using Data-URI
        saveAs(out, fileName);
    });
};
