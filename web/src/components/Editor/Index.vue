<template>
    <div>
        <el-upload
                class="avatar-uploader-image"
                multiple
                :http-request="uploadFile"
                :before-upload="beforeUploadImage"
                :on-success="onSuccessImage"
                :data="data"
                style="display: none"
                list-type="picture">
        </el-upload>
        <quill-editor ref="myTextEditor"
                      class="ql-editor"
                      v-model="detailContent"
                      :options="editorOption"
                      @blur="onEditorBlur($event)"
                      @focus="onEditorFocus($event)"
                      @change="onEditorChange($event)">
        </quill-editor>
        <p class="tip" :style="tipColor">{{tip}}</p>
    </div>
</template>

<script>
    import {quillEditor} from 'vue-quill-editor'
    import 'quill/dist/quill.core.css'
    import 'quill/dist/quill.snow.css'
    import 'quill/dist/quill.bubble.css'
    import upload from "../../common/upload.js";
    import * as Quill from 'quill' //引入编辑器       '
    //quill图片可拖拽上传
    import {ImageDrop} from 'quill-image-drop-module';

    Quill.register('modules/imageDrop', ImageDrop);
    // quill编辑器的字体
    var fonts = ['SimSun', 'SimHei', 'Microsoft-YaHei', 'KaiTi', 'FangSong', 'Arial', 'Times-New-Roman', 'sans-serif'];
    var Font = Quill.import('formats/font');
    Font.whitelist = fonts; //将字体加入到白名单
    Quill.register(Font, true);

    // 工具栏配置
    const toolbarOptions = [
        ["bold", "italic", "underline", "strike"], // 加粗 斜体 下划线 删除线
        ["blockquote", "code-block"], // 引用  代码块
        [{header: 1}, {header: 2}], // 1、2 级标题
        [{list: "ordered"}, {list: "bullet"}], // 有序、无序列表
        [{script: "sub"}, {script: "super"}], // 上标/下标
        [{indent: "-1"}, {indent: "+1"}], // 缩进
        [{'direction': 'rtl'}],            // 文本方向
        [{size: ["small", false, "large", "huge"]}], // 字体大小
        [{header: [1, 2, 3, 4, 5, 6, false]}], // 标题
        [{color: []}, {background: []}], // 字体颜色、字体背景颜色
        [{font: fonts}], // 字体种类
        [{align: []}], // 对齐方式
        ["clean"], // 清除文本格式
        ["link", "image"] // 链接、图片、视频video
    ];
    export default {
        components: {
            quillEditor
        },
        props: {
            content: String,
            name: String,
            self: Object,
        },
        data() {
            return {
                tip: '当前已输入0个字符, 您还可以输入10000个字符。',
                max: 5000,
                tipColor: 'color:#ccc',
                name: 'Index',
                imageSize: 5,
                videoSize: 200,
                data: {},
                detailContent: '', // 富文本内容
                editorOption: {
                    placeholder: '请输入内容',
                    modules: {
                        imageDrop: true,
                        toolbar: {
                            container: toolbarOptions,  // 工具栏
                            handlers: {
                                'image': function (value) {
                                    if (value) {
                                        document.querySelector('.avatar-uploader-image input').click()
                                    } else {
                                        this.quill.format('image', false);
                                    }
                                },
                            }
                        },
                    }
                }
            }
        },
        created: function () {
            this.detailContent = this.content
            var editorLength = this.content.replace(/<[^>]+>/g, "").length
            var text = this.max - editorLength
            this.tip = "当前已输入" + editorLength + "个字符, 您还可以输入" + text + "个字符。";
        },
        methods: {
            /**
             * 复制给父组件model
             */
            setParentData() {
                const self = this.self;
                var name = this.name.split('.')
                switch (name.length) {
                    case 1:
                        self[name[0]] = this.detailContent
                        break
                    case 2:
                        self[name[0]][name[1]] = this.detailContent
                        break
                    case 3:
                        self[name[0]][name[1]][name[2]] = this.detailContent
                        break
                    case 4:
                        self[name[0]][name[1]][name[2]][name[3]] = this.detailContent
                        break
                }
            },
            onEditorBlur(editor) {
                this.setParentData()
            },
            onEditorFocus(editor) {
                this.setParentData()
            },
            onEditorChange(editor) {
                var editorLength = this.detailContent.replace(/<[^>]+>/g, "").length
                var text = this.max - editorLength
                if (this.max < editorLength) {
                    this.tipColor = "color:red"
                    this.tip = "字数超出最大允许值，服务器可能拒绝保存！";
                } else {
                    this.tipColor = "color:#ccc"
                    this.tip = "当前已输入" + editorLength + "个字符, 您还可以输入" + text + "个字符。";
                }
                this.setParentData()
            },
            beforeUploadImage(file) {//图片上传之前
                const type = file.type;
                var fileRules = /\/(gif|jpg|jpeg|png|bmp)$/.test(type);
                var tip = '只能上传gif/jpeg/jpg/png文件!';
                if (!fileRules) {
                    this.$message.error(tip);
                    return false
                } else if (file.size > (1024 * 1024 * this.imageSize)) {
                    this.$message.error('上传文件大小不能超过 ' + this.imageSize + 'MB!');
                    return false
                }
                this.data['key'] = file.lastModified + '_' + file.name
            },
            async uploadFile(file) {//图片上传成功
                //执行同步上传
                const ImageUrl = await upload.uploadFile(file);
                this.detailContent = this.detailContent + '<img src="' + ImageUrl + '"/>'
            },
        },
        computed: {
            editor() {
                return this.$refs.myTextEditor.quill
            },
        },
        watch: {
            content(value) {//监听值的变化
                this.detailContent = value
            }
        }
    }
