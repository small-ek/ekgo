<template>
    <div>
        <el-upload class="upload-demo" multiple :http-request="uploadFile" :file-list="newFileList" :before-upload="beforeUpload" :data="data" :headers="headers" :limit="limit" list-type="picture-card" v-loading="loading">
            <i slot="default" class="el-icon-plus"></i>
            <div slot="tip" class="el-upload__tip">{{tip}}</div>

            <div slot="file" slot-scope="{file}" v-if=file.url>
                <video v-if="isType(file)==1" :src="file.url" controls="controls" width="145" height="146">
                    your browser does not support the video tag
                </video>
                <img v-else class="el-upload-list__item-thumbnail" :src="file.url" alt="">
                <span class="el-upload-list__item-actions">
				<span class="el-upload-list__item-preview" @click="handlePreview(file)">
					<i class="el-icon-zoom-in"></i>
				</span>
				<span v-if="!disabled" class="el-upload-list__item-delete" @click="handleRemove(file)">
					<i class="el-icon-delete"></i>
				</span>
			</span>
            </div>
        </el-upload>
        <el-dialog width="60%" :visible.sync="dialog" @close="dialogClose">
            <template v-if="file_type==2">
                <video :src="image_url" controls="controls">
                    您的浏览器不支持视频
                </video>
            </template>
            <template v-else>
                <img width="100%" :src="image_url" alt="">
            </template>
        </el-dialog>
    </div>
</template>

<script>
    import upload from "../../common/upload.js";

    export default {
        name: "Index",
        props: {
            //限制数量
            limit: {
                type: Number,
                default: 5
            },
            //原有数据
            fileList: Array,
            data: Array,
            width: {
                type: Number,
                default: 40
            },
            height: {
                type: Number,
                default: 40
            },
            rules: {
                type: String,
                default: "image"
            },
            //每张图片最小kb-未开发
            min: Number,
            //每张图片最大多少kb
            max: {
                type: Number,
                default: 500
            },
            //空间名称
            name: String,
        },
        data() {
            return {
                dialog: false,
                image_url: '',
                file_type: 1,
                newFileList: [],
                loading: false,
                headers: {},
                tip: "",
                action: ""
            };
        },

        created: function () {
            switch (this.rules) {
                case "image":
                    this.tip =
                        this.$t('image_tip') +
                        this.limit +
                        this.$t('upload_tip') +
                        this.max +
                        "KB";
                    break;
                case "video":
                    this.tip =
                        this.$t('video_tip') +
                        this.limit +
                        this.$t('upload_tip') +
                        this.max +
                        "KB";
                    break;
                case "all":
                    this.tip =
                        this.$t('all_tip') +
                        this.limit +
                        this.$t('upload_tip') +
                        this.max +
                        "KB";
                    break;
                default:
                    this.tip =
                        this.$t('all_tip') +
                        this.limit +
                        this.$t('upload_tip') +
                        this.max +
                        "KB";
                    break;
            }
            //解析原有图片地址
            this.parsingFileList();
        },
        watch: {
            /**
             * 监听图片变化
             * @param {Object} value
             * @param {Object} oldValue
             */
            fileList(value, oldValue) {
                if (value.length != oldValue.length) {
                    this.parsingFileList();
                }
            }
        },
        methods: {
            parsingFileList(data) {
                const self = this;
                if (typeof this.fileList == "object") {
                    this.fileList.forEach(function (item) {
                        self.newFileList.push({
                            name: item,
                            url: item
                        });
                    });
                }
            },
            //判断图片或者视频
            isType(file) {
                if (file.url.lastIndexOf('.mp4') != -1 || file.name.lastIndexOf('.mp4') != -1) {
                    return 1;
                } else {
                    return 2;
                }

            },
            //删除图片
            handleRemove(file) {
                var fileList = this.fileList;
                //新的原有数据
                this.newFileList = this.$Func.deleteArrayJson(
                    this.newFileList,
                    "url",
                    file.url
                );
                //原有数据删除
                this.fileList = this.$Func.deleteArray(this.fileList, file.url);

            },
            //点击图片查看
            handlePreview(file) {
                var fileRules = /\.(gif|jpg|jpeg|png|bmp|JPG|GIF|JPEG|PNG|BMP)$/.test(
                    file.url
                );
                if (fileRules) {
                    this.file_type = 1
                } else {
                    this.file_type = 2
                }
                this.image_url = file.url;
                this.dialog = true;
            },
            dialogClose() {
                this.image_url = '';
            },
            //上传之前拦截
            beforeUpload(file) {
                const type = file.type;
                var fileRules = "";
                var tip = "";

                switch (this.rules) {
                    case "image":
                        fileRules = /\/(gif|jpg|jpeg|png|bmp|JPG|GIF|JPEG|PNG|BMP)$/.test(
                            type
                        );
                        tip = this.$t('upload_limit_tip') + "jpeg、jpg、png" + this.$t('files') + '!';
                        break;
                    case "video":
                        fileRules = /\/(mp4)$/.test(type);
                        tip = this.$t('upload_limit_tip') + "mp4" + this.$t('files') + '!';
                        break;
                    case "all":
                        fileRules = /\/(gif|jpg|jpeg|png|bmp|mp4|JPG|GIF|JPEG|PNG|BMP)$/.test(
                            type
                        );
                        tip = this.$t('upload_limit_tip') + "jpeg、jpg、png、mp4" + this.$t('files') + '!';
                        break;
                    default:
                        fileRules = /\/(gif|jpg|jpeg|png|bmp|mp4|JPG|GIF|JPEG|PNG|BMP)$/.test(
                            type
                        );
                        tip = this.$t('upload_limit_tip') + "jpeg、jpg、png、mp4" + this.$t('files') + '!';
                        break;
                }

                if (!fileRules) {

                    this.$message.error(tip);
                    return false;

                } else if (file.size > 1024 * this.max) {

                    this.$message.error(this.$t('upload_size') + this.max + "KB!");
                    return false;

                }

            },
            async uploadFile(file) {
                this.loading = true;
                //执行同步上传
                const ImageUrl = await upload.uploadFile(file);
                this.loading = false;
                this.newFileList.push({
                    name: ImageUrl,
                    url: ImageUrl
                });
                this.fileList.push(ImageUrl);
            }
        }
    };
</script>
