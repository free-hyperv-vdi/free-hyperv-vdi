export default {
	created() {
		if (this.type === 'share') {
			
			this.mkclick = false
		}
	},
	methods: {
		customOpen() {
			console.log('share 打开了');
		},
		customClose() {
			console.log('share 关闭了');
		}
	}
}
