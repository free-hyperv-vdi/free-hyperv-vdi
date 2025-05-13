import message from './message.js';

const config = {
	
	top: 'top',
	
	bottom: 'bottom',
	
	center: 'center',
	
	message: 'top',
	
	dialog: 'center',
	
	share: 'bottom',
}

export default {
	data() {
		return {
			config: config,
			popupWidth: 0,
			popupHeight: 0
		}
	},
	mixins: [message],
	computed: {
		isDesktop() {
			return this.popupWidth >= 500 && this.popupHeight >= 500
		}
	},
	mounted() {
		const fixSize = () => {
			const {
				windowWidth,
				windowHeight,
				windowTop
			} = uni.getSystemInfoSync()
			this.popupWidth = windowWidth
			this.popupHeight = windowHeight + windowTop
		}
		fixSize()
		
		window.addEventListener('resize', fixSize)
		this.$once('hook:beforeDestroy', () => {
			window.removeEventListener('resize', fixSize)
		})
		
	},
}
