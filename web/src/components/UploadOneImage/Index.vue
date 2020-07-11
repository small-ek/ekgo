<template>
    <div>
        <el-upload class="avatar-uploader" :http-request="uploadFile" :show-file-list="false" :data="data" :headers="headers" :before-upload="beforeUpload" v-loading="loading">
            <img v-if="image" :src="image" class="avatar" :style="{width:width+'px',height:height+'px',lineHeight:height+'px'}">
            <i v-else class="el-icon-plus avatar-uploader-icon" :style="{width:width+'px',height:height+'px',lineHeight:height+'px'}"></i>
        </el-upload>
    </div>

</template>

<script>
    import upload from "../../common/upload.js";

    export default {
        name: "Index",
        props: {
            //原有数据
            model: String,
            data: Array,
            //每张图片最小kb-未开发
            min: Number,
            //每张图片最大多少kb
            max: {
                type: Number,
                default: 500
            },
            width: {
                type: Number,
                default: 45
            },
            height: {
                type: Number,
                default: 45
            },
            //空间名称
            name: String,
        },
        data() {
            return {
                image: '',
                loading: false,
                headers: {},
            };
        },
        created() {
            this.image = this.model
        },
        methods: {
            //点击图片查看
            handlePreview(file) {
                window.open(file.url);
            },
            //上传之前拦截
            beforeUpload(file) {
                const type = file.type;
                var fileRules = "";
                fileRules = /\/(gif|jpg|jpeg|png|bmp|JPG|GIF|JPEG|PNG|BMP)$/.test(
                    type
                );
                if (!fileRules) {
                    this.$message.error('请上传正确图片格式');
                    return false;

                } else if (file.size > 1024 * this.max) {
                    this.$message.error(this.$t('upload_size') + this.max + "KB!");
                    return false;
                }

            },
            /**
             * 上传文件
             * @param file
             * @returns {Promise<void>}
             */
            async uploadFile(file) {
                this.loading = true;
                //执行同步上传
                const ImageUrl = await upload.uploadFile(file);
                this.loading = false;
                this.image = ImageUrl

            },
        }
    };
</script>
<style scoped>
    .avatar-uploader .el-upload {
        border: 1px dashed #d9d9d9;
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
    }

    .avatar-uploader .el-upload:hover {
        border-color: #409EFF;
    }

    .avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        text-align: center;
        border: 1px dashed #8c939d;
    }

    .avatar {
        display: block;
    }
</style>
