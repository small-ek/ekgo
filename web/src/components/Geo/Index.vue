<template>
    <el-cascader :props="props" v-model="model" @change="getArray"></el-cascader>
</template>

<script>
    import {getGeo} from "../../common/api";

    export default {
        props: {
            model: Array,
        },
        data() {
            return {
                props: {
                    lazy: true,
                    lazyLoad(node, resolve) {
                        const pid = node.value ? node.value : 0
                        setTimeout(() => {
                            getGeo(pid).then(function (result) {
                                var list = []
                                result.data.data.forEach(function (item) {
                                    list.push({
                                        value: item.id,
                                        label: item.ext_name,
                                        leaf: item.deep > 2
                                    })
                                })
                                resolve(list);
                            })
                        }, 20);
                    }
                }
            };
        },
        methods: {
            async getArray(value) {
                this.$emit('update:model', value);
            }
        },

    }
</script>
