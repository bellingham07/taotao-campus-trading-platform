<template>
    <van-form @failed="onFailed">
        <van-cell-group inset>
            <!-- 通过 pattern 进行正则校验 -->
            <van-field
                    v-model="value1"
                    name="pattern"
                    placeholder="正则校验"
                    :rules="[{ pattern, message: '请输入正确内容' }]"
            />
            <!-- 通过 validator 进行函数校验 -->
            <van-field
                    v-model="value2"
                    name="validator"
                    placeholder="函数校验"
                    :rules="[{ validator, message: '请输入正确内容' }]"
            />
        </van-cell-group>
        <div style="margin: 16px;">
            <van-button round block type="primary" native-type="submit">
                提交
            </van-button>
        </div>
    </van-form>


</template>

<script setup lang="ts">

</script>


<style scoped>

</style>