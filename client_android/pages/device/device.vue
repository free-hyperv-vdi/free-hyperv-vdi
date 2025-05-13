<template>
    <view class="container">
        <view class="table-body">
            <view v-for="(item, index) in vms" :key="index" class="table-row">
                <view class="record-box">
                    <view class="item-info">
                        <label>云桌面: </label>
                        <text>{{ item.name }}</text>
                    </view>
                    <view class="item-info">
                        <label>状态: </label>
                        <text>{{ item.status }}</text>
                    </view>
                    <view class="item-info">
                        <label>IP: </label>
                        <text>{{ item.virtualIp }}</text>
                    </view>

                    <view class="button-group">
                        <uni-button @click="start(item)">开机</uni-button>
                        <uni-button @click="shutdown(item)">关机</uni-button>
                        <uni-button @click="restart(item)">重启</uni-button>
                        <uni-button @click="connect(item)">连接</uni-button>
                    </view>
                </view>
            </view>
        </view>
    </view>
</template>


<script>
    export default {
        data() {
            return {
                vms: [],
            };
        },
        onLoad() {
            console.log("load vm data")
            console.log(uni.getStorageSync("cookie"))
            this.getVmList()
            this.startTimer()
        },
        onUnload() {
            clearInterval(this.timer)
        },
        methods: {
            startTimer() {
                this.timer = setInterval(() => {
                    
                    this.getVmList()
                }, 5000); 
            },
            getVmList() {
                
                var serverIp = uni.getStorageSync("serverIp")
                var targetUrl = 'http:
                console.log("targetUrl:", targetUrl)
                const self = this
                uni.request({
                    url: targetUrl, 
                    method: 'GET',
                    mode: 'cors',
                    header: {
                        'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',
                        'Cookie': uni.getStorageSync("cookie") 
                    },
                    success: function(res) {
                        console.log('GET 请求成功', res.data.data.devices);
                        let devices = res.data.data.devices
                        self.vms = devices.filter(item => item.username === uni.getStorageSync("username"))
                            .map(item => {
                                let tempItem = {
                                    name: item.name,
                                    status: item.status,
                                    virtualIp: item.virtualIp,
                                    id: item.id
                                }
                                return tempItem
                            })
                        console.log(self.vms)
                        
                    },
                    fail: function(err) {
                        console.error('GET 请求失败', err);
                        
                    }
                });

            },
            operateVm(vmId, action) {
                var serverIp = uni.getStorageSync("serverIp")
                var targetUrl = 'http:
                const self = this
                uni.request({
                    url: targetUrl,
                    method: 'POST',
                    header: {
                        'content-type': 'application/x-www-form-urlencoded' 
                    },
                    data: {
                        vm_id: vmId,
                        action: action
                    },
                    success: function(res) {
                        console.log('POST 请求成功', res)
                        if (res.data.code == 0) {
                            self.getVmList()
                        }
                    },
                    fail: function(err) {
                        console.error('POST 请求失败', err);
                        
                    }
                });
            },
            start(item) {
                this.operateVm(item.id, '1')
            },
            shutdown(item) {
                this.operateVm(item.id, '2')
            },
            restart(item) {
                this.operateVm(item.id, '3')
            },
            connect(desktop) {
                
                console.log('Connect to desktop:', desktop);
                const testPlugin = uni.requireNativePlugin('callrdc')
                console.log("sssss1")
                testPlugin.CallRdcStart(desktop.virtualIp, uni.getStorageSync("username"), (res) => {
                    console.log('res', res)
                    console.log(res.msg)
                })
            }
        },
    };
</script>

<style>
    .container {
        padding: 20px;
    }

    .table-body {
        display: flex;
        flex-direction: column;
    }

    .table-row {
        margin-bottom: 20px;
    }

    .record-box {
        border: 2px solid #0f0;
        padding: 10px;
        border-radius: 5px;
    }

    .button-group {
        display: flex;
    }

    .button-group uni-button {
        margin-right: 10px;
    }

    .item-info {
        margin-bottom: 10px;
        
    }

    label {
        font-weight: bold;
        margin-right: 5px;
    }
</style>