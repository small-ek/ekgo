
import { onUnmounted } from "vue";

/**
 * 是否为生产环境
 * @returns {boolean}
 */
export const isNotProduction = () => {
  return process.env.NODE_ENV !== 'production'
}

/**
 * 超时处理回调
 * 
 * @param {number [ref]}
 * @param {callback}
 *
 * @returns {void}
 */
export const isTimeout = (number,callback) => {
  setTimeout(() => {
    callback();
  }, number.value * 1000);
  const Interval = setInterval(() => {
    number.value--;
  }, 1000);
  onUnmounted(() => {
    clearInterval(Interval);
  });
}
