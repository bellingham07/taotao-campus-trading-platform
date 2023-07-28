import { defineStore } from "pinia";

let userStore = defineStore("user", {
    // 存储数据:state
    state: () => {
        return {
            count: 99,
            arr: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
        }
    },
    actions: {

        updateNum(a: number, b: number) {
            this.count += a;
        }
    },
    getters: {
        total() {
            let result:any = this.arr.reduce((prev: number, next: number) => {
                return prev + next;
            }, 0);
            return result;
        }
    }
});

export default userStore;