</script>
<style>
    .quill-editor:not(.bubble) .ql-container,
    .quill-editor:not(.bubble) .ql-container .ql-editor {
        max-height: 30rem;
        /*height: 80px;*/
        padding-bottom: 1rem;
    }

    .ql-snow .ql-picker-label {
        top: -2px;
        height: 26px;
        line-height: 26px;
    }

    .tip {
        color: #ccc;
        height: 25px;
        float: right;
        line-height: 25px;
        font-size: 12px;
    }

    .editor {
        line-height: normal !important;
        height: 800px;
    }

    .ql-snow .ql-tooltip[data-mode=link]::before {
        content: "请输入链接地址:";
    }

    .ql-snow .ql-tooltip.ql-editing a.ql-action::after {
        border-right: 0px;
        content: '保存';
        padding-right: 0px;
    }

    .ql-snow .ql-tooltip[data-mode=video]::before {
        content: "请输入视频地址:";
    }

    .ql-snow .ql-picker.ql-size .ql-picker-label::before,
    .ql-snow .ql-picker.ql-size .ql-picker-item::before {
        content: '14px';
    }

    .ql-snow .ql-picker.ql-size .ql-picker-label[data-value=small]::before,
    .ql-snow .ql-picker.ql-size .ql-picker-item[data-value=small]::before {
        content: '10px';
    }

    .ql-snow .ql-picker.ql-size .ql-picker-label[data-value=large]::before,
    .ql-snow .ql-picker.ql-size .ql-picker-item[data-value=large]::before {
        content: '18px';
    }

    .ql-snow .ql-picker.ql-size .ql-picker-label[data-value=huge]::before,
    .ql-snow .ql-picker.ql-size .ql-picker-item[data-value=huge]::before {
        content: '32px';
    }

    .ql-snow .ql-picker.ql-header .ql-picker-label::before,
    .ql-snow .ql-picker.ql-header .ql-picker-item::before {
        content: '文本';
    }

    .ql-snow .ql-picker.ql-header .ql-picker-label[data-value="1"]::before,
    .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="1"]::before {
        content: '标题1';
    }

    .ql-snow .ql-picker.ql-header .ql-picker-label[data-value="2"]::before,
    .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="2"]::before {
        content: '标题2';
    }

    .ql-snow .ql-picker.ql-header .ql-picker-label[data-value="3"]::before,
    .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="3"]::before {
        content: '标题3';
    }

    .ql-snow .ql-picker.ql-header .ql-picker-label[data-value="4"]::before,
    .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="4"]::before {
        content: '标题4';
    }

    .ql-snow .ql-picker.ql-header .ql-picker-label[data-value="5"]::before,
    .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="5"]::before {
        content: '标题5';
    }

    .ql-snow .ql-picker.ql-header .ql-picker-label[data-value="6"]::before,
    .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="6"]::before {
        content: '标题6';
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item::before {
        content: '标准字体';
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=serif]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=serif]::before {
        content: '衬线字体';
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=monospace]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=monospace]::before {
        content: '等宽字体';
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=SimSun]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=SimSun]::before {
        content: "宋体";
        font-family: "SimSun";
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=SimHei]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=SimHei]::before {
        content: "黑体";
        font-family: "SimHei";
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=Microsoft-YaHei]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=Microsoft-YaHei]::before {
        content: "微软雅黑";
        font-family: "Microsoft YaHei";
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=KaiTi]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=KaiTi]::before {
        content: "楷体";
        font-family: "KaiTi";
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=FangSong]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=FangSong]::before {
        content: "仿宋";
        font-family: "FangSong";
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=Arial]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=Arial]::before {
        content: "Arial";
        font-family: "Arial";
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=Times-New-Roman]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=Times-New-Roman]::before {
        content: "Times New Roman";
        font-family: "Times New Roman";
    }

    .ql-snow .ql-picker.ql-font .ql-picker-label[data-value=sans-serif]::before,
    .ql-snow .ql-picker.ql-font .ql-picker-item[data-value=sans-serif]::before {
        content: "sans-serif";
        font-family: "sans-serif";
    }

    .ql-font-SimSun {
        font-family: "SimSun";
    }

    .ql-font-SimHei {
        font-family: "SimHei";
    }

    .ql-font-Microsoft-YaHei {
        font-family: "Microsoft YaHei";
    }

    .ql-font-KaiTi {
        font-family: "KaiTi";
    }

    .ql-font-FangSong {
        font-family: "FangSong";
    }

    .ql-font-Arial {
        font-family: "Arial";
    }

    .ql-font-Times-New-Roman {
        font-family: "Times New Roman";
    }

    .ql-font-sans-serif {
        font-family: "sans-serif";
    }
</style>
