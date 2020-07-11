
```
#
content=传参内容
return=返回方法内容
#
<template>
    <div class="form">
        <el-form-item label="富文本" prop="describe">
            <Editor content="测试数据" @return="Editor"></Editor>
        </el-form-item>
    </div>
</template>
<script>
    import Editor from '@/components/Editor/Index.vue'

    export default {
        data() {
            return {
                data: [],
            };
        },
        components: {Editor},
        created: function () {
        },
        methods: {
            Editor(data) {
                console.log(data)
            }
        }
    }
</script>

```
