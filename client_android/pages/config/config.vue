<template>
    <view class="content">
        <view class="input-group">
            <view class="input-row border">
                <text class="title">服务端IP：</text>
                <m-input class="m-input" type="text" clearable focus v-model="serverIp" placeholder="请输入IP"></m-input>
            </view>
        </view>
        <view class="btn-row">
            <button type="primary" class="primary larger-btn" :loading="loginBtnLoading"
                @tap="bindSetServer">保存</button>
            <button type="primary" class="primary larger-btn" :loading="loginBtnLoading"
                @tap="bindCancelServer">取消</button>
        </view>
    </view>
</template>

<script>
    import mInput from '../../components/m-input.vue'

    export default {
        components: {
            mInput
        },
        data() {
            return {
                serverIp: uni.getStorageSync('serverIp'),
                positionTop: 0,
                isDevtools: false,
                loginBtnLoading: false,
            }
        },
        methods: {
            initPosition() {
                
                this.positionTop = uni.getSystemInfoSync().windowHeight - 100;
            },
            bindSetServer() {
                
                
                
                
                
                
                uni.setStorageSync('serverIp', this.serverIp)
                
                
                
                
                
                
                
                
                
                
                
                
                

                
                
                
                uni.navigateBack()
            },
            bindCancelServer() {
                console.log("cancel server")
                
                
                
                uni.navigateBack()
            },
            onReady() {
                this.initPosition();
            }
        },
    }
</script>

<style>
    .login-type {
        display: flex;
        justify-content: center;
    }

    .login-type-btn {
        line-height: 30px;
        margin: 0px 15px;
    }

    .login-type-btn.act {
        color: #0FAEFF;
        border-bottom: solid 1px #0FAEFF;
    }

    .send-code-btn {
        width: 120px;
        text-align: center;
        background-color: #0FAEFF;
        color: #FFFFFF;
    }

    .action-row {
        display: flex;
        flex-direction: row;
        justify-content: center;
    }

    .action-row navigator {
        color: #007aff;
        padding: 0 10px;
    }

    .oauth-row {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-around;
        flex-wrap: wrap;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
    }

    .oauth-image {
        position: relative;
        width: 50px;
        height: 50px;
        border: 1px solid #dddddd;
        border-radius: 50px;
        background-color: #ffffff;
    }

    .oauth-image image {
        width: 30px;
        height: 30px;
        margin: 10px;
    }

    .oauth-image button {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        opacity: 0;
    }

    .captcha-view {
        line-height: 0;
        justify-content: center;
        align-items: center;
        display: flex;
        position: relative;
        background-color: #f3f3f3;
    }

    .btn-row {
        display: flex;
        justify-content: space-between;
        
        
    }

    
    .btn-row button {
        
        margin: 5px;
        
    }

    .larger-btn {
        padding: 5px 40px;
        
        font-size: 16px;
        
        
    }
</style>