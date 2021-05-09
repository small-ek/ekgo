let timeout = null;
let timer, flag;
export default {
    /**
     * 删除对象
     * @param parameter
     * @returns {Array}
     */
    deleteObject: function (parameter) {
        var object = parameter['object']
        var value = parameter['value']
        var data = [];
        object.forEach(function (v, i) {
            if (v[parameter.name] != value) {
                data.push(v)
            }
        })
        return data
    },
    /**
     * 更具值删除json 数组数据
     * @param {Object} list 原有数组
     * @param {Object} name 删除数组名称
     * @param {Object} value 删除数组值
     */
    deleteArrayJson(list, name, value) {
        var newArr = new Array();
        for (var i = 0; i < list.length; i++) {
            var j = list[i];
            if (j[name] != value) {
                newArr.push(j);
            }
        }
        return newArr;
    },
    /**
     * 更具值删除数组json 数组数据
     * @param {Object} list 原有数组
     * @param {Object} name 删除数组名称
     * @param {Object} value 删除数组值
     */
    deleteArray(list, value) {
        for (var i = 0; i < list.length; i++) {
            if (list[i] == value) {
                list.splice(i, 1) //删除对应的值
            }
        }
        return list;
    },
    /**
     * 跳转
     * @param self
     * @param url
     * @param query
     */
    jump: function (self, url, query) {
        self.$router.push({
            path: url,
            query: query
        })
    },
    /**
     * 日期插件
     */
    dateOption: function () {
        return {
            shortcuts: [{
                text: '最近一周',
                onClick(picker) {
                    const end = new Date();
                    const start = new Date();
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
                    picker.$emit('pick', [start, end]);
                }
            }, {
                text: '最近一个月',
                onClick(picker) {
                    const end = new Date();
                    const start = new Date();
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
                    picker.$emit('pick', [start, end]);
                }
            }, {
                text: '最近三个月',
                onClick(picker) {
                    const end = new Date();
                    const start = new Date();
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
                    picker.$emit('pick', [start, end]);
                }
            }]
        }
    },
    /**
     * 菜单树形转换
     * @param data
     * @param key
     * @param pid
     * @returns {[]}
     */
    toTree(data, key, pid) {
        if (data.length > 0) {
            var key = key ? key : 'id'
            var pid = pid ? pid : 'parent_id'
            // 删除 所有 children,以防止多次调用
            data.forEach(function (item) {
                delete item.children;
            });
            // 将数据存储为 以 id 为 KEY 的 map 索引数据列
            var map = {};
            data.forEach(function (item) {
                map[item[key]] = item;
            });
            var val = [];
            data.forEach(function (item) {
                // 以当前遍历项，的pid,去map对象中找到索引的id
                var parent = map[item[pid]];
                // 好绕啊，如果找到索引，那么说明此项不在顶级当中,那么需要把此项添加到，他对应的父级中
                if (parent) {
                    (parent.children || (parent.children = [])).push(item);
                } else {
                    //如果没有在map中找到对应的索引ID,那么直接把 当前的item添加到 val结果集中，作为顶级
                    val.push(item);
                }
            });
            return val;
        }
    },
    /**
     * form表单是否换行
     */
    formPosition() {
        var width = document.body.clientWidth;
        if (width < 768) {
            return 'top';
        }
        return 'right'
    },
    /**
     * 页面搜索
     * @param {Object} self
     */
    search(self, name) {
        self.fullscreen_loading = true;
        setTimeout(() => {
            if (name == undefined) {
                self.$refs.page.search();
            } else {
                self.$refs[name].search();
            }
            self.fullscreen_loading = false
        }, 300);
    },
    /**
     * 数据分组
     * @param {Object} array
     * @param {Object} f
     */
    groupBy(list, name) {
        const groups = [];
        list.forEach(value => {
            const key = {};
            const data = {};

            let group = groups.find(value2 => {
                return value2.title === value[name];
            });
            if (!group) {
                group = {
                    title: value[name],
                };
                groups.push(group);
            }
            group.children = group.children || [];
            group.children.push(value);
        });
        return groups;

    },
    /**
     * 获取数组最大值
     * @param {Object} array
     */
    maxArray(array) {
        return Math.max.apply(Math, array)
    },
    /**
     * 获取数组最小值
     * @param {Object} array
     */
    minArray(array) {
        return Math.min.apply(Math, array)
    },
    /**
     * 搜索数组key合并一维数组
     * @param {Object} array
     */
    arrayColumn(array, key) {
        var result = [];
        for (var i = 0; i < array.length; i++) {
            if (array[i][key] || array[i][key] === 0) {
                result.push(array[i][key])
            }
        }
        return result;
    },
    /**
     * 笛卡尔积算法 一般用于库存生成规格
     * @param {Object} array [[1,2,3],[4,5,6]]
     */
    cartesian(array) {
        if (array.length < 2) return array[0] || [];
        return [].reduce.call(array, function (col, set) {
            var res = [];
            col.forEach(function (c) {
                set.forEach(function (s) {
                    var t = [].concat(Array.isArray(c) ? c : [c]);
                    t.push(s);
                    res.push(t);
                })
            });
            return res;
        });
    },
    /**
     * 格式化时间
     * @param {Object} date 时间
     * @param {Object} fmt  格式化类型
     */
    formatDate(date, fmt) {
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
    /**
     * 精确计算回调
     * @param f
     * @param digit
     * @returns {number}
     */
    formatFloat(f, digit = 2) {
        var m = Math.pow(10, digit);
        const result = Math.round(f * m, 10) / m

        return result;
    },
    /**
     * 删除筛选条件
     * @param self
     */
    clearSearch(where) {
        for (var key in where) {
            if (key != 'page_size' && key != 'current_page' && key != 'filter') {
                where[key] = ""
            }
        }
        if (where.filter != undefined) {
            var newWhere = JSON.parse(JSON.stringify(where.filter))
            newWhere.forEach(function (item) {
                item[2] = ""
            })
            where.filter = newWhere
        }
    },
    /**
     * 搜索开关
     * @param self
     */
    searchSwitch(self) {
        self.search_status = self.search_status == true ? false : true
    },
    //浮点数乘法运算
    numberRide(arg1, arg2) {
        var m = 0,
            s1 = arg1.toString(),
            s2 = arg2.toString();
        try {
            m += s1.split(".")[1].length;
        } catch (e) {
        }
        try {
            m += s2.split(".")[1].length;
        } catch (e) {
        }
        return (
            (Number(s1.replace(".", "")) * Number(s2.replace(".", ""))) /
            Math.pow(10, m)
        );
    },
    /**
     * 防抖
     * @param func
     * @param wait
     * @param immediate
     */
    debounce(func, wait = 500, immediate = false) {
        // 清除定时器
        if (timeout !== null) clearTimeout(timeout);
        // 立即执行，此类情况一般用不到
        if (immediate) {
            var callNow = !timeout;
            timeout = setTimeout(function () {
                timeout = null;
            }, wait);
            if (callNow) typeof func === 'function' && func();
        } else {
            // 设置定时器，当最后一次操作后，timeout不会再被清除，所以在延时wait毫秒后执行func回调方法
            timeout = setTimeout(function () {
                typeof func === 'function' && func();
            }, wait);
        }
    },
    /**
     * 节流
     * @param func
     * @param wait
     * @param immediate
     */
    throttle(func, wait = 500, immediate = true) {
        if (immediate) {
            if (!flag) {
                flag = true;
                // 如果是立即执行，则在wait毫秒内开始时执行
                typeof func === 'function' && func();
                timer = setTimeout(() => {
                    flag = false;
                }, wait);
            }
        } else {
            if (!flag) {
                flag = true
                // 如果是非立即执行，则在wait毫秒内的结束处执行
                timer = setTimeout(() => {
                    flag = false
                    typeof func === 'function' && func();
                }, wait);
            }
        }
    },
    /**
     * 两个数组合并去除重复
     * @param arr1
     * @param arr2
     * @returns {*}
     */
    uniqueArray(arr1, arr2) {
        var result = []
        for (let i = 0; i < arr1.length; i++) {
            if (arr2.indexOf(arr1[i]) == -1) {
                result.push(arr1[i])
            }
        }
        return result;
    },
    assign(result, list, callback, name) {
        if (result.length == list.length) {
            return Promise.all(result).then(res => {
                for (let i = 0; i < list.length; i++) {
                    var value = list[i]
                    if (name && res[i].data.data) {
                        value[name] = res[i].data.data
                    } else if (res[i] != undefined && res[i].data.data) {
                        value["with"] = res[i].data.data
                    }
                }
                callback(list)
            })
        }
    }
}

