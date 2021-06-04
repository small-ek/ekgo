/**
 * 时间格式化
 * @type {{set: (function(*=, *=): (string|string))}}
 */
var format = {
    /**
     * 格式化时间
     * @param date 原始时间
     * @param fmt 格式化字符串 YYYY-mm-dd HH:MM:SS
     * @returns {string|*}
     */
    set: function (date, fmt) {
        if (date == "" || date == undefined || date == null || date == '0001-01-01T00:00:00Z') {
            return ""
        } else {
            var GMT = new Date(date)
        }

        fmt == undefined ? fmt = "YYYY-mm-dd HH:MM:SS" : fmt;
        const opt = {
            "Y+": GMT.getFullYear().toString(), // 年
            "m+": (GMT.getMonth() + 1).toString(), // 月
            "d+": GMT.getDate().toString(), // 日
            "H+": GMT.getHours().toString(), // 时
            "M+": GMT.getMinutes().toString(), // 分
            "S+": GMT.getSeconds().toString() // 秒
            // 有其他格式化字符需求可以继续添加，必须转化成字符串
        };

        let ret;

        for (let k in opt) {
            ret = new RegExp("(" + k + ")").exec(fmt);
            if (ret) {
                fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
            }
        }
        return fmt;
    },

}
export {
    format,
}
