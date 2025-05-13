<script>
    import {
        mapState,
        mapMutations
    } from 'vuex'

    export default {
        onLaunch: function() {
            plus.screen.lockOrientation("landscape");
            console.log('App Launch');
            let uniIdToken = uni.getStorageSync('uni_id_token')
            if (uniIdToken) {
                this.login(uni.getStorageSync('username'))
            }

            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
            
        },
        onShow: function() {
            console.log('App Show');
        },
        onHide: function() {
            console.log('App Hide');
        },
        methods: {
            ...mapMutations(['login', 'setUniverifyErrorMsg', 'setHideUniverify']),
        }
    }
</script>

<style>
    
    @import "components/m-icon/m-icon.css";

    
    page {
        min-height: 100%;
        display: flex;
        font-size: 14px;
    }

    input,
    textarea,
    button {
        font-size: 14px;
    }

    
    page {
        width: 100%;
        height: 100%;
        display: block;
    }

    swan-template {
        width: 100%;
        min-height: 100%;
        display: flex;
    }

    
    custom-component {
        width: 100%;
        min-height: 100%;
        display: flex;
    }

    

    
    page {
        min-height: 100vh;
    }

    

    
    m-input {
        width: 100%;
        
        display: flex;
        flex: 1;
    }

    .content {
        display: flex;
        flex: 1;
        flex-direction: column;
        background-color: #efeff4;
        padding: 10px;
    }

    .input-group {
        background-color: #ffffff;
        margin-top: 20px;
        position: relative;
    }

    .input-group::before {
        position: absolute;
        right: 0;
        top: 0;
        left: 0;
        height: 1px;
        content: '';
        -webkit-transform: scaleY(.5);
        transform: scaleY(.5);
        background-color: #c8c7cc;
    }

    .input-group::after {
        position: absolute;
        right: 0;
        bottom: 0;
        left: 0;
        height: 1px;
        content: '';
        -webkit-transform: scaleY(.5);
        transform: scaleY(.5);
        background-color: #c8c7cc;
    }

    .input-row {
        display: flex;
        flex-direction: row;
        position: relative;
        
        height: 40px;
        line-height: 40px;
    }

    .input-row .title {
        width: 70px;
        padding-left: 15px;
    }

    .input-row.border::after {
        position: absolute;
        right: 0;
        bottom: 0;
        left: 8px;
        height: 1px;
        content: '';
        -webkit-transform: scaleY(.5);
        transform: scaleY(.5);
        background-color: #c8c7cc;
    }

    .btn-row {
        margin-top: 25px;
        padding: 10px;
    }

    button.primary {
        background-color: #0faeff;
    }
</style>