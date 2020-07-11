<template>
    <div>
        <span class="nav-right nav-icon" @click="onRight">
            <i class="el-icon-arrow-left"></i>
        </span>
        <span class="nav-left nav-icon" @click="onLeft">
            <i class="el-icon-arrow-right"></i>
        </span>
        <div class="scroll-container" ref="scrollContainer" @wheel.prevent="handleScroll">
            <div class="scroll-wrapper" ref="scrollWrapper" :style="{left: left + 'px'}">
                <slot></slot>
            </div>
        </div>
    </div>

</template>

<script>
    const padding = 0; // tag's padding
    export default {
        name: 'scrollPane',
        data() {
            return {
                left: 0,
            }
        },
        methods: {
            handleScroll(e) {
                const eventDelta = e.wheelDelta || -e.deltaY * 3
                const $container = this.$refs.scrollContainer
                const $containerWidth = $container.offsetWidth
                const $wrapper = this.$refs.scrollWrapper
                const $wrapperWidth = $wrapper.offsetWidth
                if (eventDelta > 0) {
                    this.left = Math.min(0, this.left + eventDelta)
                } else {
                    if ($containerWidth - padding < $wrapperWidth) {
                        if (this.left < -($wrapperWidth - $containerWidth + padding)) {
                            this.left = this.left
                        } else {
                            this.left = Math.max(this.left + eventDelta, $containerWidth - $wrapperWidth - padding)
                        }
                    } else {
                        this.left = 0
                    }
                }
            },
            moveToTarget($target) {
                const $container = this.$refs.scrollContainer
                const $containerWidth = $container.offsetWidth
                const $targetLeft = $target.offsetLeft
                const $targetWidth = $target.offsetWidth
                if ($targetLeft < -this.left) {
                    // tag in the left
                    this.left = -$targetLeft + padding
                } else if ($targetLeft + padding > -this.left && $targetLeft + $targetWidth < -this.left + $containerWidth -
                    padding) {
                    // tag in the current view
                    // eslint-disable-line
                } else {
                    // tag in the right
                    this.left = -($targetLeft - ($containerWidth - $targetWidth) + padding)
                }
            },
            onRight() {
                if (this.left < 0) {
                    this.left = (this.left + 300) < 0 ? this.left + 300 : 0
                }
            },
            onLeft() {
                const $container = this.$refs.scrollContainer
                const $containerWidth = $container.offsetWidth

                const $wrapper = this.$refs.scrollWrapper
                const $wrapperWidth = $wrapper.offsetWidth
                const right = $containerWidth - $wrapperWidth - padding
                if (this.left > right) {
                    this.left = (this.left - 300) > right ? this.left - 300 : right
                }
            }
        }
    }
</script>

<style lang="less" scoped>
    .nav-icon {
        line-height: 44px;
        cursor: pointer;
        color: #909399;
        height: 44px;
        font-weight: bold;
        width: 30px;
        text-align: center;
    }

    .nav-right {
        position: absolute;
        left: 0px;
    }

    .nav-left {
        position: absolute;
        right: 0px;
    }

    .scroll-container {
        margin: 0px 25px;
        cursor: ew-resize;
        white-space: nowrap;
        position: relative;
        overflow: hidden;

        .scroll-wrapper {
            padding: 0 4px;
            position: absolute;
            transition: left 0.3s ease-in-out;
            padding-right: 110px;
        }
    }
</style>
