/**
 * 树操作
 * @type {{recursion: ((function(*, *, *): (string|*))|*)}}
 */
const tree = {
    /**
     * 递归
     * @param data 数据
     * @param key 筛选的key名称
     * @param pid 父节点名称
     * @returns {*[]}
     */
    recursion: function (list, key = "id", pid = "parent_id") {
        if (list.length > 0) {
            // 删除 所有 children,以防止多次调用
            for (let i = 0; i < list.length; i++) {
                delete list[i].children;
            }

            //将数据存储为 以 id 为 KEY 的 map 索引数据列
            var map = {};
            for (let i = 0; i < list.length; i++) {
                map[list[i][key]] = list[i];
            }

            var val = [];
            for (let i = 0; i < list.length; i++) {
                // 以当前遍历项，的pid,去map对象中找到索引的id
                var parent = map[list[i][pid]];
                // 如果找到索引，那么说明此项不在顶级当中,那么需要把此项添加到，对应的父级中
                if (parent) {
                    (parent.children || (parent.children = [])).push(list[i]);
                } else {
                    //如果没有在map中找到对应的索引ID,那么直接把 当前的item添加到 val结果集中，作为顶级
                    val.push(list[i]);
                }
            }
            return val;
        }
    },
    /**
     * 递归获取所有父节点
     * @param list
     * @param id
     * @returns {T[]|*[]}
     */
    getParentId: function (list, id, key = "id") {
        for (let i = 0; i < list.length; i++) {
            if (list[i][key] == id) {
                return [list[i]]
            }
            if (list[i].children) {
                let node = this.getParentId(list[i].children, id);
                if (node !== undefined) {
                    return node.concat(list[i])
                }
            }
        }
    }

}
export {
    tree,
}
